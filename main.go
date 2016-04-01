// go_tcp_client project main.go
package main

import (
	"fmt"
	"net"
)

//ip和port常量的定义
const (
	serverip   = "localhost"
	serverport = "54321"
)

var (
	maxRead = 1100
	msgSend = "{\"hello\":\"world\", \"t\":true, \"f\":false, \"n\":null, \"i\":123, \"pi\":3.1416, \"a\":[1,2,3,4]}"
)

func main() {
	hostAndPort := serverip + ":" + serverport
	conn, err := net.Dial("tcp", hostAndPort)
	checkError(err, "Dial:")
	fmt.Println("连接成功...")
	fmt.Println("发送数据:" + msgSend)
	conn.Write([]byte(msgSend))
	var ibuf []byte = make([]byte, maxRead+1)
	length, err := conn.Read(ibuf[0:maxRead])
	checkError(err, "Read:")
	ibuf[maxRead] = 0 //防止overflow
	fmt.Printf("接收数据: %v\n", string(ibuf[0:length]))
}

//错误检查，
func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error()) //终止程序
	}
}
