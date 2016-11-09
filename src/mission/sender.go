package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	//"net/url"
	"encoding/json"
	//"bytes"
)

type BaseJsonBean struct {
	Code    int         `json:"code"`
	//Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func httpGet() {
	resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{

	}
	fmt.Println(string(body))
}

func httpPost(){

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	fmt.Println(jsonStr)

	//v := url.Values{}
	////v.Set("huifu", "hello world")
	bean := &BaseJsonBean{
		Code: 200,
		Message: "hello world",
	}
	resBean, _ := json.Marshal(bean);
	fmt.Println(resBean)

	client := &http.Client{}
	// - 1
	//req, _ := http.NewRequest("POST","http://127.0.0.1:8001/submit", bytes.NewBuffer(jsonStr))
	// - 2
	req, _ := http.NewRequest("POST","http://127.0.0.1:8001/submit", strings.NewReader(string(resBean)))

	req.Header.Set("Content-Type","application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main(){
	httpPost()
}