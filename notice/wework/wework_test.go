package main

import "testing"

func TestSendData(t *testing.T) {
	var we = new(NoticeWework)
	sendBody := `{"Config":{"AgentID":1000006,"Corpid":"ww4863d34f18330c5e","EncodingAESKey":"IaiXRoS2mk9e2cI4vpPsvSrVbkxpM578tOnpbpZkTmo","Secret":"nM_rK-hCpsgQvYkfEH1AsrqYUN86yo6czsrg6q9cXEg","Token":"p6uzoqddwX4B84SW2OUDb7msk45GWWi"},"Msg":{"come_from":"","config_id":"","method":"","method_cron":"","method_num":0,"msg_body":"你好，你的系统有如下告警：\n产品：模拟测试电表2022\n设备：\n级别：超紧急 \n触发规则：测试 (va \u003e 200)","msg_title":"告警通知模版","msg_url":"","party_ids":"","template_code":"m002","totag":[{"name":"wework","value":"xjy@sagoo.cn"}],"user_ids":""},"SendParam":{}}`
	we.Send([]byte(sendBody))

}
