package handler

import (
	"fmt"
	"path_project/internal/db"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	db *db.Queries
	e *echo.Echo
}

func NewServer(db *db.Queries) *Server {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	
	return &Server{
		db: db,
		e: e,
	}
}


func (s *Server) Start(addr string) {
	s.registerRoutes()

	for _, route := range s.e.Router().Routes() {
		fmt.Println(route.Method, route.Path, route.Parameters)
	}
	
	
	err := s.e.Start(addr)
	if err != nil {
		s.e.Logger.Error("failed to start server", "error", err)
	}
}
