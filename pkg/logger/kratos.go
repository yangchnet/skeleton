package logger

import "github.com/go-kratos/kratos/v2/log"

// Log is a kratos log helper.
func (l *logger) Log(level log.Level, keyvals ...interface{}) error {
	switch level {
	case log.LevelDebug:
		l.Debug(keyvals...)
	case log.LevelInfo:
		l.Info(keyvals...)
	case log.LevelWarn:
		l.Warn(keyvals...)
	case log.LevelError:
		l.Error(keyvals...)
	case log.LevelFatal:
		l.Fatal(keyvals...)
	}

	return nil
}
