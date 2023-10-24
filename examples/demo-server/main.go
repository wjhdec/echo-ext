package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	slogecho "github.com/samber/slog-echo"
	"github.com/spf13/viper"
	"github.com/wjhdec/echo-ext/pkg/config"
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
		return &ResultInfo{Value: req.Name + "_" + fmt.Sprintf("%f", req.Value)}, errors.Errorf("this is error")
	})
}

func NewDemoRouter(group *echo.Group) server.Router {
	router := server.NewRouter(group)
	router.AddHandler(NewTest1Handler())
	return router
}

type DemoConfig struct {
	viper.Viper
}

func NewConfig() *DemoConfig {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../../configs/")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return &DemoConfig{*v}
}

func (c *DemoConfig) Reload() error {
	return c.ReadInConfig()
}

// UnmarshalByKey 根据key填充内容
func (c *DemoConfig) UnmarshalByKey(key string, v any) error {
	cfg := c.Sub(key)
	if cfg != nil {
		return cfg.Unmarshal(v)
	}
	return nil
}

// ValueByKey 根据key获取内容，返回 any，找不到则返回 nil
func (c *DemoConfig) ValueByKey(key string) any {
	return c.Get(key)
}

// StrValueByKey 根据key获取字符串内容，找不到返回空字符串
func (c *DemoConfig) StrValueByKey(key string) string {
	return c.GetString(key)
}

// SetByKey 覆盖配置
func (c *DemoConfig) SetByKey(key string, v any) error {
	c.Set(key, v)
	return nil
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
	svr.Run()
}
