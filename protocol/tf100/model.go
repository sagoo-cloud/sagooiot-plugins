package main

type DeviceAlarm struct {
	MsgType int   `json:"msgType"`
	DevId   int   `json:"devId"`
	MsgId   int   `json:"msgId"`
	Ports   int   `json:"ports"`
	ErrType []int `json:"errType"`
}
type MsgObj struct {
	MsgType int   `json:"msgType"`
	DevId   int   `json:"devId"`
	MsgId   int   `json:"msgId"`
	MainVer int   `json:"mainVer"`
	PowVer  []int `json:"powVer"`
	Ports   int   `json:"ports"`
	VsId    int   `json:"vsId"`
	DevType int   `json:"devType"`
	PeType  int   `json:"peType"`
	EnUpd   int   `json:"enUpd"`
}
