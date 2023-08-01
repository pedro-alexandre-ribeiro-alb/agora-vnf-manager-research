package log

import (
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	GetLevel() logrus.Level
	SetLevel(logrus.Level)
}

var LogTerminal = logrus.New()

func LogPrettyfier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return "", " " + fileName + "\t| "
}

func init() {

	LogTerminal.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableSorting:   false,
		PadLevelText:     true,
		FullTimestamp:    true,
		CallerPrettyfier: LogPrettyfier,
	})
	LogTerminal.SetReportCaller(true)

	LogTerminal.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	LogTerminal.SetLevel(logrus.InfoLevel)

	/* 	pathMap := lfshook.PathMap{
		logrus.ErrorLevel: "/var/log/error.log",
		logrus.FatalLevel: "/var/log/fatal.log",
	}

	LogFile.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		},
	)) */
}

func Debug(format string) {
	LogTerminal.Debug(format)
}

func Debugf(format string, args ...interface{}) {
	LogTerminal.Debugf(format, args)
}

func Info(format string) {
	LogTerminal.Info(format)
}

func Infof(format string, args ...interface{}) {
	LogTerminal.Infof(format, args)
}

func Warn(format string) {
	LogTerminal.Warn(format)
}

func Warnf(format string, args ...interface{}) {
	LogTerminal.Warnf(format, args)
}

func Error(format string) {
	LogTerminal.Error(format)
}

func Errorf(format string, args ...interface{}) {
	LogTerminal.Errorf(format, args)
}

func Trace(format string) {
	LogTerminal.Trace(format)
}

func Tracef(format string, args ...interface{}) {
	LogTerminal.Tracef(format, args)
}

func Print(format string) {
	LogTerminal.Print(format)
}

func Printf(format string, args ...interface{}) {
	LogTerminal.Printf(format, args)
}

func Fatal(format string) {
	LogTerminal.Fatal(format)
}

func Fatalf(format string, args ...interface{}) {
	LogTerminal.Fatalf(format, args)
}

func GetLevel() logrus.Level {
	return LogTerminal.GetLevel()
}

func SetLevel(level logrus.Level) {
	LogTerminal.SetLevel(level)
}
