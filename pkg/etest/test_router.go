package etest

import (
	"github.com/labstack/echo/v4"

	"github.com/wjhdec/echo-ext/pkg/server"
)

func LoadRouter(routerFunc server.RouterFnc, authorMap map[string]server.Author, middleware ...echo.MiddlewareFunc) (*server.Server, error) {
	opt := server.NewServerOptions().SetName("test").AddAuthMap(authorMap)
	serv, err := server.NewServer(opt)
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
	if r, err := routerFunc(baseGroup, serv.Config); err != nil {
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
