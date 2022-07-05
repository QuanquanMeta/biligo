package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("../frontpage.html")
	if err != nil {
		fmt.Println("failed to read the html, err:", err)
		return
	}
	// str := "<h1>hello world!<h1>"
	w.Write(b) // type connversion. string to []byte
}

func queryHandlerFunc(w http.ResponseWriter, r *http.Request) {
	queryPara := r.URL.Query()
	name := queryPara.Get("name")
	age := queryPara.Get("age")
	fmt.Println(name, age) // automatically get the query parameters
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func runServer() {
	http.HandleFunc("/home/", homeHandlerFunc)
	http.HandleFunc("/query/", queryHandlerFunc)
	http.ListenAndServe("127.0.0.1:9090", nil)
}

func main() {
	runServer()
}
