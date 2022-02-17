package logger

import (
	"io"
	"os"
)

const (
	// FmtEmptySeparate represents empty format string.
	FmtEmptySeparate = ""
)

// Level is the log level.
type Level uint8

const (
	// DebugLevel only use when dev.
	DebugLevel Level = iota

	// InfoLevel log some useful info.
	InfoLevel

	// WarnLevel is more important than InfoLevel.
	WarnLevel

	// ErrorLevel indicate that the program had an error occurred.
	ErrorLevel

	// PanicLevel log is print when a panic occurred.
	PanicLevel

	// FatalLevel log is print when program can not run anymore.
	FatalLevel
)

// LevelNameMapping is a mapping from log level to level name.
var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
}

// log options.
type options struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
	disableColor  bool
}

type Option func(*options)

func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

// WithOutput set log output.
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

// WithLevel sets log level.
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

// WithStdLevel sets log std level.
func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

// WithFormatter sets log formatter.
func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

// WithDisableCaller sets log disable caller. if caller is true, then the log location will not be printed.
func WithDisableCaller(disableCaller bool) Option {
	return func(o *options) {
		o.disableCaller = disableCaller
	}
}

// WithDisableColor sets log color. if disableColor is true, then the log color will not be printed.
func WithDisableColor(disableColor bool) Option {
	return func(o *options) {
		o.disableColor = disableColor
	}
}
