package main

import (
	"sagooiot-plugin/notice/sms/provider/alisms"
	"testing"
)

// 测试阿里云短信发送 测试的时候，注意确认keyId与secret是否正确
func TestSender_Request(t *testing.T) {
	var sms = alisms.New("cn-hangzhou", "LTAI5tHLKfkUjspf7M", "dnNGFYpF3ROdV6Oi8vurPqV3x3pS", "sagoo")
	result, err := sms.Request("SMS_234795741", "{'code':'19001'}", []string{"13700005102"})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(result)
}

// 测试短信插件发送 测试的时候，注意确认keyId与secret是否正确
func TestSmsData(t *testing.T) {
	var nSms = new(NoticeSms)
	sendBody := `{"Config":{"AccessKeyId":"LTAI5tHLKfkUjspf7M","AccessSecret":"dnNGFYpF3ROdV6Oi8vurPqV3x3pS","ProviderName":"alisms","RegionId":"cn-hangzhou","SignName":"sagoo","Title":"阿里云"},"Msg":{"come_from":"","config_id":"","method":"","method_cron":"","method_num":0,"msg_body":"{'device':'19001','msg':'aaaaaaaaaa'}","msg_title":"title111112222","msg_url":"","party_ids":"","totag":"[{\"name\":\"mail\",\"value\":\"940290@qq.com\"},{\"name\":\"webhook\",\"value\":\"cccc\"},{\"name\":\"sms\",\"value\":\"13700005102\"}]","user_ids":""},"SendParam":{"code":"SMS_464035361"}}`
	res := nSms.Send([]byte(sendBody))
	t.Log(res)
}
