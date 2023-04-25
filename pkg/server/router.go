package server

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(group *echo.Group) *Router {
	return &Router{
		Group: group,
	}
}

type Router struct {
	Group    *echo.Group
	Handlers []HandlerEnable
}

func (r *Router) AddHandler(handler ...HandlerEnable) {
	r.Handlers = append(r.Handlers, handler...)
}

// Register 注册 handler
func (r *Router) Register() error {
	for _, h := range r.Handlers {
		r.Group.Add(h.Method(), h.Path(), h.HandlerFunc(), h.Middlewares()...)
	}
	return nil
}
