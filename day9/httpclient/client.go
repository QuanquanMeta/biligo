package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// global client. to be used for request is often
var (
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
	}
)

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:9090/query/?name=sb&age=18")
	if err != nil {
		fmt.Println("Get url failed, err:", err)
		return
	}

	// get values from resp
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.body failed, err:", err)
		return
	}

	// get data
	fmt.Println((string(b)))
}

func httpNewRequest() {
	data := url.Values{}
	urlObj, _ := url.ParseRequestURI("http://127.0.0.1:9090/query/") //
	data.Set("name", "abc")
	data.Set("age", "9000")
	urlStr := data.Encode() // get a string after the url encode
	urlObj.RawQuery = urlStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)

	// resp, err := http.DefaultClient.Do(req)

	// if err != nil {
	// 	fmt.Println("get url failed, err:", err)
	// 	return
	// }

	// short connection // it is used for short requests
	// tr := &http.Transport{
	// 	DisableKeepAlives: true,
	// }

	// client := http.Client{
	// 	Transport: tr,
	// }

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("get url failed, err:", err)
		return
	}

	defer resp.Body.Close() // must close the resp

	// get values from resp
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.body failed, err:", err)
		return
	}

	// get data
	fmt.Println((string(b)))
}

// net client
func main() {
	// httpGet()
	httpNewRequest()
}
