package Logger

type ILogger interface {
	Debug(fields ...interface{})
	Info(fields ...interface{})
	Warn(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
	Panic(fields ...interface{})
}
