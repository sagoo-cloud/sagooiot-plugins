package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	gplugin "github.com/hashicorp/go-plugin"
	"net/rpc"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/plugins/consts/PluginHandleType"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	plugin "sagooiot/pkg/plugins/module"
	"strings"
)

// ProtocolTgn52 实现
type ProtocolTgn52 struct{}

func (p *ProtocolTgn52) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Name = "tgn52"
	res.Types = PluginType.Notice
	res.HandleType = PluginHandleType.TcpServer
	res.Title = "TG-N5 v2设备协议"
	res.Author = "Microrain"
	res.Description = "对TG-N5插座设备进行数据采集v2"
	res.Version = "0.01"
	return res
}

func (p *ProtocolTgn52) Encode(args model.DataReq) model.JsonRes {
	var resp model.JsonRes
	fmt.Println("接收到参数：", args)
	return resp
}

func (p *ProtocolTgn52) Decode(data model.DataReq) model.JsonRes {
	var resp model.JsonRes
	resp.Code = 0

	tmpData := strings.Split(string(data.Data), ";")
	var rd = make(map[string]interface{})

	l := len(tmpData)
	nowTime := gtime.Now().Unix()
	if l > 7 {
		rd["HeadStr"] = sagooProtocol.PropertyNode{Value: tmpData[0], CreateTime: nowTime}
		rd["DeviceID"] = sagooProtocol.PropertyNode{Value: tmpData[1], CreateTime: nowTime}
		rd["Signal"] = sagooProtocol.PropertyNode{Value: tmpData[2], CreateTime: nowTime}
		rd["Battery"] = sagooProtocol.PropertyNode{Value: tmpData[3], CreateTime: nowTime}
		rd["Temperature"] = sagooProtocol.PropertyNode{Value: tmpData[4], CreateTime: nowTime}
		rd["Humidity"] = sagooProtocol.PropertyNode{Value: tmpData[5], CreateTime: nowTime}
		rd["Cycle"] = sagooProtocol.PropertyNode{Value: tmpData[6], CreateTime: nowTime}
		//处理续传数据
		updateStr := make([]string, 0)
		for i := 7; i < l; i++ {
			updateStr = append(updateStr, tmpData[i])
		}
		rd["Update"] = sagooProtocol.PropertyNode{Value: updateStr, CreateTime: nowTime}
	}

	resp.Code = 0
	resp.Data = sagooProtocol.ReportPropertyReq{
		Id:      guid.S(),
		Version: "1.0",
		Sys:     sagooProtocol.SysInfo{Ack: 0},
		Params:  rd,
		Method:  "thing.event.property.post",
	}
	return resp
}

// Tgn52Plugin 插件接口实现
// 这有两种方法：服务器必须为此插件返回RPC服务器类型。我们为此构建了一个RPCServer。
// 客户端必须返回我们的接口的实现通过RPC客户端。我们为此返回RPC。
type Tgn52Plugin struct{}

// Server 此方法由插件进程延迟调
func (t *Tgn52Plugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &plugin.ProtocolRPCServer{Impl: new(ProtocolTgn52)}, nil
}

// Client 此方法由宿主进程调用
func (t *Tgn52Plugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &plugin.ProtocolRPC{Client: c}, nil
}

func main() {
	//调用plugin.Serve()启动侦听，并提供服务
	//ServeConfig 握手配置，插件进程和宿主机进程，都需要保持一致
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

// 插件进程必须指定Impl，此处赋值为greeter对象
var pluginMap = map[string]gplugin.Plugin{
	"tgn52": new(Tgn52Plugin),
}
