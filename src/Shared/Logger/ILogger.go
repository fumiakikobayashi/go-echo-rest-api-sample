package Logger

type ILogger interface {
	Close() error
	Debug(fields ...interface{})
	Info(fields ...interface{})
	Warn(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
	Panic(fields ...interface{})
}
