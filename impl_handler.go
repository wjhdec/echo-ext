package echoext

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

//#region 基础实现

type baseHandler struct {
	// method 执行方法 e.g. http.MethodGet http.MethodPut
	method      string
	path        string
	middleware  []echo.MiddlewareFunc
	handlerFunc echo.HandlerFunc
}

func (h *baseHandler) Method() string {
	return h.method
}

func (h *baseHandler) Path() string {
	return h.path
}

func (h *baseHandler) Middlewares() []echo.MiddlewareFunc {
	return h.middleware
}

func (h *baseHandler) HandlerFunc() echo.HandlerFunc {
	return h.handlerFunc
}

//#endregion 基础实现

// NewJsonHandler 返回json
func NewJsonHandler[T, R any](path, method string, f func(ctx echo.Context, req T) (resp R, err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return NewJsonStatusHandler(path, method, http.StatusOK, f, middleware...)
}

// NewJsonStatusHandler 返回带状态json
func NewJsonStatusHandler[T, R any](path, method string, successStatus int, f func(ctx echo.Context, req T) (resp R, err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return &baseHandler{
		method:     method,
		path:       path,
		middleware: middleware,
		handlerFunc: func(c echo.Context) error {
			var t T
			if err := c.Bind(&t); err != nil {
				return err
			}
			r, err := f(c, t)
			if err != nil {
				return err
			}
			if successStatus == 0 {
				successStatus = http.StatusOK
			}
			return c.JSON(successStatus, r)
		},
	}
}

// NewNoContentHandler 无返回
func NewNoContentHandler[T any](path, method string, f func(ctx echo.Context, req T) (err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return &baseHandler{
		method:     method,
		path:       path,
		middleware: middleware,
		handlerFunc: func(c echo.Context) error {
			var t T
			if err := c.Bind(&t); err != nil {
				return err
			}
			err := f(c, t)
			if err != nil {
				return err
			}
			return c.NoContent(http.StatusNoContent)
		},
	}
}

// NewStreamHandler 返回流，用于下载等
func NewStreamHandler[T any](path, method string, f func(ctx echo.Context, req T) (header http.Header, reader io.ReadCloser, err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return &baseHandler{
		method:     method,
		path:       path,
		middleware: middleware,
		handlerFunc: func(c echo.Context) error {
			var t T
			if err := c.Bind(&t); err != nil {
				return err
			}
			header, reader, err := f(c, t)
			if err != nil {
				return err
			}
			defer reader.Close()
			// 当前echo逻辑，如果content type 已经定义，则不会再填写，为避免后期echo修改，也加一下
			cntType := header.Get(echo.HeaderContentType)
			respHeader := c.Response().Header()
			for k := range header {
				respHeader.Set(k, header.Get(k))
			}

			return c.Stream(http.StatusOK, cntType, reader)
		},
	}
}

// NewHttpHandler 支持 httpHandler 扩展
func NewHttpHandler[T any](path, method string, f func(ctx echo.Context, req T) (handler http.Handler, err error), middleware ...echo.MiddlewareFunc) HandlerEnable {
	return &baseHandler{
		method:     method,
		path:       path,
		middleware: middleware,
		handlerFunc: func(c echo.Context) error {
			var t T
			if err := c.Bind(&t); err != nil {
				return err
			}
			handler, err := f(c, t)
			if err != nil {
				return err
			}
			handler.ServeHTTP(c.Response().Writer, c.Request())
			return nil
		},
	}
}
