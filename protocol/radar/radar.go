package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"net/rpc"
	"sagooiot-plugin/protocol/radar/base"
	"sagooiot/pkg/plugins/consts/PluginHandleType"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	"strconv"

	gplugin "github.com/hashicorp/go-plugin"
	plugin "sagooiot/pkg/plugins/module"
)

// radr 实现
type ProtocolRadar struct{}

func (ProtocolRadar) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Name = "radar"
	res.Types = PluginType.Protocol
	res.HandleType = PluginHandleType.TcpServer
	res.Title = "雷达协议"
	res.Author = "Fade"
	res.Description = "微波雷达检测器进行数据采集"
	res.Version = "0.01"
	return res
}

func (ProtocolRadar) Encode(args model.DataReq) model.JsonRes {
	var resp model.JsonRes
	resp.Code = 0

	var packet base.Packet
	err := gjson.Unmarshal(args.Data, packet)
	if err != nil {
		resp.Code = 1
		resp.Message = "参数错误,不是base.Packet类型"
		return resp
	}

	//packet, ok := args.(base.Packet)
	//if !ok {
	//	resp.Code = 1
	//	resp.Message = "参数错误,不是base.Packet类型"
	//	return resp
	//}
	encodeData, err := packet.Encode()
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return resp
	}
	resp.Data = string(encodeData)
	return resp
}

func (ProtocolRadar) Decode(data model.DataReq) model.JsonRes {
	var resp model.JsonRes
	p, err := base.Decode(data.Data)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return resp
	}
	dataIdentType, err := strconv.Atoi(data.DataIdent)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		return resp
	}
	var action *base.CommonOp
	switch dataIdentType {
	case base.IdentHeartBeat:
		action = &base.ActionHeartBeat
	case base.IdentIPInfo:
		action = &base.ActionQueryIpInfo
	case base.IdentVersion:
		action = &base.ActionQueryVersionInfo
	case base.IdentRadar:
		action = &base.ActionQueryRadarInfo
	case base.IdentLane:
		action = &base.ActionQueryRadarInfo
	}
	if action == nil {
		resp.Code = 1
		resp.Message = fmt.Sprintf("未知的数据标识:%s", data.DataIdent)
		return resp
	}
	if int(p.Op) == action.OpTypeResponse {
		fieldMap, err := action.Convert(p.Payload)
		if err != nil {
			resp.Code = 1
			resp.Message = err.Error()
			return resp
		}
		respByte, _ := json.Marshal(fieldMap)
		resp.Data = string(respByte)
		return resp
	} else {
		resp.Code = 1
		resp.Message = fmt.Sprintf("接收到错误数据:%x 标识:%s", data.Data, data.DataIdent)
		return resp
	}

}

type RadarPlugin struct{}

func (RadarPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &plugin.ProtocolRPCServer{Impl: new(ProtocolRadar)}, nil
}

func (RadarPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &plugin.ProtocolRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"radar": new(RadarPlugin),
}
