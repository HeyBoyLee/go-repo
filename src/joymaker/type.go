package JM

import (
	//"go/token"
	//"crypto"
)

type AuthInfo_t struct {
	AppId   string		`json:"appid"`
	Sign    string		`json:"sign"`
	Version string    `json:"version"`
}

type AuthResultInfo_t struct {
	Token		string		`json:"token"`
	//Result 	bool			`json:"result"`
	//MessageCode	int8	`json:"messageCode"`
}

type CommonResultInfo_t struct {
	Code	int8					`json:"code"`
	Msg 	string				`json:"msg"`
	//Data 	interface{}		`json:"data"`
	Data 	AuthResultInfo_t		`json:"data"`
}

//type CommonResultTest_t struct {
//	Code	int8					`json:"code"`
//	Msg 	string				`json:"msg"`
//	Data 	interface{}		`json:"data"`
//	//Data 	AuthResultInfo_t		`json:"data"`
//}
//
//type AuthResultData interface {
//	GetAuthData() string
//}
//
//type SubmitResultData interface {
//	GetSubmitData() ScanResultInfo_t
//}
//
//func (authResult *AuthResultInfo_t) GetAuthData() string{
//	return authResult.Token
//}

type ConnectInfo_t struct {
	ApMac string										`json:"apMAC"`
	Channel int8										`json:"channel"`
	Ssid string											`json:"ssid"`
	LocationTime int64							`json:"locationTime"`
	TerminalMac string							`json:"terminalMAC"`
	TerminalImei string							`json:"terminalIMEI"`
	LocationId string								`json:"locationId"`
}

type ConSubmitInfo_t struct{
	ConnectInfoList []ConnectInfo_t	`json:"connectInfoList"`
	AccessToken string							`json:"accessToken"`
	AppId string										`json:"appId"`
}

type ApInfo_t struct {
	ApMac string										`json:"apMAC"`
	Channel int8										`json:"channel"`
	Ssid string											`json:"ssid"`
	Rssi int												`json:"rssi"`
}

type TerminalInfo_t struct {
	ApList []ApInfo_t								`json:"scanInfo"`
	LocationTime int64							`json:"locationTime"`
	TerminalMac string							`json:"terminalMAC"`
	TerminalImei string							`json:"terminalImei"`
	Product string									`json:"product"`
	Brand string										`json:"brand"`
	Manufacturer string							`json:"manufacturer"`
	LocationId string								`json:"locationId"`
}

type ScanInfo_t struct {
	//TerminalList []TerminalInfo_t		`json:"scanInfoList"`
	LocationTime int64							`json:"locationTime"`
	TerminalMac string							`json:"terminalMAC"`
	TerminalImei string							`json:"terminalIMEI"`
	Product string									`json:"product"`
	Brand string										`json:"brand"`
	Manufacturer string							`json:"manufacturer"`
	LocationId string								`json:"locationId"`
	AccessToken string							`json:"accessToken"`
	AppId string										`json:"appId"`
	ApList []ApInfo_t								`json:"scanInfo"`
}

type ScanResultInfo_t struct {
	UnitId 				string		`json:"unitId"`
	LocationId 		string		`json:"locationId"`
	CityName 			string		`json:"cityName"`
	ZoneName 			string		`json:"zoneName"`
	BuildingName 	string		`json:"buildingName"`
	UnitName 			string		`json:"unitName"`
	FloorName 		string		`json:"floorName"`
	CatagoryName 	string		`json:"catagoryName"`
	AddtionalDesc string		`json:"addtionalDesc"`
	Address 			string		`json:"address"`
}

type SubmitResultInfo_t struct {
	Code	int					`json:"code"`
	Msg 	string				`json:"msg"`
	//Data 	interface{}		`json:"data"`
	Data 	ScanResultInfo_t	`json:"data"`
}


