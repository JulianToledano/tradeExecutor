package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger = log.New()

func Set(file io.Writer, lvl string) {
	logger.SetOutput(io.MultiWriter(os.Stdout, file))
	l, _ := log.ParseLevel(lvl)
	logger.SetLevel(l)
}

func Errorf(format string, args ...interface{}) {
	if len(args) > 0 {
		logger.Errorf(format, args...)
	} else {
		logger.Errorf(format)
	}
}

func Infof(format string, args ...interface{}) {
	if len(args) > 0 {
		logger.Infof(format, args...)
	} else {
		logger.Infof(format)
	}

}

func Warnf(format string, args ...interface{}) {
	if len(args) > 0 {
		logger.Warnf(format, args...)
	} else {
		logger.Warnf(format)
	}
}

func Fatalf(format string, args ...interface{}) {
	if len(args) > 0 {
		logger.Fatalf(format, args...)
	} else {
		logger.Fatalf(format)
	}
}
