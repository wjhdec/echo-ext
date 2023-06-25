package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewJsonHandler 新建方法
func NewJsonHandler[T, R any](path, method string, f func(ctx echo.Context, req T) (resp R, err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return &jsonHandler[T, R]{method: method, path: path, handlerFunc: f, middleware: middleware}
}

type jsonHandler[T, R any] struct {
	// method 执行方法 e.g. http.MethodGet http.MethodPut
	method      string
	path        string
	handlerFunc func(echo.Context, T) (R, error)
	middleware  []echo.MiddlewareFunc
}

func (h *jsonHandler[T, R]) Method() string {
	return h.method
}

func (h *jsonHandler[T, R]) Path() string {
	return h.path
}

func (h *jsonHandler[T, R]) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		var t T
		if err := c.Bind(&t); err != nil {
			return err
		}
		r, err := h.handlerFunc(c, t)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, r)
	}
}

func (h *jsonHandler[T, R]) Middlewares() []echo.MiddlewareFunc {
	return h.middleware
}
