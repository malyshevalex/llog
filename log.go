package llog

// Level defines current log level
type Level int32

const (
	LFatal Level = iota
	LError
	LWarning
	LInfo
	LDebug
)

// Logger is an interface for abstract logging instance
type Logger interface {
	// AddPrefix make new nested logger with prefix
	AddPrefix(prefix string) Logger

	// Fatal outputs message with LFatal level and exit 1
	Fatal(v ...interface{})

	// Fatalf outputs formatted message and exit 1
	Fatalf(f string, v ...interface{})

	// Error outputs message with LError level
	Error(v ...interface{})

	// Errorf outputs formatted message with LError level
	Errorf(f string, v ...interface{})

	// Warning outputs message with LWarning level
	Warning(v ...interface{})

	// Warningf outputs formatted message with LWarning level
	Warningf(f string, v ...interface{})

	// Info outputs message with LInfo level (default level)
	Info(v ...interface{})

	// Infof outputs formatted message with LInfo level (default level)
	Infof(f string, v ...interface{})

	// Debug outputs message with LDebug level
	Debug(v ...interface{})

	// Debugf outputs formatted message with LDebug level
	Debugf(f string, v ...interface{})

	// Close logger
	Close()
}
