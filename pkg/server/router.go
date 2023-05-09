package server

import (
	"github.com/labstack/echo/v4"
)

func NewRouter(group *echo.Group) Router {
	return &router{
		Group: group,
	}
}

type Router interface {
	AddHandler(handler ...HandlerEnable)
	Register() error
}

type router struct {
	Group    *echo.Group
	handlers []HandlerEnable
}

func (r *router) AddHandler(handler ...HandlerEnable) {
	r.handlers = append(r.handlers, handler...)
}

func (r *router) Register() error {
	for _, h := range r.handlers {
		r.Group.Add(h.Method(), h.Path(), h.HandlerFunc(), h.Middlewares()...)
	}
	return nil
}
