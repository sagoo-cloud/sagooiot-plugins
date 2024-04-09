package tencentcloud

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot-plugin/notice/sms/provider"
	"sagooiot/pkg/plugins/model"
)

type Instance struct {
}

func (i *Instance) SendSms(ctx *provider.Context, msg model.NoticeInfoData) (result model.JsonRes, err error) {
	smsConfig := ctx.SmsConfig
	g.Log().Debug(context.TODO(), "sms发送开始")
	secretKey := gconv.String(smsConfig["AccessSecret"])
	secretId := gconv.String(smsConfig["AccessKeyId"])
	signName := gconv.String(smsConfig["SignName"]) //短信签名

	//发送的信息内容采用|线进行内容分割
	TemplateParam := gstr.Explode("|", msg.MsgBody)

	var phoneNumbers []string
	for _, object := range msg.Totag {
		if object.Name == "sms" {
			phoneNumbers = append(phoneNumbers, object.Value)
		}
	}

	res, err := New(secretId, secretKey, signName).
		Request(msg.TemplateCode, TemplateParam, phoneNumbers)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		result.Code = 1
		result.Message = err.Error()
		return
	}
	result.Code = 0
	result.Message = res
	return
}
