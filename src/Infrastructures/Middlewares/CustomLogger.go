package Middlewares

import (
	"github.com/labstack/echo/v4"
	sLogger "go-echo-rest-api-sample/src/Shared/Logger"
)

func CustomLogger(l sLogger.ILogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sLogger.WithContext(c, l)
			return next(c)
		}
	}
}
