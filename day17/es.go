package main

// ES insert data demo

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Student{Name: "rion", Age: 22, Married: false}

	// chain
	put1, err := client.Index().
		Index("student").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("index student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
