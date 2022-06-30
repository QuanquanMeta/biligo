package mylogger

import (
	"fmt"
	"time"
)

// ConsoleLogger log struct
type ConsoleLogger struct {
	Level LogLevel
}

// logging into console
// NewLog new log constructor
func NewConsoleLogger(level LogLevel) ConsoleLogger {
	// level, err := parseStrtoLogLevel(levelStr)
	// if err != nil {
	// 	panic(err)
	// }
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) log(lv, format string, a ...any) {

	level, err := parseStrtoLogLevel(lv)
	if err != nil {
		fmt.Println("parseStrtoLogLevel() failed")
		return
	}

	if c.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInformation(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02: 15:04:05"), lv, fileName, funcName, lineNo, msg)
	}
}

// enable
func (c ConsoleLogger) enable(level LogLevel) bool {
	return level&c.Level != 0
}

func (c ConsoleLogger) Debug(format string, a ...any) {
	c.log("DEBUG", format, a...)
}

func (c ConsoleLogger) Info(format string, a ...any) {

	c.log("INFO", format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...any) {
	c.log("WARNING", format, a...)
}

func (c ConsoleLogger) Error(format string, a ...any) {
	c.log("ERROR", format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...any) {
	c.log("FATAL", format, a...)
}
