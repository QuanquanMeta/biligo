package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	//etcdWatch()
	etcdPut()
}

func etcdPut() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("new client failed, err:%v\n", err)
	}
	defer cli.Close()

	timeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	jsonValue := `[{"path":"c:/tmp/nginx.log", "topic":"web_log"},{"path":"d:/xxx/redis.log", "topic":"redis_log"},{"path":"d:/yyy/mysql.log", "topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", jsonValue)
	cancel()
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
}

func etcdPutGet() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("new client failed, err:%v\n", err)
	}
	defer cli.Close()

	timeout := time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	_, err = cli.Put(ctx, "xiang", "wang")
	cancel()
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

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "xiang")
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
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func etcdWatch() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("new client failed, err:%v\n", err)
	}
	fmt.Println("wacher connected etcd successful")
	defer cli.Close()

	ch := cli.Watch(context.Background(), "xiang")

	// get vlue from ch, which is watiching key
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v Key:%v, Value:%v", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}

}
