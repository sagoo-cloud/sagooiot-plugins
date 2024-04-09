package main

import "fmt"

func main() {
	data := createMessage(`{"msgType": 20,"devId": 13900001,"msgId": 5,"mainVer": 101,"powVer": [101, 101],"ports": 2,"vsId": 0,"devType": 1,"peType": 0,"enUpd": 1}`)
	fmt.Println("设备端发送的内容：", string(data))
}

func createMessage(content string) []byte {
	contentWrapped := fmt.Sprintf("%s", content)
	size := fmt.Sprintf("%03d", len(contentWrapped))
	message := fmt.Sprintf("CCMD:%s%s", size, contentWrapped)
	return []byte(message)
}
