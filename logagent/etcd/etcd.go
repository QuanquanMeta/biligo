package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
)

// log info
type LogEntry struct {
	Path  string `json:"path"`  // log path
	Topic string `json:"topic"` // log to be sent to kafka
}

// init etcd
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{ // do NOT add := to make it local!
		Endpoints: []string{addr},
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Printf("new client failed, err:%v\n", err)
	}
	return
}

// get key from etcd
// C:/temp/nginx.log web_log
// D:/xxx/redis.log redis_log
func GetConf(key string) (logEntries []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	// use the response
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err := json.Unmarshal(ev.Value, &logEntries)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, err:%v\n", err)
			return nil, err
		}
	}
	return
}

// etcd watch func
func WatchConf(key string, newConfch chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)

	// get vlue from ch, which is watiching key
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v Key:%v, Value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// notify taillog.tskmgr
			// 1. decide type first
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				// not delete
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
			}

			fmt.Printf("Get new conf!:%v\n", newConf)

			newConfch <- newConf
		}
	}
}
