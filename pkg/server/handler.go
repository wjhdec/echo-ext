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
func NewJsonHandler[T, R any](path, method string, f func(*T) (*R, error)) HandlerEnable {
	return &JsonHandler[T, R]{method: method, path: path, handlerFunc: f}
}

type JsonHandler[T, R any] struct {
	// method 执行方法 e.g. http.MethodGet http.MethodPut
	method      string
	path        string
	handlerFunc func(*T) (*R, error)
	middleware  []echo.MiddlewareFunc
}

func (h *JsonHandler[T, R]) Method() string {
	return h.method
}

func (h *JsonHandler[T, R]) Path() string {
	return h.path
}

func (h *JsonHandler[T, R]) HandlerFunc() echo.HandlerFunc {
	return func(c echo.Context) error {
		t := new(T)
		if err := c.Bind(t); err != nil {
			return err
		}
		r, err := h.handlerFunc(t)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, r)
	}
}

func (h *JsonHandler[T, R]) Middlewares() []echo.MiddlewareFunc {
	return h.middleware
}
