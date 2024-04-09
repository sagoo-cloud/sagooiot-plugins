package main

import (
	"fmt"
	"net/rpc"
	"sagooiot/pkg/plugins/consts/PluginHandleType"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"

	gplugin "github.com/hashicorp/go-plugin"
	plugin "sagooiot/pkg/plugins/module"
)

// ProtocolModbus 实现
type ProtocolModbus struct{}

func (ProtocolModbus) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Name = "modbus"
	res.Types = PluginType.Protocol
	res.HandleType = PluginHandleType.TcpServer
	res.Title = "Modbus TCP协议"
	res.Author = "Microrain"
	res.Description = "对modbus TCP模式的设备进行数据采集"
	res.Version = "0.01"
	return res
}

func (ProtocolModbus) Encode(args model.DataReq) model.JsonRes {
	var resp model.JsonRes
	fmt.Println("接收到参数：", args)
	return resp
}

func (ProtocolModbus) Decode(data model.DataReq) model.JsonRes {
	var resp model.JsonRes
	resp.Code = 0
	resp.Data = data.Data
	return resp

}

// ModbusPlugin 插件接口实现
// 这有两种方法：服务器必须为此插件返回RPC服务器类型。我们为此构建了一个RPCServer。
// 客户端必须返回我们的接口的实现通过RPC客户端。我们为此返回RPC。
type ModbusPlugin struct{}

// Server 此方法由插件进程延迟调
func (ModbusPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &plugin.ProtocolRPCServer{Impl: new(ProtocolModbus)}, nil
}

// Client 此方法由宿主进程调用
func (ModbusPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
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
	"modbus": new(ModbusPlugin),
}
