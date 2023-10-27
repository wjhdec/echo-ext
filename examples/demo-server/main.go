package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/server"
)

type ResultInfo struct {
	Value float64
}

// Req 参数，使用 echo.Bind 匹配
type Req struct {
	V1 float64 `query:"v1"`
	V2 float64 `query:"v2"`
}

// NewSumFunctionHandler 加和接口
func NewSumFunctionHandler() server.HandlerEnable {
	return server.NewJsonHandler("/sum", http.MethodGet, func(_ echo.Context, req Req) (*ResultInfo, error) {
		return &ResultInfo{Value: req.V1 + req.V2}, nil
	})
}

// NewErrorDemoHandler 错误示例接口
func NewErrorDemoHandler() server.HandlerEnable {
	return server.NewJsonHandler("/demo-error", http.MethodGet, func(_ echo.Context, req Req) (*ResultInfo, error) {
		// 如果想记录堆栈信息，可使用 github.com/pkg/errors
		return &ResultInfo{}, fmt.Errorf("this is error")
	})
}
func NewDemoRouter(group *echo.Group) server.Router {
	router := server.NewRouter(group)
	router.AddHandler(NewSumFunctionHandler(), NewErrorDemoHandler())
	return router
}

func main() {
	config.SetDefaultConfig(NewConfig())
	logOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
	slog.SetDefault(logger)
	svr, err := server.NewServer(server.NewServerOptions())
	if err != nil {
		slog.Error("", slog.Any("error", err))
	}
	svr.AddMiddleware(slogecho.New(logger), middleware.Recover())
	svr.AddRouter(NewDemoRouter(svr.RootGroup()))
	slog.Info("test url:", slog.String("url", "http://localhost:8888/my-test/sum?v1=1&v2=10"))
	slog.Info("test error url:", slog.String("url", "http://localhost:8888/my-test/demo-error"))
	svr.Run()
}
