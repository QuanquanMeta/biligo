package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// empty interface{}
type any = interface{}

// Severity
type LogLevel uint16

type LogType uint8

const (
	ConsoleLoggerType LogType = iota
	FileLoggerType
)

func NewLogger(lt LogType) Logger {

	switch lt {
	case ConsoleLoggerType:
		return NewConsoleLogger(DEBUG | WARNING | ERROR | FATAL)
	case FileLoggerType:
		return NewFileLogger(DEBUG|WARNING|ERROR|FATAL, "./log", "d5.log", 4*1024)
	default:
		return nil
	}
}

type Logger interface {
	Debug(format string, a ...any)
	Info(format string, a ...any)
	Warning(format string, a ...any)
	Error(format string, a ...any)
	Fatal(format string, a ...any)
	log(lv, format string, a ...any)
}

const (
	INVALID LogLevel = 1 << iota // 1 << 0 which is 00000000 00000001
	TRACE                        // 1 << 1 which is 00000000 00000010
	DEBUG                        // 1 << 2 which is 00000000 00000100
	INFO                         // 1 << 3 which is 00000000 00001000
	WARNING                      // 1 << 4 which is 00000000 00010000
	ERROR                        // 1 << 5 which is 00000000 00100000
	FATAL                        // 1 << 6 which is 00000000 01000000
)

func parseStrtoLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("invalid string")
		return INVALID, err
	}
}

func parseLogLevelToString(lv LogLevel) (s string, err error) {
	switch lv {
	case TRACE:
		return "TRACE", nil
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		err = errors.New("invalid LogLevel")
		return
	}
}

// runtime caller
func getInformation(n int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	file = path.Base(file)
	return funcName, file, line
}
