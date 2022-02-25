package logger

import (
	"fmt"
	"time"
)

// TextFormatter implemented Formatter interface.
type TextFormatter struct {
	IgnoreBasicFields bool
}

// Format format log to text.
func (f *TextFormatter) Format(e *Entry) error {
	if !e.logger.opt.disableColor {
		e.Buffer.WriteString(string(e.Level.Color()))
		defer e.Buffer.WriteString(string(reset))
	}

	defer e.Buffer.WriteString("\n")

	if !f.IgnoreBasicFields {
		e.Buffer.WriteString(fmt.Sprintf("%s %s", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]))
		if e.File != "" {
			if !e.logger.opt.enableAbsPath {
				short := e.File
				for i := len(e.File) - 1; i > 0; i-- {
					if e.File[i] == '/' {
						if e.File[i] == '/' {
							short = e.File[i+1:]
							break
						}
					}
				}
				e.Buffer.WriteString(fmt.Sprintf(" %s:%d", short, e.Line))
			} else {
				e.Buffer.WriteString(fmt.Sprintf(" %s:%d", e.File, e.Line))
			}
		}
		e.Buffer.WriteString(" ")
	}

	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}

	return nil
}
