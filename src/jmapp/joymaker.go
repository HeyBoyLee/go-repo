package main

import (
	"fmt"
	"encoding/json"
	"joymaker"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	_ "github.com/robfig/cron"
)

type dbInfo struct{
	Session *mgo.Session
	LocComputeLog *mgo.Collection
	JmPoiLog *mgo.Collection
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
	computeLog := s.DB("metok_server").C("geolocComputeLog")
	poiLog := s.DB("metok_server").C("jmPoiLog")
	db.LocComputeLog = computeLog
	db.JmPoiLog = poiLog
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
	result := JM.CommonResultInfo_t{}
	res := JM.HttpPost(JM.URL+JM.AUTH_PATH, string(resBean))

	json.Unmarshal([]byte(res) , &result)

	if result.Code == 0 && result.Msg == "" {
		//fmt.Println(result)
		return  true, result.Data.Token
	}
	return false, ""
}

func conSubmit(t string){
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

func scanSubmit(t string , value *JM.LocComputeLog_t) int {
	if len(value.Wifis) == 0 {
		return 0
	}

	scanInfo := JM.ScanInfo_t{}
	apInfo := make([]JM.ApInfo_t , 0)
	for i:=0;i<len(value.Wifis);i++ {
		ap := JM.ApInfo_t{Channel : 6}
		ap.ApMac = value.Wifis[i].Bssid
		ap.Rssi = value.Wifis[i].Dbm
		ap.Ssid = JM.GetMd5String(value.Wifis[i].Bssid)
		apInfo = append(apInfo , ap)
	}

	scanInfo.ApList= apInfo
	scanInfo.LocationTime = time.Now().UnixNano()
	scanInfo.TerminalMac = "78:02:f8:3e:e5:77"
	scanInfo.TerminalImei = "861414033593558"
	scanInfo.Product = "mi5"
	scanInfo.Brand = "xiaomi"
	scanInfo.Manufacturer = "xiaomi"
	scanInfo.LocationId = "23werdsdf3i43uh3u34h"
	scanInfo.AppId = JM.APPID
	scanInfo.AccessToken = t

	resBean, _ := json.Marshal(&scanInfo)
	fmt.Println(string(resBean))

	res := JM.HttpPost(JM.URL+JM.SCANSUBMIT_PATH , string(resBean))
	result := JM.SubmitResultInfo_t{}
	json.Unmarshal([]byte(res) , &result)
	if result.Code == 0 && result.Msg == "" {
		fmt.Println(result);
		query := buildPoiLog(result.Data)
		info, _ :=db.JmPoiLog.Upsert(query , bson.M{"$set":query})
		fmt.Println(info)
	}else{
		fmt.Println("scanSubmit failed!")
	}
	fmt.Println(time.Now())
	db.LocComputeLog.Upsert(bson.M{"_id":value.Id} , bson.M{"$set":bson.M{"updatedAt":time.Now()}})
	return result.Code
}

func buildPoiLog(s JM.ScanResultInfo_t) bson.M{
	return bson.M{
		"unitId":s.UnitId,
		"locationId": s.LocationId,
		"cityName":s.CityName,
		"zoneName":s.ZoneName,
		"buildingName":s.BuildingName,
		"unitName":s.UnitName,
		"floorName":s.FloorName,
		"categoryName":s.CatagoryName,
		"addtionalDesc":s.AddtionalDesc,
		"address":s.Address}
}

func cronJob(){
	bPass, token := getAuth()
	if !bPass {
		fmt.Println("get Token Failed!")
	}
	fmt.Println(token)

	value := JM.LocComputeLog_t{}
	//err := false
	iter := db.LocComputeLog.Find(nil).Iter()
	for iter.Next(&value) {
		err :=scanSubmit(token , &value)
		if err == -4 {
			for {
				bPass, token = getAuth()
				if bPass {
					err = scanSubmit(token , &value)
					break;
				}
			}
		}
	}
	//2 , 接入定位提交
	//conSubmit(token)
	//3, 扫描定位接口
	//scanSubmit(token)
}

func main(){
	connectDB()
	cronJob()
	//c := cron.New()
	//spec := "*/1 * * * * ?"
	//c.AddFunc(spec , cronJob)

	//c.Start()

	//select{}

	disConnectDB()
}