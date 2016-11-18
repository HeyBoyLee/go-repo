package main

import (
	"fmt"
	"encoding/json"
	"joymaker"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type dbInfo struct{
	Session *mgo.Session
	LocComputeLog *mgo.Collection
}
var (
	db *dbInfo
)

func connectDB(){
	db = &dbInfo{}
	s, err := mgo.Dial("127.0.0.1:27017")
	db.Session = s
	if err != nil {
		panic(err)
	}
	//defer session.Close();
	collection := s.DB("metok_server").C("geolocComputeLog")
	db.LocComputeLog = collection
}

func disConnectDB(){
	db.Session.Close()
}

func getAuth() (bool, string) {
	params := []string{JM.APPID, JM.APPKEY, JM.IP}
	sign := JM.GenerateSignature(params)

	bean := JM.AuthInfo_t{
		AppId: JM.APPID,
		Sign: sign,
		Version: JM.VERSION }
	resBean, _ := json.Marshal(&bean);
	fmt.Println(string(resBean))
	result := JM.AuthResultInfo_t{}
	res := JM.HttpPost(JM.URL+JM.AUTH_PATH, string(resBean))
	json.Unmarshal([]byte(res) , &result)

	if result.Result && result.MessageCode == 0 {
		return true, result.Token
	}
	return false, ""
}

func conSubmit(){
	value := JM.LocComputeLog_t{}
	db.LocComputeLog.Find(bson.M{}).One(&value)

	conSubmitInfo := JM.ConSubmitInfo_t{}
	ciList := make([]JM.ConnectInfo_t,0)

	for i:=0;i<len(value.Wifis); i++{
		//fmt.Printf("[%d]:%T\n" , i , value.Wifis[i])
		ci := JM.ConnectInfo_t{
			Channel : 10,
			Ssid:"340a22be65b860da290eea7a3e171920",
			LocationTime:time.Now().Unix(),
			TerminalImei:"355533059452624",
			TerminalMac: "FC:C2:DE:DF:4B:28",
			LocationId:"jd-3423werdsdf3i43uh3u34h"}
		ci.ApMac = value.Wifis[i].Bssid
		//fmt.Println(ci)
		ciList = append(ciList , ci)
	}
	conSubmitInfo.ConnectInfoList = ciList
	conSubmitInfo.AccessToken = "5419e41f39f4464caa6090fa15ce393e"
	conSubmitInfo.AppId = "joymake123"
	resBean, _ := json.Marshal(&conSubmitInfo)
	res := JM.HttpPost(JM.URL+JM.CONSUBMIT_PATH , string(resBean))
	result := JM.CommonResultInfo_t{}
	json.Unmarshal([]byte(res) , &result)
	if result.Result && result.MessageCode == 0 {
		fmt.Println("consubmit success!")
	}else{
		fmt.Println("consubmit failed!")
	}
}

func scanSubmit(){
	value := JM.LocComputeLog_t{}
	db.LocComputeLog.Find(bson.M{}).One(&value)
	scanInfo := JM.ScanInfo_t{}

	apInfo := make([]JM.ApInfo_t , 0)
	for i:=0;i<len(value.Wifis);i++ {
		ap := JM.ApInfo_t{
			Channel : 10,
			Ssid:"340a22be65b860da290eea7a3e171920"}
		ap.ApMac = value.Wifis[i].Bssid
		//fmt.Println(ci)
		apInfo = append(apInfo , ap)
	}
	terminalInfo := make([]JM.TerminalInfo_t, 0)
	terminal := JM.TerminalInfo_t{
		ApList: apInfo,
		LocationTime: time.Now().Unix(),
		TerminalMac:"FC:C2:DE:DF:4B:28",
		TerminalImei:"355533059452624",
		Product:"mi5p",
		Brand:"xiaomi",
		Manufacturer: "xiaomi",
		LocationId:"jd-3423werdsdf3i43uh3u34h",
	}
	terminalInfo = append(terminalInfo , terminal)
	scanInfo.TerminalList = terminalInfo
	scanInfo.AccessToken = "5419e41f39f4464caa6090fa15ce393e"
	scanInfo.AppId = "joymake123"
	resBean, _ := json.Marshal(&scanInfo)
	res := JM.HttpPost(JM.URL+JM.CONSUBMIT_PATH , string(resBean))
	result := JM.CommonResultInfo_t{}
	json.Unmarshal([]byte(res) , &result)
	if result.Result && result.MessageCode == 0 {
		fmt.Println("consubmit success!")
	}else{
		fmt.Println("consubmit failed!")
	}
}

func main(){
	x := -1
	fmt.Printf("--- %T\n" , x);
	connectDB()
	// 1 , get access token
	//bPass, token :=getAuth()
	//if(!bPass){
	//	fmt.Println("get Token Error!")
	//}
	//2 , 接入定位提交
	conSubmit()

	disConnectDB()
}