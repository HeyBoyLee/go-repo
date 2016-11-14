package main

import (
	"fmt"
	"encoding/json"
	"joymaker"
)

type AuthInfo struct {
	AppId   string		`json:"appid"`
	Sign    string		`json:"sign"`
	Version string      	`json:"version"`
}

type ConnectInfo struct {
	ApMac string		`json:"apMAC"`
	Channel int8		`json:"channel"`
	Ssid string		`json:"ssid"`
	LocationTime int64	`json:"locationTime"`
	TerminalMac string	`json:"terminalMAC"`
	TerminalImei string	`json:"terminalIMEI"`
	LocationId string	`json:"locationId"`
}

type ConSubmitInfo struct{
	ConnectInfoList []ConnectInfo	`json:"connectInfoList"`
	AccessToken string		`json:"accessToken"`
	AppId string			`json:"appId"`
}

func getAuth() bool {
	params := []string{JM.APPID, JM.APPKEY, JM.IP}
	sign := JM.GenerateSignature(params)

	bean := &AuthInfo{
		AppId: JM.APPID,
		Sign: sign,
		Version: JM.VERSION }
	resBean, _ := json.Marshal(bean);
	fmt.Println(string(resBean))
	JM.HttpPost(JM.URL+JM.AUTH_PATH, string(resBean))
	return false
}

func main(){
	getAuth()
}