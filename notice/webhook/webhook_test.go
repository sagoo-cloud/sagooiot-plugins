package main

import "testing"

func TestNoticeWebhook_Send(t *testing.T) {
	var nw = new(NoticeWebhook)
	//sendBody := `{"Config":{"webhook":[{"PayloadURL":"https://qyapi.weixin.qq.com/cgi-bin/webhook/send","Secret":"4f832276-e04e-47ba-8e1b-913811522fd2","From":"wework"}]},"Msg":{"come_from":"","config_id":"","method":"","method_cron":"","method_num":0,"msg_body":"{'device':'19001','msg':'aaaaaaaaaa'}","msg_title":"title111112222","msg_url":"","party_ids":"","totag":"[{\"name\":\"mail\",\"value\":\"940290@qq.com\"},{\"name\":\"webhook\",\"value\":\"cccc\"},{\"name\":\"sms\",\"value\":\"13700005102\"},{\"name\":\"webhook\",\"value\":\"13700005102\"}]","user_ids":""},"SendParam":{"code":"SMS_464035361"}}`
	sendBody := `{"Config":{"webhook":[{"PayloadURL":"https://qyapi.weixin.qq.com/cgi-bin/webhook/send","Secret":"4f832276-e04e-47ba-8e1b-913811522fd2","From":"wework"}]},"Msg":{"come_from":"","config_id":"","method":"","method_cron":"","method_num":0,"msg_body":"{'code':'19001'}","msg_title":"title111112222","msg_url":"","party_ids":"","template_code":"SMS_464050874","totag":[{"name":"webhook","value":"@all"}],"user_ids":""},"SendParam":{}}`
	res := nw.Send([]byte(sendBody))
	t.Log(res)
}
