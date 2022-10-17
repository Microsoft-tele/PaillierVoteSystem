package main

import (
	"SockGo/ConveyUtils"
	"SockGo/RSAUtils"
	"SockGo/ShellUtils"
	"SockGo/paillier"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var PubKey paillier.PublicKey

func main() {
	ServiceHostIp := "192.168.1.103"
	ServiceHostPort := "8888"
	address := ServiceHostIp + ":" + ServiceHostPort
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Client dial err =", err)
		return
	}
	fmt.Printf("连接成功：%v 服务器地址是：%v\n", conn, conn.RemoteAddr().String())
	//var data []byte

	//data = make([]byte, 1024)
	for {
		data := ConveyUtils.RecvStringFrom(conn)
		fmt.Println("服务端发来的数据:", string(data))
		if strings.Contains(data, "选择您的身份") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "输入您的基本信息以获得投票资格证") {
			ScanlnToSock(conn)
		} else if strings.Contains(data, "传您的RSA公钥以获得") {
			// 向服务器传输RSA公钥
			Case3(conn)

		} else if strings.Contains(data, "是否修改加入投票的总人数") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "请输入修改的数字") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "功能菜单") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "是否生成新的Paillier密钥对") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "请输入您本次投票使用的密钥") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "第三方公证人的初始化操作完成") {
			// "\033[31m update apt-get  更新apt-get \033[0m"
			fmt.Println("\033[31m 第三方公证人的初始化操作完成 \033[0m")

		} else if strings.Contains(data, "请选择从本地录入或者等待候选人提供信息") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "请输入本次投票制作的选票数") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "请输入本次参加选举的人数") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "候选人姓名") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "候选人自我介绍") {
			ScanlnToSock(conn)

		} else if strings.Contains(data, "开始录入候选者信息") {

		} else if strings.Contains(data, "继续等待其他候选人或") {
			ScanlnToSock(conn)
		} else if strings.Contains(data, "准备开始分发Paillier公钥") {
			PubKey = RecvPaillierKey(conn)
		} else if strings.Contains(data, "请输入所有候选人的姓名") {
			ScanlnToSock(conn)
		}
	}
}
func RecvPaillierKey(conn net.Conn) paillier.PublicKey {
	fmt.Println("等待paillier")
	data := ConveyUtils.RecvFrom(conn)
	fmt.Println("Paillier json is:", data)
	var PaillierPublicKey paillier.PublicKey
	err := json.Unmarshal(data, &PaillierPublicKey)
	if err != nil {
		fmt.Println("反射失败:", err)
	}
	return PaillierPublicKey
}

func ScanlnToSock(conn net.Conn) {
	fmt.Println("请输入对应的数据:")
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}

func Case3(conn net.Conn) {
	for {
		fmt.Println("是否生成新的RSA密钥对？ (Y/n)")
		var choice string
		fmt.Scanf("%s", &choice)
		if choice == "Y" || choice == "y" {
			RSAUtils.GenerateRSAKey(1024)
			break
		} else if choice == "N" || choice == "n" {
			break
		} else {
			fmt.Println("您的输入有误，请重新输入:")
		}
	}
	fmt.Println("开始准备传输RSA公钥")
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	out := ShellUtils.GetOutFromStdout("ls " + home + "/rsa/keys/")
	for i, v := range out {
		fmt.Printf("[%d : %v]\n", i, v)
	}
	var op int
	fmt.Println("请输入您要上传的公钥文件：")
	for {
		scanf, err := fmt.Scanf("%d", &op)
		if err != nil {
			fmt.Println("您的输入不合法，请重新输入：", scanf)
			return
		} else {
			break
		}
	}
	pubpath := home + "/rsa/keys/" + out[op]
	fmt.Println("您选择的文件是：", pubpath)
	ConveyUtils.ConveyFile(conn, pubpath)
}
