package echoext

import "github.com/labstack/echo/v4"

type HandlerEnable interface {
	HandlerFunc() echo.HandlerFunc
	Method() string
	Path() string
	Middlewares() []echo.MiddlewareFunc
}
