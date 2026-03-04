package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (s *Server) registerRoutes(){
	s.e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	s.e.GET("/api", func(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello Wrold!")})
}
