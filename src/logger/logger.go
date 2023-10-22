package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	gray      *color.Color = color.New(color.FgHiBlack)
	blue      *color.Color = color.New(color.FgHiBlue, color.Bold)
	green     *color.Color = color.New(color.FgHiGreen, color.Bold)
	yellow    *color.Color = color.New(color.FgHiYellow, color.Bold)
	brightRed *color.Color = color.New(color.FgHiRed, color.Bold)
	red       *color.Color = color.New(color.BgRed, color.FgWhite, color.Bold)
	reset     *color.Color = color.New(color.Reset)
)

const (
	logTypeDebug int = iota
	logTypeInfo
	logTypeWarning
	logTypeError
	logTypeFatal
)

func writePrefix(logType int) {
	gray.Print(time.Now().Format("Mon 15:04:05"))
	reset.Print(" ")

	switch logType {
	case logTypeDebug:
		blue.Print("[DEBUG]")
	case logTypeInfo:
		green.Print("[ INFO]")
	case logTypeWarning:
		yellow.Print("[ WARN]")
	case logTypeError:
		brightRed.Print("[ERROR]")
	case logTypeFatal:
		red.Print("[FATAL]")
	}
}

func writeArgs(args ...interface{}) {
	for _, v := range args {
		reset.Printf(" %s", v)
	}

	reset.Print("\n")
}

func Debug(args ...interface{}) {
	writePrefix(logTypeDebug)
	writeArgs(args...)
}

func Debugf(format string, args ...interface{}) {
	writePrefix(logTypeDebug)
	writeArgs(fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	writePrefix(logTypeInfo)
	writeArgs(args...)
}

func Infof(format string, args ...interface{}) {
	writePrefix(logTypeInfo)
	writeArgs(fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	writePrefix(logTypeWarning)
	writeArgs(args...)
}

func Warnf(format string, args ...interface{}) {
	writePrefix(logTypeWarning)
	writeArgs(fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	writePrefix(logTypeError)
	writeArgs(args...)
}

func Errorf(format string, args ...interface{}) {
	writePrefix(logTypeError)
	writeArgs(fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	writePrefix(logTypeFatal)
	writeArgs(args...)

	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	writePrefix(logTypeFatal)
	writeArgs(fmt.Sprintf(format, args...))

	os.Exit(1)
}
