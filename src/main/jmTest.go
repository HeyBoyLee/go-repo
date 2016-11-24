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

func connectDB1(){
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

func disConnectDB1(){
	db.Session.Close()
}

func getAuth1() (bool, string) {
	params := []string{JM.APPID, JM.APPKEY, JM.IP}
	sign := JM.GenerateSignature(params)

	bean := JM.AuthInfo_t{
		AppId: JM.APPID,
		Sign: sign,
		Version: JM.VERSION }
	resBean, _ := json.Marshal(&bean);
	fmt.Println(string(resBean))
	result := JM.CommonResultTest_t{}
	res := JM.HttpPost(JM.URL+JM.AUTH_PATH, string(resBean))

	json.Unmarshal([]byte(res) , &result)

	if result.Code == 0 && result.Msg == "" {
		//authResult := JM.AuthResultInfo_t{}

		//authResult,_ := data.([]string)
		authResult,_ := result.Data.(JM.AuthResultData)
		fmt.Println(authResult.GetAuthData())
		return  true, "token"//result.Data.Token//authResult["token"]
	}
	return false, ""
}

func conSubmit1(t string){
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
	conSubmitInfo.AccessToken = t
	conSubmitInfo.AppId = JM.APPID
	resBean, _ := json.Marshal(&conSubmitInfo)
	res := JM.HttpPost(JM.URL+JM.CONSUBMIT_PATH , string(resBean))
	result := JM.CommonResultInfo_t{}
	json.Unmarshal([]byte(res) , &result)
	if result.Code == 0 && result.Msg == "" {
		fmt.Println("consubmit success!")
	}else{
		fmt.Println("consubmit failed!")
	}
}

func scanSubmit1(t string){
	value := JM.LocComputeLog_t{}
	db.LocComputeLog.Find(bson.M{}).One(&value)
	scanInfo := JM.ScanInfo_t{}

	apInfo := make([]JM.ApInfo_t , 0)
	for i:=0;i<len(value.Wifis);i++ {
		ap := JM.ApInfo_t{
			Channel : 10}
		ap.ApMac = value.Wifis[i].Bssid
		ap.Rssi = value.Wifis[i].Dbm
		ap.Ssid = JM.GetMd5String(value.Wifis[i].Bssid)
		//fmt.Println(ci)
		apInfo = append(apInfo , ap)
	}
	//terminalInfo := make([]JM.TerminalInfo_t, 0)
	//terminal := JM.TerminalInfo_t{
	scanInfo.ApList= apInfo
	scanInfo.LocationTime = time.Now().UnixNano()
	scanInfo.TerminalMac = "78:02:f8:3e:e5:77"
	scanInfo.TerminalImei = "861414033593558"
	scanInfo.Product = "mi5"
	scanInfo.Brand = "xiaomi"
	scanInfo.Manufacturer = "xiaomi"
	scanInfo.LocationId = "jd-3423werdsdf3i43uh3u34h"
	scanInfo.AppId = JM.APPID
	scanInfo.AccessToken = t
	//
	resBean, _ := json.Marshal(&scanInfo)
	res := JM.HttpPost(JM.URL+JM.SCANSUBMIT_PATH , string(resBean))
	result := JM.SubmitResultInfo_t{}
	json.Unmarshal([]byte(res) , &result)
	if result.Code == 0 && result.Msg == "" {
		//scanResult := JM.ScanResultInfo_t{}
		//scanResult = result.Data.(JM.ScanResultInfo_t)
		//fmt.Println(scanResult)
		//fmt.Println("consubmit success!")
		fmt.Println(result);
	}else{
		fmt.Println("consubmit failed!")
	}
}

func main(){
	connectDB1()
	// 1 , get access token
	bPass, token :=getAuth1()
	if(!bPass){
		fmt.Println("get Token Failed!")
	}
	fmt.Println(token)
	//2 , 接入定位提交
	//conSubmit1(token)
	//3, 扫描定位接口
	//scanSubmit1(token)

	disConnectDB1()
}
