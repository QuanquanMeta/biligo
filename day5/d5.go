package main

import (
	"time"

	"github.com/biligo/day5/mylogger"
)

func main() {
	//myConsoleLogger()
	myLogger()
}

var log mylogger.Logger

// log
// func myLog() {

// 	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Printf("open file failed, %v\n", err)
// 	}
// 	log.SetOutput(fileObj)
// 	log.Println("this is a log")
// }

// func myConsoleLogger() {
// 	log = mylogger.NewConsoleLogger(mylogger.DEBUG | mylogger.WARNING | mylogger.ERROR | mylogger.FATAL)

// 	log.Debug("this is debug")
// 	log.Info("this is infor")
// 	log.Warning("this is warning")
// 	log.Error("this is error")
// 	log.Fatal("this is fatal %v, %d", "err: unpackage", 100)
// 	time.Sleep(time.Second)
// }

func myLogger() {
	log = mylogger.NewLogger(mylogger.ConsoleLoggerType)
	log = mylogger.NewLogger(mylogger.FileLoggerType)

	for i := 0; i < 100; i++ {
		log.Debug("this is debug")
		log.Info("this is infor")
		log.Warning("this is warning")
		log.Error("this is error")
		log.Fatal("this is fatal %v, %d", "err: unpackage", 100)
		time.Sleep(time.Second)
	}
}
