package main

// log transfer
// get log from kafka and send it to ES

import (
	"fmt"

	"github.com/biligo/logagent/conf"
	"github.com/biligo/logagent/kafka"
	"github.com/biligo/logtransfer/es"
	"gopkg.in/ini.v1"
)

func main() {
	// 0. load ini file
	cfg := new(conf.AppConf)
	err := ini.MapTo(cfg, "../logagent/conf/config.ini")

	if err != nil {
		fmt.Printf("ini config, err:%v\n", err)
	}

	fmt.Printf("cfg:%v\n", cfg.ESConf.Address)

	// 1. send log to ES
	err = es.InitES(cfg.ESConf.Address, cfg.ESConf.ChanSize, cfg.ESConf.NumsGorutine)
	if err != nil {
		fmt.Printf("init es client failed, err:%v\n", err)
		return
	}
	// 2. init kafka
	// 2.1 connect to kafka, create partition
	// 2.2 each partition get data and send to es
	err = kafka.InitConsumer([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		fmt.Printf("kafka init failed, err:%v\n", err)
		return
	}
	//fmt.Println("init kafka succeed")

	// wait in the main
	select {}
}
