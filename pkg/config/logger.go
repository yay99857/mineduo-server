package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	warn   *log.Logger
	err    *log.Logger
	info   *log.Logger
	debug  *log.Logger
	writer io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ltime|log.Ldate)

	return &Logger{
		warn:   log.New(writer, " WARN ==> ", logger.Flags()),
		err:    log.New(writer, " ERROR ==> ", logger.Flags()),
		info:   log.New(writer, " INFO ==> ", logger.Flags()),
		debug:  log.New(writer, " DEBUG ==> ", logger.Flags()),
		writer: writer,
	}
}

func (l *Logger) Warn(v ...interface{}) {
    l.warn.Println(v...)
}
func (l *Logger) Error(v ...interface{}) {
    l.err.Println(v...)
}
func (l *Logger) Info(v ...interface{}) {
    l.info.Println(v...)
}
func (l *Logger) Debug(v ...interface{}) {
    l.debug.Println(v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
    l.warn.Printf(format, v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
    l.err.Printf(format, v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
    l.info.Printf(format, v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
    l.debug.Printf(format, v...)
}
