package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/biligo/logagent/conf"
	"github.com/biligo/logagent/etcd"
	"github.com/biligo/logagent/kafka"
	"github.com/biligo/logagent/taillog"
	"github.com/biligo/logagent/utils"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	//1. read log
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 			//2. send to kafka
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

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
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMax)
	if err != nil {
		fmt.Printf("kafka init failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka succeed")

	// 2. init ETCD
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("etcd init failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd succeed")

	// in order to get ip for each logagent
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdconfkey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 2.1 get log info from etcd
	logEntryConf, err := etcd.GetConf(etcdconfkey)
	if err != nil {
		fmt.Printf("GetConf failed, err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success %v\n", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index:%v, value:%v\n", index, value)
	}

	// 2.2 send a watcher to monitor the changes from etcd and notify the logagent
	taillog.InitMgr(logEntryConf)
	newConfChan := taillog.NewConf()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdconfkey, newConfChan)
	wg.Wait()
	// 3 collect the logs

	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("taillog init failed, err:%v\n", err)
	// 	return
	// }
	// fmt.Println("init tailog succeed")

	// run()
}
