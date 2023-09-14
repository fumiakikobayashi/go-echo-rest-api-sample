package Logger

import (
	"github.com/rs/zerolog"
	"go-echo-rest-api-sample/src/Shared"
	"os"
)

type zeroLog struct {
	logger zerolog.Logger
}

// NewLogger
// ex.`{"level":"info","message":"Your log message","time":1631618390,"caller":"/path/to/your/source/file.go:123"}`
func NewLogger() ILogger {
	logLevelMap := map[string]zerolog.Level{
		"debug": zerolog.DebugLevel,
		"info":  zerolog.InfoLevel,
		"warn":  zerolog.WarnLevel,
		"error": zerolog.ErrorLevel,
		"fatal": zerolog.FatalLevel,
		"panic": zerolog.PanicLevel,
	}
	var level zerolog.Level
	if val, ok := logLevelMap[os.Getenv("LOG_LEVEL")]; ok {
		level = val
	} else {
		level = zerolog.InfoLevel
	}

	// FIXME 深さを「4」と指定しているが、実装に依存するため別の方法を検討する
	zerolog.CallerSkipFrameCount = 4
	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &zeroLog{
		logger: zerolog.New(os.Stdout).With().Timestamp().Caller().Logger(),
	}
}

func (l *zeroLog) Debug(fields ...interface{}) {
	l.log(zerolog.DebugLevel, fields...)
}

func (l *zeroLog) Info(fields ...interface{}) {
	l.log(zerolog.InfoLevel, fields...)
}

func (l *zeroLog) Warn(fields ...interface{}) {
	l.log(zerolog.WarnLevel, fields...)
}

func (l *zeroLog) Error(fields ...interface{}) {
	l.log(zerolog.ErrorLevel, fields...)
}

func (l *zeroLog) Fatal(fields ...interface{}) {
	l.log(zerolog.FatalLevel, fields...)
}

func (l *zeroLog) Panic(fields ...interface{}) {
	l.log(zerolog.PanicLevel, fields...)
}

func (l *zeroLog) log(level zerolog.Level, fields ...interface{}) {
	if len(fields) == 0 {
		return
	}

	event := l.logger.WithLevel(level)

	// スタックトレースの設定
	if shouldLogStackTrace(level, Shared.NewEnv(os.Getenv("APP_ENV"))) {
		event = event.Stack()
	}

	// fieldsの1つ目は必ず"message"の値として扱う
	if msg, ok := fields[0].(string); ok {
		event = event.Str("message", msg)
	} else {
		event = event.Interface("message", fields[0])
	}

	// 残りのfieldsはキーと値のペアとして扱う
	for i := 1; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		// キーが文字列でない場合はスルー
		if !ok {
			continue
		}

		// 値がある場合は、その型に応じてログを出力
		if i+1 < len(fields) {
			switch v := fields[i+1].(type) {
			case int:
				event = event.Int(key, v)
			case string:
				event = event.Str(key, v)
			case bool:
				event = event.Bool(key, v)
			default:
				event = event.Interface(key, v)
			}
		} else {
			// 値がない場合は、空文字列として扱う
			event = event.Str(key, "")
		}
	}

	event.Send()
}

// shouldLogStackTrace ログレベルがError＆本番環境以外の場合は、スタックトレースを出力する
func shouldLogStackTrace(level zerolog.Level, env Shared.Environment) bool {
	return level == zerolog.ErrorLevel && !env.IsProduction()
}
