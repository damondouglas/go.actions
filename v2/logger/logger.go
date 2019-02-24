package logger

// Logger logs to console.
type Logger interface {
	Errorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
}
