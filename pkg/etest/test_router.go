package etest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/server"
)

func LoadRouter(cfg config.Config, rfunc func(group *echo.Group, cfg config.Config) (server.Router, error)) (*server.Server, error) {
	serv, err := server.NewServer("test", cfg)
	if err != nil {
		return nil, err
	}
	baseGroup := serv.Echo().Group("")
	serv.AddMiddleware(middleware.Recover(), middleware.Logger())
	if err != nil {
		return nil, err
	}
	r, err := rfunc(baseGroup, cfg)
	if err != nil {
		return nil, err
	}
	serv.AddRouter(r)
	serv.RegisterRouters()
	return serv, nil
}
