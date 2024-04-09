package alisms

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot-plugin/notice/sms/provider"
	"sagooiot/pkg/plugins/model"
)

type Instance struct {
}

func (i *Instance) SendSms(ctx *provider.Context, msg model.NoticeInfoData) (result model.JsonRes, err error) {

	smsConfig := ctx.SmsConfig
	g.Log().Debug(context.TODO(), "sms发送开始")
	keyId := gconv.String(smsConfig["AccessKeyId"])
	secret := gconv.String(smsConfig["AccessSecret"])
	regionId := gconv.String(smsConfig["RegionId"])
	signName := gconv.String(smsConfig["SignName"]) //短信签名

	//发送的信息内容采用|线进行内容分割
	var phoneNumbers []string
	for _, object := range msg.Totag {
		if object.Name == "sms" {
			phoneNumbers = append(phoneNumbers, object.Value)
		}
	}

	if phoneNumbers == nil || len(phoneNumbers) == 0 {
		result.Code = 1
		result.Message = "手机号码不能为空"
		return
	}

	res, err := New(regionId, keyId, secret, signName).
		Request(msg.TemplateCode, msg.MsgBody, phoneNumbers)

	if err != nil {
		result.Code = 1
		result.Message = err.Error()
		return
	}
	result.Code = 0
	result.Message = res

	return
}
