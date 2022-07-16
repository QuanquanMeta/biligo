package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
)

type EsLogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *EsLogData
)

// init es
func InitES(addr string, chanSize int, numsGorutine int) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))

	if err != nil {
		return err
	}

	fmt.Println("connect to es success, start to send to es.")

	ch = make(chan *EsLogData, chanSize)

	for i := 0; i < numsGorutine; i++ {
		go Send()
	}

	return
}

func SendToESChan(msg *EsLogData) {
	ch <- msg
}

func Send() (err error) {
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Topic %s: %s to index %s, type %s\n", msg.Topic, put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Microsecond * 500)
		}
	}
}
