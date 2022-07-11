package main

import (
	"fmt"
	"time"

	"github.com/biligo/logagent/conf"
	"github.com/biligo/logagent/kafka"
	"github.com/biligo/logagent/taillog"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	//1. read log
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
			//2. send to kafka
		default:
			time.Sleep(time.Second)
		}
	}
}

// logagent entry program
func main() {
	// 0. load ini file
	// cfg, err := ini.Load("./conf/config.ini")
	// fmt.Println(cfg.Section("kafka").Key("address"))
	// fmt.Println(cfg.Section("kafka").Key("topic"))
	// fmt.Println(cfg.Section("taillog").Key("path"))

	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini kafka failed, err:%v\n", err)
	}

	// 1. init kafka
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("kafka init failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka succeed")

	// 2. open log to collect log
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("taillog init failed, err:%v\n", err)
		return
	}
	fmt.Println("init tailog succeed")

	run()
}
