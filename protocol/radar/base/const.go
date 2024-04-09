package base

const (
	OpHeartbeat      = 7
	OpHeartbeatReply = 8
	OpQuery          = 1
	OpQueryReply     = 3
	OpSet            = 2
	OpSetReply       = 4
	OpSubscribe      = 5
	OpReport         = 6
	OpFile           = 9
	OpUpgrade        = 11
	OpUpgradeReply   = 12
	OpCanSend        = 13
	OpCanRecv        = 14
	OpRs485Send      = 15
	OpRs485Recv      = 16
)

var OperationDesc = map[int]string{
	OpQuery:          "查询",
	OpSet:            "设置",
	OpQueryReply:     "查询应答",
	OpSetReply:       "设置应答",
	OpSubscribe:      "订阅",
	OpReport:         "主动上报",
	OpHeartbeat:      "心跳",
	OpHeartbeatReply: "心跳应答",
	OpFile:           "文件",
	OpUpgrade:        "固件刷新",
	OpUpgradeReply:   "固件刷新应答",
	OpCanSend:        "网口转CAN-发送",
	OpCanRecv:        "网口转CAN-接收",
	OpRs485Send:      "网口转RS485-发送",
	OpRs485Recv:      "网口转RS485-接收",
}

const (
	IdentHeartBeat     = 0
	IdentIPInfo        = 1
	IdentTime          = 2
	IdentComm          = 3
	IdentComm3rd       = 4
	IdentReboot        = 5
	IdentVersion       = 6
	IdentRadar         = 7
	IdentLane          = 8
	IdentLine          = 9
	IdentVLine         = 10
	IdentArea          = 11
	IdentEvent         = 12
	IdentCycle         = 13
	IdentReport        = 14
	IdentRadarGPS      = 16
	IdentSubs          = 30
	IdentPulse         = 31
	IdentQueue         = 32
	IdentLive          = 33
	IdentInstant       = 34
	IdentPass          = 35
	IdentRoad          = 36
	IdentStat          = 37
	IdentEval          = 38
	IdentEvent2        = 39
	IdentFault         = 40
	IdentMode          = 128
	IdentApp           = 129
	IdentFilter        = 130
	IdentTrigger       = 131
	IdentFile          = 132
	IdentData          = 133
	IdentTrack         = 145
	IdentTriggerStatus = 146
	IdentDesc          = 147
	IdentBootloader    = 148
	IdentCan           = 149
	IdentStore         = 150
	IdentGyro          = 151
	IdentGyroCal       = 152
	IdentRs485         = 153
	IdentRaw           = 255
)
