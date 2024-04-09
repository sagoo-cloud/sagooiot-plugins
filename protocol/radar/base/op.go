package base

import (
	"fmt"
	"net"
	"strconv"
)

type CommonOp struct {
	OpTypeRequest  byte
	OpTypeResponse int
	Ident          byte
	Convert        func([]byte) (map[string]interface{}, error)
}

var ActionHeartBeat = CommonOp{OpHeartbeat, OpHeartbeatReply, IdentHeartBeat, func([]byte) (map[string]interface{}, error) {
	return map[string]interface{}{"alive": true}, nil
}}

var ActionQueryIpInfo = CommonOp{OpQuery, OpQueryReply, IdentIPInfo, func(data []byte) (map[string]interface{}, error) {
	if len(data) != 1+4+4+4+2+6 {
		return nil, fmt.Errorf("invalid data length")
	}
	ipAddr := net.IPv4(data[1], data[2], data[3], data[4]).String()
	ipMask := net.IPv4(data[5], data[6], data[7], data[8]).String()
	ipGate := net.IPv4(data[9], data[10], data[11], data[12]).String()
	ipPort := int(data[13])<<8 + int(data[14])
	macAddr := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", data[15], data[16], data[17], data[18], data[19], data[20])
	return map[string]interface{}{
		"ip":   ipAddr,
		"mask": ipMask,
		"gate": ipGate,
		"port": ipPort,
		"mac":  macAddr,
	}, nil
}}

var ActionQueryRadarInfo = CommonOp{OpQuery, OpQueryReply, IdentLane, func(data []byte) (map[string]interface{}, error) {
	//deviceId := int(data[1])<<8 + int(data[2])
	laneNum := int(data[3])
	laneInfo := make([]map[string]interface{}, laneNum)
	for i := 0; i < laneNum; i++ {
		laneInfo[i] = map[string]interface{}{
			"laneId": int(data[4+i*10]),
			"startX": int(data[5+i*10])<<8 + int(data[6+i*10]),
			"startY": int(data[7+i*10])<<8 + int(data[8+i*10]),
			"width":  int(data[9+i*10]),
			"length": int(data[10+i*10])<<8 + int(data[11+i*10]),
			"type":   int(data[12+i*10]),
			"flow":   int(data[13+i*10]),
		}
	}
	return map[string]interface{}{
		//"deviceId": deviceId,
		"laneNum":  laneNum,
		"laneInfo": laneInfo,
	}, nil
}}

var ActionQueryVersionInfo = CommonOp{OpQuery, OpQueryReply, IdentVersion, func(data []byte) (map[string]interface{}, error) {
	if len(data) < 1+128+25 {
		return nil, fmt.Errorf("invalid data length")
	}
	fmt.Println(1231222313123123)
	return map[string]interface{}{
		//"version": string(data[1 : 1+128]),
		"serial": string(data[1+128 : 1+128+25]),
		//"reserved": string(data[1+128+25 : 1+128+25+10]),
	}, nil
}}

var ActionQueryRadarInstallInfo = CommonOp{OpQuery, OpQueryReply, IdentRadar, func(data []byte) (map[string]interface{}, error) {
	if len(data) < 1+2+1+2+1 {
		return nil, fmt.Errorf("invalid data length")
	}
	return map[string]interface{}{
		//"deviceId":   int(data[1])<<8 + int(data[2]),
		"pitch":      int(data[3]),
		"pitchAngle": int(data[4]),
		"horizon":    int(data[5])<<8 + int(data[6]),
		"stopLine":   int(data[7])<<8 + int(data[8]),
		"direct":     strconv.Itoa(int(data[9])),
		"roadName":   string(data[10:]),
	}, nil
}}
