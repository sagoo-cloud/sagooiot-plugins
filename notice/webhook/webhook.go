package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
	gplugin "github.com/hashicorp/go-plugin"
	"net/rpc"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	extend "sagooiot/pkg/plugins/module"
	"sagooiot/pkg/plugins/sdk"
)

type Options struct {
	From       string
	PayloadURL string
	Secret     string
	Subject    string
	Body       interface{}
}

// NoticeWebhook 实现
type NoticeWebhook struct{}

func (NoticeWebhook) Info() model.PluginInfo {
	var res = model.PluginInfo{}
	res.Types = PluginType.Notice
	res.Name = "webhook"
	res.Title = "WebHook触发通知"
	res.Author = "Microrain"
	res.Description = "通过Webhook方式发送通知"
	res.Version = "0.01"
	return res
}

func (NoticeWebhook) Send(data []byte) (res model.JsonRes) {
	//解析通知数据
	nd, err := sdk.DecodeNoticeData(data)
	if err != nil {
		res.Code = 2
		res.Message = "插件数据解析失败"
		res.Data = err.Error()
		return res
	}

	var sendObjectList []model.NoticeSendObject
	err = gjson.DecodeTo(nd.Msg.Totag, &sendObjectList)

	//发送的信息内容采用|线进行内容分割
	var touser string
	if len(sendObjectList) > 1 {
		for _, object := range sendObjectList {
			if object.Name == "webhook" {
				touser = touser + object.Value + "|"
			}
		}
	} else {
		touser = "@all"
	}

	weConfigs := gconv.Interfaces(nd.Config["webhook"])
	var resData = make(map[string]interface{})
	for _, opData := range weConfigs {
		op := new(Options)
		if err := gconv.Struct(opData, op); err != nil {
			g.Log().Error(context.TODO(), err)
		}
		if op.From == "" {
			op.From = "default"
		}
		switch op.From {
		case "wework":
			op.Subject = nd.Msg.MsgTitle
			var sendData = MessageData{}
			sendData.Touser = touser
			sendData.Msgtype = "text"
			sendData.Agentid = nd.Msg.TemplateCode
			sendData.Text.Content = nd.Msg.MsgTitle + "\n\n" + nd.Msg.MsgBody
			sendData.Safe = 0
			sendData.EnableIdTrans = 0
			sendData.EnableDuplicateCheck = 0
			sendData.DuplicateCheckInterval = 1800
			op.Body = sendData
			op.PayloadURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + op.Secret
		case "default":
			op.Subject = nd.Msg.MsgTitle
			op.Body = nd.Msg.MsgBody
		}
		resData[op.From] = PostData(op)
	}
	res.Code = 0
	res.Message = "webhook本地发送完成"
	res.Data = gconv.String(resData)
	return res

}

// PostData 通过API修改数据
func PostData(o *Options) (res string) {
	c := g.Client()
	c.SetHeader("Secret", o.Secret)
	c.SetHeader("Accept", "application/json")
	c.SetHeader("Content-Type", "application/json")
	if r, e := c.Post(context.TODO(), o.PayloadURL, o.Body); e != nil {
		g.Log().Error(context.TODO(), e)
		return e.Error()

	} else {
		defer func(r *gclient.Response) {
			if err := r.Close(); err != nil {
				fmt.Println(err)
			}
		}(r)
		g.Log().Debug(context.TODO(), o.PayloadURL, r.StatusCode)
		//body := []byte(r.ReadAllString())
		//g.Log().Debug(context.TODO(), body)
		return r.ReadAllString()
	}
}

// WebhookPlugin 插件接口实现
type WebhookPlugin struct{}

// Server 此方法由插件进程延迟调
func (WebhookPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &extend.NoticeRPCServer{Impl: new(NoticeWebhook)}, nil
}

// Client 此方法由宿主进程调用
func (WebhookPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &extend.NoticeRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: extend.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"webhook": new(WebhookPlugin),
}
