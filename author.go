package echoext

import "github.com/labstack/echo/v4"

type Author interface {
	Auth() echo.MiddlewareFunc
}
