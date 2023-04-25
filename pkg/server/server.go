package server

import (
	"context"
	"echoext/pkg/config"
	"echoext/pkg/elog"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	e         *echo.Echo
	rootGroup *echo.Group
	version   string
	options   *Options
	routers   []*Router
}

func NewServer(version string, options ...*Options) *Server {
	e := echo.New()
	e.Logger = newEchoLogger(elog.Default())
	e.HideBanner = true
	e.HTTPErrorHandler = CustomHttpErrorHandler
	opt := new(Options)
	if len(options) > 0 {
		opt = options[0]
	} else {
		opt = NewOptions(config.New())
	}

	rootGroup := e.Group(opt.BasePath)
	return &Server{e: e, rootGroup: rootGroup, version: version, options: opt}
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

func (s *Server) AddRouter(router ...*Router) {
	s.routers = append(s.routers, router...)
}

func (s *Server) Run() {

	for _, r := range s.routers {
		if err := r.Register(); err != nil {
			panic(err)
		}
	}

	s.rootGroup.GET("/info", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"version":      s.version,  // git 中对应版本
			"current_time": time.Now(), // git 当前时间
		})
	})
	opt := s.options
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", opt.Port),
		Handler: s.e,
	}

	go func() {
		log.Infof("start server %s on port: %d", s.version, opt.Port)
		if opt.TLSKey == "" || opt.TLSPem == "" {
			log.Debug("not use tls")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.Fatalf("Start service error: %+v", err)
			}
		} else {
			log.Debug("use tls")
			if err := srv.ListenAndServeTLS(opt.TLSPem, opt.TLSKey); err != http.ErrServerClosed {
				log.Fatalf("Start service error: %+v", err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown error: %+v", err)
	}
}
