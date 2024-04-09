package main

import "testing"

func TestMailData(t *testing.T) {
	var mail = new(NoticeMail)
	sendBody := `{"Config":{"MailHost":"smtp.qq.com","MailPort":"465","MailUser":"cyw.fly@qq.com","MailPass":"hhipgbgmkichbgcg"},"Msg":{"come_from":"","config_id":"","method":"","method_cron":"","method_num":0,"msg_body":"{'device':'19001','msg':'aaaaaaaaaa'}","msg_title":"title111112222","msg_url":"","party_ids":"","totag":[{"name":"mail","value":"xinjy@qq.com"}],"user_ids":""},"SendParam":{"code":"SMS_464035361"}}`
	mail.Send([]byte(sendBody))
}
