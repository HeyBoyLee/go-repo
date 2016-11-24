package JM

import (
	"gopkg.in/mgo.v2/bson"
	//"time"
	"time"
)

type Position_t struct {
	Lng float64					`lng`
	Lat float64					`lat`
}

type MixedRes_t struct {
	Position Position_t	`position`
	//Position map[string]float64
}

type Wifi_t struct {
	Id bson.ObjectId			`_id`
	Connected bool
	Accuracy int
	Dbm int
	Lat float64
	Lng float64
	Bssid string
}

type Cell_t struct {
	Id bson.ObjectId			`_id`
	Connected bool
	Accuracy int
	Lat float64
	Lng float64
	Dbm int
}

type LocComputeLog_t struct {
	Id string			`_id`
	//Role int8
	//Accuracy int
	//MixedRes MixedRes_t  `mixedRes`
	Wifis []Wifi_t
	Cells []Cell_t
	//Ts time.Time
	//Sg string
	//Uid string
	UpdatedAt time.Time
}