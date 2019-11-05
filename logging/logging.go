package logging

import (
	"github.com/sirupsen/logrus"
	//	"io"
	//	logging "log"
	"os"
)

var (
	log *logrus.Logger
)

func init() {

	log = logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false, ForceColors: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.DebugLevel)

	log.SetReportCaller(false)

	// mw := io.MultiWriter(os.Stdout, f)
	// log.SetOutput(mw)
}

// Info ...
func Info(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

func Fatal(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
