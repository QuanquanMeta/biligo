package main

import (
	"github.com/biligo/logagent/kafka"
)

func main() {
	//kafka.InitConsumer([]string{"127.0.0.1:9092"}, "web_log")
	kafka.Consumer()
}
