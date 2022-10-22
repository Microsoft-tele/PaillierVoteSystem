package main

import (
	"SockGo/VoteUtils"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

var Can = VoteUtils.Candidate{}

func main() {

	Can.SetCandidateInfo()
	fmt.Println(Can)
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
	write, err := conn.Write([]byte("Candidate"))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}

	fmt.Println("[本地]开始与服务器通话：")

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
		if strings.Contains(str, "请上传您的信息") {
			PushInfoToSock(conn)
			fmt.Println("请等待投票结果：")
		}
	}
}
func PushInfoToSock(conn net.Conn) {
	fmt.Println("[本地]进入上传信息函数：")
	CanJson, err := json.Marshal(Can)
	if err != nil {
		fmt.Println("候选人对象序列化失败:", err)
		return
	}
	write, err := conn.Write(CanJson)
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}
	fmt.Printf("[本地]向[ %v ]发送[ %v ]字节数据\n[%v]\n", conn.RemoteAddr(), write, CanJson)
}
