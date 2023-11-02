package echoext

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg       *Options
	e         *echo.Echo
	rootGroup *echo.Group
	routers   []Router
}

func NewServer(opt *Options) (*Server, error) {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = CustomHttpErrorHandler

	rootGroup := e.Group(opt.BasePath)
	return &Server{
		cfg:       opt,
		e:         e,
		rootGroup: rootGroup,
	}, nil
}

func (s *Server) AddMiddleware(middleware ...echo.MiddlewareFunc) {
	s.e.Use(middleware...)
}

func (s *Server) Echo() *echo.Echo {
	return s.e
}

func (s *Server) Options() *Options {
	return s.cfg
}

func (s *Server) RootGroup() *echo.Group {
	return s.rootGroup
}

type ServerGroup struct {
	*echo.Group
	Server *Server
}

// SubGroup 创建子组
func (s *ServerGroup) SubGroup(prefix string, funcs ...echo.MiddlewareFunc) *echo.Group {
	return s.Group.Group(prefix, funcs...)
}

func (s *ServerGroup) Author(key string) Author {
	return s.Server.Options().GetAuthor(key)
}

type RouterFnc func(group *ServerGroup) (Router, error)

func (s *Server) AddRouterFnc(fncs ...RouterFnc) error {
	for _, fnc := range fncs {
		r, err := fnc(&ServerGroup{s.rootGroup, s})
		if err != nil {
			return err
		}
		s.routers = append(s.routers, r)
	}
	return nil
}

func (s *Server) RegisterRouters() {
	for _, r := range s.routers {
		if err := r.Register(); err != nil {
			panic(err)
		}
	}
}

func (s *Server) Run() {
	s.RegisterRouters()
	opt := s.cfg
	s.rootGroup.GET("/info", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"version":      opt.Version,
			"current_time": time.Now(),
		})
	})
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.Port),
		Handler: s.e,
	}

	go func() {
		slog.Info("start server", "version", opt.Version, "port", opt.Port)
		if opt.TLSKey == "" || opt.TLSPem == "" {
			slog.Debug("not use tls")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				logFatalWithMsg("start service error", err)
			}
		} else {
			slog.Debug("use tls")
			if err := srv.ListenAndServeTLS(opt.TLSPem, opt.TLSKey); err != http.ErrServerClosed {
				logFatalWithMsg("start service error", err)
			}
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		logErrorWithMsg("server shutdown error", err)
	}
}
