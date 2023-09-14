package Logger

import (
	"github.com/labstack/echo/v4"
	"go-echo-rest-api-sample/src/Shared"
)

func WithContext(ctx echo.Context, logger ILogger) {
	ctx.Set("Logger", logger)
}

func FromContext(ctx echo.Context) (ILogger, error) {
	logger := ctx.Get("Logger")
	if logger == nil {
		return nil, Shared.NewSampleError(
			"001-001",
			"logger.FromContext() failed to extract Logger from context",
		)
	}

	return logger.(ILogger), nil
}
