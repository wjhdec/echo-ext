package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HandlerEnable interface {
	HandlerFunc() echo.HandlerFunc
	Method() string
	Path() string
	Middlewares() []echo.MiddlewareFunc
}

// NewJsonHandler 新建方法
func NewJsonHandler[T any](path, method string, f func(T) (any, error)) HandlerEnable {
	return &JsonHandler[T]{method: method, path: path, handlerFunc: f}
}

type JsonHandler[T any] struct {
	// method 执行方法 e.g. http.MethodGet http.MethodPut
	method      string
	path        string
	handlerFunc func(T) (any, error)
	middleware  []echo.MiddlewareFunc
}

func (h *JsonHandler[T]) Method() string {
	return h.method
}

func (h *JsonHandler[T]) Path() string {
	return h.path
}

func (h *JsonHandler[T]) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		var t T
		if err := c.Bind(&t); err != nil {
			return err
		}
		r, err := h.handlerFunc(t)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, r)
	}
}

func (h *JsonHandler[T]) Middlewares() []echo.MiddlewareFunc {
	return h.middleware
}
