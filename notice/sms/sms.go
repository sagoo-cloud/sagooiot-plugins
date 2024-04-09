package main

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	gplugin "github.com/hashicorp/go-plugin"
	"net/rpc"
	"sagooiot-plugin/notice/sms/provider"
	"sagooiot-plugin/notice/sms/provider/alisms"
	"sagooiot-plugin/notice/sms/provider/tencentcloud"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	extend "sagooiot/pkg/plugins/module"
	"sagooiot/pkg/plugins/sdk"
)

// NoticeSms 实现
type NoticeSms struct{}

func (NoticeSms) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Types = PluginType.Notice
	res.Name = "sms"
	res.Title = "短信通知"
	res.Author = "Microrain"
	res.Description = "通过短信发送通知"
	res.Version = "0.01"
	return res
}

func (NoticeSms) Send(data []byte) (res model.JsonRes) {
	//解析通知数据
	nd, err := sdk.DecodeNoticeData(data)
	if err != nil {
		res.Code = 2
		res.Message = "短信插件数据解析失败"
		res.Data = err.Error()
		return res
	}

	title := gconv.String(nd.Config["title"])
	//初始化上下文
	ctx := &provider.Context{
		ProviderName:  gconv.String(nd.Config["ProviderName"]),
		ProviderTitle: title,
		SmsConfig:     nd.Config,
		SendParam:     nd.SendParam,
	}
	res, err = SmsData(ctx, nd.Msg)
	res.Data = gconv.String(nd)

	return
}

// SmsData 短信发送
func SmsData(ctx *provider.Context, msg model.NoticeInfoData) (result model.JsonRes, err error) {
	var instance provider.SmsProviderInterface
	switch ctx.ProviderName {
	case "alisms":
		instance = &alisms.Instance{}
	case "tencentcloud":
		instance = &tencentcloud.Instance{}

	default:
		err = gerror.New("未选择短信发送供应商")
		return
	}

	g.Log().Debug(context.TODO(), "发达短信供应商：", ctx.ProviderTitle)
	result, err = instance.SendSms(ctx, msg)

	return

}

// SmsPlugin 插件接口实现
type SmsPlugin struct{}

// Server 此方法由插件进程延迟调
func (SmsPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &extend.NoticeRPCServer{Impl: new(NoticeSms)}, nil
}

// Client 此方法由宿主进程调用
func (SmsPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &extend.NoticeRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: extend.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"sms": new(SmsPlugin),
}
