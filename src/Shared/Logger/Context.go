package Logger

import (
	"github.com/labstack/echo/v4"
)

func WithContext(ctx echo.Context, logger ILogger) {
	ctx.Set("Logger", logger)
}

func FromContext(ctx echo.Context) ILogger {
	logger := ctx.Get("Logger")
	if logger == nil {
		panic("コンテキストからロガーの取得に失敗しました")
	}

	return logger.(ILogger)
}
