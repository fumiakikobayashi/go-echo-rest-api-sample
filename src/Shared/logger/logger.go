package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type zeroLog struct {
	logger zerolog.Logger
}

func NewLogger() LoggerInterface {
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

	zerolog.SetGlobalLevel(level)
	return &zeroLog{
		logger: zerolog.New(os.Stdout),
	}
}

func (l *zeroLog) Debug(message string) {
	l.logger.Debug().Msg(message)
}

func (l *zeroLog) Info(message string) {
	l.logger.Info().Msg(message)
}

func (l *zeroLog) Warn(message string) {
	l.logger.Warn().Msg(message)
}

func (l *zeroLog) Error(message string) {
	l.logger.Error().Msg(message)
}

func (l *zeroLog) Fatal(message string) {
	l.logger.Fatal().Msg(message)
}

func (l *zeroLog) Panic(message string) {
	l.logger.Panic().Msg(message)
}
