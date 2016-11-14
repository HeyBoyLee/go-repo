package JM


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

type ApInfo struct {
	ApMac string			`json:"ApMAC"`
	Channel int8			`json:"channel"`
	Ssid string			`json:"ssid"`
	Rssi int8			`json:"rssi"`
}

type TerminalInfo struct {
	ApList []ApInfo			`json:"scanInfo"`
	LocationTime int64		`json:"locationTime"`
	TerminalMac string		`json:"terminalMAC"`
	TerminalImei string		`json:"terminalImei"`
	Product string			`json:"product"`
	Brand string			`json:"brand"`
	Manufacturer string		`json:"manufacturer"`
	LocationId string		`json:"locationId"`
}

type ScanInfo struct {
	TerminalList []TerminalInfo	`json:"scanInfoList"`
	AccessToken string		`json:"accessToken"`
	AppId string			`json:"appId"`
}