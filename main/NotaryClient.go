package main

import (
	"RemoteRouter/VoteUtils"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	Notary := VoteUtils.Notary{}
	Notary.InitNotary() // 创建对象
	fmt.Println(Notary)
LOOP:
	ServiceHostIp := "127.0.0.1"
	ServiceHostPort := "8888"
	address := ServiceHostIp + ":" + ServiceHostPort
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("建立连接失败：", err)
		goto LOOP
	}
	fmt.Println("建立连接成功：", conn.RemoteAddr())
	// 开始验证身份
	write, err := conn.Write([]byte("Notary"))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}

	fmt.Println("开始与服务器通话：")

	data := make([]byte, 0) // 进入对话模式
	for {
		buf := make([]byte, 1024)
		fmt.Printf("等待[ %v ]的输入\n", conn.RemoteAddr())
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("从[ %v ]接收失败 : err:[ %v ] : 接收的字节数:[ %v ]\n", conn.RemoteAddr(), err, read)
			return
		}
		buf = buf[:read]
		fmt.Printf("接收到的数据[byte]: [ %v ]\n", buf)
		data = append(data, buf...)

		str := string(buf)
		fmt.Printf("接收到的数据[char]: [ %v ]\n", str)
		if strings.Contains(str, "请确认所有候选人已经成功上传信息") {
			ScanlnToSock(conn)
			NotaryJson, err := json.Marshal(Notary)
			if err != nil {
				fmt.Println("Notary序列化失败：", err)
			}
			n, err := conn.Write(NotaryJson)
			if err != nil {
				fmt.Println(n)
				fmt.Printf("向[ %v ]发送数据失败\n", conn.RemoteAddr())
				return
			}
			fmt.Printf("向[ %v ]发送数据[ %v ]\n", conn.RemoteAddr(), NotaryJson)
		} else if strings.Contains(str, "请等待投票结果") {

		}
	}
}

func ScanlnToSock(conn net.Conn) {
	fmt.Println("请输入对应的数据:")
	sendData, _ := reader.ReadString('\n')
	_, err := conn.Write([]byte(sendData))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败:", conn.RemoteAddr())
		return
	}
}
