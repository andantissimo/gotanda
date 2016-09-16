package log

import (
	"fmt"
	"log"
)

var theWriter logWriter = &stdWriter{}

// Facility is one of LOCAL0..LOCAL7
type Facility int

// LOCAL0 -> syslog.LOG_LOCAL0
const (
	LOCAL0 Facility = iota
	LOCAL1
	LOCAL2
	LOCAL3
	LOCAL4
	LOCAL5
	LOCAL6
	LOCAL7
)

type logWriter interface {
	Info(m string) error
	Warning(m string) error
	Err(m string) error
}

type stdWriter struct{}

func (w *stdWriter) Info(m string) error {
	log.Println("[info] " + m)
	return nil
}

func (w *stdWriter) Warning(m string) error {
	log.Println("[warning] " + m)
	return nil
}

func (w *stdWriter) Err(m string) error {
	log.Println("[error] " + m)
	return nil
}

// Infof logs a message with severity LOG_INFO|EVENTLOG_INFORMATION_TYPE
func Infof(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	theWriter.Info(message)
}

// Warningf logs a message with severity LOG_WARNING|EVENTLOG_WARNING_TYPE
func Warningf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	theWriter.Warning(message)
}

// Errf logs a message with severity LOG_ERR|EVENTLOG_ERROR_TYPE
func Errf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	theWriter.Err(message)
}
