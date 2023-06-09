package sdk

type LoggerInterface interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
	Panic(message string)
}
