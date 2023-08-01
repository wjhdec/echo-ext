package etest

import (
	"github.com/labstack/echo/v4"
	"github.com/wjhdec/echo-ext/pkg/config"
	"github.com/wjhdec/echo-ext/pkg/server"
)

func LoadRouter(cfg config.Config, routerFunc server.RouterFnc, authorMap map[string]server.Author, middleware ...echo.MiddlewareFunc) (*server.Server, error) {
	serv, err := server.NewServer("test", authorMap, cfg)
	if err != nil {
		return nil, err
	}
	baseGroup := serv.Echo().Group("")
	if err != nil {
		return nil, err
	}
	if len(middleware) > 0 {
		serv.AddMiddleware(middleware...)
	}
	if r, err := routerFunc(baseGroup, serv.ServerConfig); err != nil {
		return nil, err
	} else {
		serv.AddRouter(r)
	}
	if err != nil {
		return nil, err
	}
	serv.RegisterRouters()
	return serv, nil
}
