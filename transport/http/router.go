package http

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// client => server => go application => | middleware => | router

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	s.App.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	s.App.GET("/ready", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	s.App.GET("/live", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	v1.GET("/hello123123", s.handler.Get)
	v1.POST("/user/", s.handler.CreateUser)
	v1.POST("/some", s.handler.Get)
}
