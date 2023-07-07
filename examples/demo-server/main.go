// Package main 测试内容
// 启动后访问 http://localhost:8888/my-test?name=a&value=200
package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/elog"
	"github.com/wjhdec/echo-ext/pkg/server"
)

type ResultInfo struct {
	Value string
}

type Req struct {
	Name  string  `query:"name"`
	Value float64 `query:"value"`
}

func NewTest1Handler() server.HandlerEnable {
	return server.NewJsonHandler("", http.MethodGet, func(_ echo.Context, req Req) (*ResultInfo, error) {
		return &ResultInfo{Value: req.Name + "_" + fmt.Sprintf("%f", req.Value)}, nil
	})
}

func NewDemoRouter(group *echo.Group) server.Router {
	router := server.NewRouter(group)
	router.AddHandler(NewTest1Handler())
	return router
}

func main() {
	cfg, err := config.New()
	if err != nil {
		elog.Panic(err)
	}
	elog.OverrideGlobalLogger(elog.NewLogger(cfg))
	svr, err := server.NewServer("v0.0")
	if err != nil {
		elog.Error(err)
	}
	svr.AddMiddleware(middleware.Logger(), middleware.Recover())
	svr.AddRouter(NewDemoRouter(svr.RootGroup()))
	svr.Run()
}
