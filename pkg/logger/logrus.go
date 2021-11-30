package logger

import (
	"fmt"
	"io"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Init(out io.Writer) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	})

	logrus.SetOutput(out)
}

func Trace(msg ...interface{}) {
	logrus.Trace(msg...)
}
func Tracef(format string, msg ...interface{}) {
	logrus.Tracef(format, msg...)
}

func Debug(msg ...interface{}) {
	logrus.Debug(msg...)
}
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(msg ...interface{}) {
	logrus.Fatal(msg...)
}
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
