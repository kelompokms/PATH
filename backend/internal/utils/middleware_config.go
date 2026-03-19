package utils

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var RequestLoggerConfig = middleware.RequestLoggerConfig{
	LogStatus:    true,
	LogURI:       true,
	LogMethod:    true,
	LogLatency:   true,
	LogUserAgent: true,
	LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
		log.Printf("%v %v, status: %v, latency: %v, user_agent: %v", v.Method, v.URI, v.Status, v.Latency, v.UserAgent)
		return nil
	},
}
