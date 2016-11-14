package JM

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"sort"
	"crypto/sha1"
)

func GenerateSignature(arr []string) string{
	sort.Strings(arr)
	content := strings.Join(arr , "")
	fmt.Println(content)
	h := sha1.New()
	h.Write([]byte(content))
	fmt.Println("------")
	sign := fmt.Sprintf("%x\n" , h.Sum(nil))
	return sign
}

func HttpGet() {
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

func HttpPost(url string, params string){
	fmt.Println(url)
	client := &http.Client{}
	// - 1
	//req, _ := http.NewRequest("POST","http://127.0.0.1:8001/submit", bytes.NewBuffer(jsonStr))
	// - 2

	req, _ := http.NewRequest("POST",url, strings.NewReader(params))

	req.Header.Set("Content-Type","application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
