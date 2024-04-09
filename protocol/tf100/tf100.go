package main

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/guid"
	gplugin "github.com/hashicorp/go-plugin"
	"net/rpc"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/plugins/consts/PluginHandleType"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	plugin "sagooiot/pkg/plugins/module"
	"time"
)

// ProtocolTF100 实现
type ProtocolTF100 struct{}

func (p *ProtocolTF100) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Types = PluginType.Protocol
	res.HandleType = PluginHandleType.TcpServer
	res.Name = "tf100"
	res.Title = "TF100 设备协议"
	res.Author = "Microrain"
	res.Description = "TF100设备与服务器通信协议"
	res.Version = "0.01"
	return res
}

func (p *ProtocolTF100) Encode(args model.DataReq) model.JsonRes {
	var resp model.JsonRes
	resp.Code = 0
	resp.Message = string(args.Data)
	resp.Data = args.Data
	fmt.Println("接收到参数：", args)
	return resp
}

func (p *ProtocolTF100) Decode(data model.DataReq) model.JsonRes {
	var resp model.JsonRes
	resp.Code = 0
	bodyDataLen := data.Data[5:8]

	bodyData := data.Data[8:]

	var msg MsgObj
	err := json.Unmarshal(bodyData, &msg)
	if err != nil {
		resp.Code = 1
		resp.Message = fmt.Sprintf("json字串：%s 解析失败：%s", string(bodyData), err.Error())
		return resp
	}
	fmt.Println(msg.MsgId, msg.MsgType)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	shanghaiTime := time.Now().In(loc)
	timestamp := shanghaiTime.Unix()
	var rd = make(map[string]interface{})
	//设置属性
	rd["msgType"] = sagooProtocol.PropertyNode{Value: msg.MsgType, CreateTime: timestamp}
	rd["devId"] = sagooProtocol.PropertyNode{Value: msg.DevId, CreateTime: timestamp}
	rd["msgId"] = sagooProtocol.PropertyNode{Value: msg.MsgId, CreateTime: timestamp}
	rd["ports"] = sagooProtocol.PropertyNode{Value: msg.Ports, CreateTime: timestamp}
	rd["errTypes"] = sagooProtocol.PropertyNode{Value: 0, CreateTime: timestamp}

	resp.Code = 0

	resp.Data = sagooProtocol.ReportPropertyReq{
		Id:      guid.S(),
		Version: string(bodyDataLen), // "1.0",
		Sys:     sagooProtocol.SysInfo{Ack: 0},
		Params:  rd,
		Method:  "thing.event.property.post",
	}
	return resp
}

// TF100Plugin 插件接口实现
// 这有两种方法：服务器必须为此插件返回RPC服务器类型。我们为此构建了一个RPCServer。
// 客户端必须返回我们的接口的实现通过RPC客户端。我们为此返回RPC。
type TF100Plugin struct{}

// Server 此方法由插件进程延迟调
func (t *TF100Plugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &plugin.ProtocolRPCServer{Impl: new(ProtocolTF100)}, nil
}

// Client 此方法由宿主进程调用
func (t *TF100Plugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &plugin.ProtocolRPC{Client: c}, nil
}

// 主函数启动插件服务
func main() {
	//调用plugin.Serve()启动侦听，并提供服务
	//ServeConfig 握手配置，插件进程和宿主机进程，都需要保持一致
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig,
		Plugins: map[string]gplugin.Plugin{
			"tf100": &TF100Plugin{},
		},
	})
}
