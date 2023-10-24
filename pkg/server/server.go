package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/wjhdec/echo-ext/pkg/config"
)

type ServerConfig struct {
	ConfigOptions *ConfigOptions
	ServerOptions *ServerOptions
}

type Server struct {
	Config ServerConfig

	e         *echo.Echo
	rootGroup *echo.Group
	routers   []Router
}

func NewServer(opt *ServerOptions) (*Server, error) {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = CustomHttpErrorHandler

	svrName := "server"
	if opt.Name != "" {
		svrName = "server." + opt.Name
	}
	cfgOpt := NewConfigOptions(svrName)
	rootGroup := e.Group(cfgOpt.BasePath)
	return &Server{
		e: e, rootGroup: rootGroup,
		Config: ServerConfig{cfgOpt, opt},
	}, nil
}

func (s *Server) AddMiddleware(middleware ...echo.MiddlewareFunc) {
	s.e.Use(middleware...)
}

func (s *Server) Echo() *echo.Echo {
	return s.e
}

func (s *Server) RootGroup() *echo.Group {
	return s.rootGroup
}

func (s *Server) AddRouter(router ...Router) {
	s.routers = append(s.routers, router...)
}

type RouterFnc func(group *echo.Group, config ServerConfig) (Router, error)

func (s *Server) AddRouterFnc(fncs ...RouterFnc) error {
	for _, fnc := range fncs {
		r, err := fnc(s.rootGroup, s.Config)
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
	svrOpt := s.Config.ServerOptions
	cfgOpt := s.Config.ConfigOptions
	s.rootGroup.GET("/info", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"version":      svrOpt.Version,
			"current_time": time.Now(),
		})
	})
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfgOpt.Port),
		Handler: s.e,
	}

	go func() {
		slog.Info("start server", "version", svrOpt.Version, "port", cfgOpt.Port, "config-file", config.ConfigFileUsed())
		if cfgOpt.TLSKey == "" || cfgOpt.TLSPem == "" {
			slog.Debug("not use tls")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				slog.Error("Start service error", err)
			}
		} else {
			slog.Debug("use tls")
			if err := srv.ListenAndServeTLS(cfgOpt.TLSPem, cfgOpt.TLSKey); err != http.ErrServerClosed {
				slog.Error("Start service error", err)
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
		slog.Error("server shutdown error", err)
	}
}
