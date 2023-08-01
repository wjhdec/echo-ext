package server

import (
	"github.com/labstack/echo/v4"
)

// Author 提供验证的接口
type Author interface {
	Auth() echo.MiddlewareFunc
}
