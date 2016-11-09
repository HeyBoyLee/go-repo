package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"net/url"
	//"encoding/json"
)

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func httpGet() {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{

	}
	fmt.Println(string(body))
}

func httpPost(){
	v := url.Values{}
	v.Set("huifu", "hello world")

	client := &http.Client{}
	req, _ := http.NewRequest("POST","http://127.0.0.1:8001/submit",
		strings.NewReader(string(v.Encode())))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.set("Content-Type","application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func main(){
	httpPost()
}