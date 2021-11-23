package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Sprintf("建立listen失败:%s", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Sprintf("建立accept失败:%s", err.Error())
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Sprintf("wirte消息失败:%s", err.Error())
		}
	}

}
