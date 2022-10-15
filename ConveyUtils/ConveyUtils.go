package ConveyUtils

import (
	"fmt"
	"net"
	"strings"
)

func ConveyDataToService(conn net.Conn, data []byte) {
	n, err2 := conn.Write(data)
	if err2 != nil {
		fmt.Println("Write err is", err2)
	}
	fmt.Printf("向服务端发送了[ %d ]个字节的数据\n", n)
}

func RecvFromClient(conn net.Conn) (data []byte) {
	data = make([]byte, 0)
	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务端在等待[%v]的输入\n", conn.RemoteAddr())
		n, err1 := conn.Read(buf) //阻塞，直到客户端发送消息
		buf = buf[:n]
		bufStr := string(buf)
		if strings.Contains(bufStr, "_over") {
			fmt.Println("_over is", buf)
			data = append(data, buf...)
			break
		}
		if err1 != nil {
			fmt.Println("服务端Read err is", err1)
			break
		}
		data = append(data, buf...)
		//fmt.Print(string(buf[:n])) // n is real data read from conn
	}
	data = data[:len(data)-5]
	fmt.Println("Recv all:", data)
	return
}

func RecvOver(conn net.Conn) {
	for {
		buf := make([]byte, 10)
		n, err := conn.Read(buf) //阻塞，直到客户端发送消息
		if err != nil {
			fmt.Println("Client read from server err:", err)
			return
		}
		fmt.Printf("接收到[ %d ]字节的数据\n", n)
		if strings.Contains(string(buf), "_over") {
			break
		}
	}
}
