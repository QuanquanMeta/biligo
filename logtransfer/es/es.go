package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
)

// init es
func InitES(addr string) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))

	if err != nil {
		return err
	}

	fmt.Println("connect to es success")
	return
}

func Send(indexStr string, data interface{}) {
	put1, err := client.Index().
		Index(indexStr).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("index %s: id:%s to index %s, type %s\n", indexStr, put1.Id, put1.Index, put1.Type)
}
