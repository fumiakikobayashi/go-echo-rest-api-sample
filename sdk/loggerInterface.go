package sdk

type LoggerInterface interface {
	Info(message string)
	Error(message string)
}
