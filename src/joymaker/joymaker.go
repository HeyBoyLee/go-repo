package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"encoding/json"
	"sort"
	"crypto/sha1"
)

const (
	APPID = "001"
	APPKEY = "f804de3b53bb42860153bb4934780000"
	VERSION = "v1"
	IP = "10.232.46.11"
	URL = "http://joymakeit.ticp.net:8888/AUTH"
)

type BaseInfo struct {
	AppId   string		`json:"appid"`
	Sign    string		`json:"sign"`
	Version string      	`json:"version"`
}

func generateSignature(arr []string){
	//arr.
}

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8001")
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

	params := []string{APPID, APPKEY, IP}
	sort.Strings(params)
	generateSignature(params)
	content := strings.Join(params , "")
	fmt.Println(content)
	h := sha1.New()
	h.Write([]byte(content))
	fmt.Println("------")
	sign := fmt.Sprintf("%x\n" , h.Sum(nil))
	//sign := string(h.Sum(nil))

	bean := &BaseInfo{
		AppId: APPID,
		Sign: sign,
		Version: VERSION }
	resBean, _ := json.Marshal(bean);
	//fmt.Println(resBean)

	client := &http.Client{}
	// - 1
	//req, _ := http.NewRequest("POST","http://127.0.0.1:8001/submit", bytes.NewBuffer(jsonStr))
	// - 2
	req, _ := http.NewRequest("POST",URL, strings.NewReader(string(resBean)))

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