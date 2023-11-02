package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	echoext "github.com/wjhdec/echo-ext/v2"
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
func NewSumFunctionHandler() echoext.HandlerEnable {
	return echoext.NewJsonHandler("/sum", http.MethodGet, func(_ echo.Context, req Req) (*ResultInfo, error) {
		return &ResultInfo{Value: req.V1 + req.V2}, nil
	})
}

// NewErrorDemoHandler 错误示例接口

func NewErrorDemoHandler() echoext.HandlerEnable {
	return echoext.NewJsonHandler("/demo-error", http.MethodGet, func(_ echo.Context, req Req) (*ResultInfo, error) {
		// 如果想记录堆栈信息，可使用 github.com/pkg/errors
		return &ResultInfo{}, fmt.Errorf("this is error")
	})
}
func NewDemoRouter(group *echoext.ServerGroup) (echoext.Router, error) {
	router := echoext.NewRouter(group.Group)
	router.AddHandler(NewSumFunctionHandler(), NewErrorDemoHandler())
	return router, nil
}

// url: http://localhost:8181/my-test/sum?v1=1&v2=10
// url: http://localhost:8888/my-test/demo-error
func main() {
	opt := &echoext.Options{
		Version:  "v1.0",
		Port:     8181,
		BasePath: "my-test",
	}
	logOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpts))
	slog.SetDefault(logger)
	opt.SetVersion("v1.1")
	svr, err := echoext.NewServer(opt)
	if err != nil {
		slog.Error("server error", slog.Any("error", err))
	}
	svr.AddMiddleware(slogecho.New(logger), middleware.Recover())
	if err := svr.AddRouterFnc(NewDemoRouter); err != nil {
		slog.Error("server error", slog.Any("error", err))
	}
	svr.Run()
}
