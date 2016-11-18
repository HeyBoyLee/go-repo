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
	Result 	bool			`json:"result"`
	MessageCode	int8	`json:"messageCode"`
}

type CommonResultInfo_t struct {
	Result	bool			`json:"result"`
	MessageCode int8	`json:"messageCode"`
}

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
	ApMac string										`json:"ApMAC"`
	Channel int8										`json:"channel"`
	Ssid string											`json:"ssid"`
	Rssi int8												`json:"rssi"`
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
	TerminalList []TerminalInfo_t		`json:"scanInfoList"`
	AccessToken string							`json:"accessToken"`
	AppId string										`json:"appId"`
}