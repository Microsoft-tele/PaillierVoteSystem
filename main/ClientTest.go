package main

import (
	"SockGo/ConveyUtils"
	"SockGo/RSAUtils"
	"SockGo/ShellUtils"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

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
	var data []byte

	data = make([]byte, 1024)
	for {
		data = ConveyUtils.RecvFrom(conn)
		strData := string(data)
		fmt.Println("服务端发来的数据:", string(strData))
		if strings.Contains(strData, "您是候选人或投票人") {
			Case1(conn)

		} else if strings.Contains(strData, "输入您的基本信息以获得投票资格证") {
			Case2(conn)
		} else if strings.Contains(strData, "传您的RSA公钥以获得") {
			// 向服务器传输RSA公钥
			Case3(conn)

		} else if strings.Contains(strData, "是否修改加入投票的总人数") {
			Case4(conn)

		} else if strings.Contains(strData, "请输入修改的数字") {
			Case5(conn)

		} else if strings.Contains(strData, "功能菜单") {
			Case6(conn)

		} else if strings.Contains(strData, "是否生成新的Paillier密钥对") {
			Case7(conn)

		} else if strings.Contains(strData, "请输入您本次投票使用的密钥") {
			Case6(conn)

		} else if strings.Contains(strData, "第三方公证人的初始化操作完成") {
			// "\033[31m update apt-get  更新apt-get \033[0m"
			fmt.Println("\033[31m 第三方公证人的初始化操作完成 \033[0m")

		} else if strings.Contains(strData, "请选择从本地录入或者等待候选人提供信息") {
			Case6(conn)

		} else if strings.Contains(strData, "请输入本次投票制作的选票数") {
			Case7(conn)

		} else if strings.Contains(strData, "请输入本次参加选举的人数") {
			Case7(conn)

		} else if strings.Contains(strData, "候选人姓名") {
			Case7(conn)

		} else if strings.Contains(strData, "候选人自我介绍") {
			Case7(conn)

		} else if strings.Contains(strData, "开始录入候选者信息") {

		}
	}
}

func ConveyFile(conn net.Conn, filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件错误：", err)
	}
	reader := bufio.NewReader(f)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕:", err, n)
			ConveyUtils.ConveyData(conn, []byte("_over"))
			break
		}
		ConveyUtils.ConveyData(conn, buf)
	}
}

func Case1(conn net.Conn) {
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
func Case2(conn net.Conn) {
	fmt.Println("请输入您的姓名：")
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
	ConveyFile(conn, pubpath)
}
func Case4(conn net.Conn) {
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
func Case5(conn net.Conn) {
	fmt.Println("请输入要修改的数字：")
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
func Case6(conn net.Conn) {
	fmt.Println("请输入对应的数字：")
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
func Case7(conn net.Conn) {
	fmt.Println("请输入：")
	sendData, _ := reader.ReadString('\n')
	ConveyUtils.ConveyData(conn, []byte(sendData))
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
func Case8(conn net.Conn) {

}
func Case9(conn net.Conn) {

}
