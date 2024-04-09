package test

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func TestTCPClient(t *testing.T) {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:5090")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			t.Error(err)
		}
	}(conn)

	fmt.Println("Connecting to " + conn.RemoteAddr().String())

	jsonText := `{"msgType": 20,"devId": 13900001,"msgId": 5,"mainVer": 101,"powVer": [101, 101],"ports": 2,"vsId": 0,"devType": 1,"peType": 0,"enUpd": 1}`
	msgData := createMessage(jsonText)
	//缓存conn中的数据
	//buf := make([]byte, 1024*4)
	for {
		// 发送数据
		fmt.Println(msgData)
		_, err = conn.Write(msgData)
		if err != nil {
			fmt.Println("Error sending:", err)
			os.Exit(1)
		}
		message := make(chan string)
		go receive(conn, message)

		//// 接收数据：接收服务端返回的数据存入buf
		//n, err := conn.Read(buf)
		//if err != nil {
		//	fmt.Println("c read failed err: ", err)
		//	return
		//}
		//// 显示服务端回传的数据
		//fmt.Println("服务端返回:", string(buf[:n]))
		time.Sleep(2 * time.Second)
	}

}

// 封包
func createMessage(content string) []byte {
	contentWrapped := fmt.Sprintf("%s", content)     //content
	size := fmt.Sprintf("%03d", len(contentWrapped)) //
	message := fmt.Sprintf("CCMD:%s%s", size, contentWrapped)
	return []byte(message)
}

// 接收事件
func receive(conn net.Conn, messages chan string) {

	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
		messages <- "Error reading: " + err.Error()
		return
	}

	fmt.Println("==========服务器回复：==", reply)

}
