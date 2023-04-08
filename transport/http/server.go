package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"repotest/config"
	_ "repotest/docs"
	"repotest/logger"
	"repotest/transport/http/handler"
	middleware2 "repotest/transport/http/middleware"

	"context"
	"fmt"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
	m       middleware2.JWTAuth
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	c := jaegertracing.New(s.App, nil)
	defer c.Close()
	s.SetupRoutes()

	go func() {
		if err := s.App.Start(fmt.Sprintf(":%d", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	l := logger.Logger(context.Background())
	e.Use(middleware.RequestID())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	//s.prom.Use(e)
	return e
}
