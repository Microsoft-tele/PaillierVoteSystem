package main

import (
	"SockGo/ConveyUtils"
	"SockGo/CryptoUtils"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var MAX_CLINET_NUM = 150 // 包括投票人和被选举人的服务线程数
var lock sync.Mutex
var NotaryNum = 0
var ClinetCnt = 0
var CandidatesConns *[]net.Conn
var VotersConns *[]net.Conn

func main() {
	fmt.Println("服务器正在监听：")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("关闭监听失败，请查明原因：", err)
		}
	}(listen)
	if err != nil {
		fmt.Println("建立监听失败，请查明原因:", err)
	}

	CandidatesConns = new([]net.Conn)
	VotersConns = new([]net.Conn)

	for ClinetCnt < MAX_CLINET_NUM { // 多线程处理
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败，请查明原因:", err)
		} else {
			ClinetCnt++
			go Process(conn)
		}
	}
	fmt.Println("达到默认并发限制上限，请重新设置参数并重启系统")
}

func Process(conn net.Conn) {
	Hello(conn)
	model := WorkModel(conn)
	fmt.Println("Model is int", model)
	if model == 1 { // 候选人模式
		*CandidatesConns = append(*CandidatesConns, conn) // 候选人进入队列
		Model_1_TEST(conn)                                // 进入候选人处理函数

	} else if model == 2 { // 投票人模式
		*VotersConns = append(*VotersConns, conn) // 投票人进入队列
		Model_2_TEST(conn)

	} else if model == 3 { // 公证人模式
		fmt.Println("有客户选择了公证人模式，请注意！！！")
		if NotaryNum < 3 {
			Model_3_TEST(conn)
		} else {
			ConveyUtils.PrintStringToSock(conn, "本次投票已经产生第三方监控员，您无权限加入管理员:\n")
			return
		}
	}
}

func Model_1_TEST(conn net.Conn) {
	ConveyUtils.PrintStringToSock(conn, "进入候选人模式界面:\n")
}
func Model_2_TEST(conn net.Conn) {
	ConveyUtils.PrintStringToSock(conn, "进入投票人模式界面:\n")

}
func Model_3_TEST(conn net.Conn) {
	ConveyUtils.PrintStringToSock(conn, "进入公证人模式界面:\n")
	// 开始分发Paillier公钥
	PaillierPrivateKey := CryptoUtils.GetKeysFromJson(conn)
	fmt.Println(PaillierPrivateKey)

	for {
		ConveyUtils.PrintStringToSock(conn, "是否继续等待其他候选人或投票人:10秒后自动停止等待:(1/0)\n")
		data := ConveyUtils.RecvStringFrom(conn)
		intData, err := strconv.Atoi(data)
		if err != nil {
			ConveyUtils.PrintStringToSock(conn, "您的输入不合法\n")
		}
		if intData == 1 {
			continue
		} else if intData == 0 {
			break
		} else {
			fmt.Println("您的输入不合法，请重新输入")
		}
	}
	for i, v := range *VotersConns {
		fmt.Printf("[第%d个客户端 : %v]\n", i, v)
		ConveyUtils.PrintStringToSock(v, "公证人发送：准备开始分发Paillier公钥:\n")
	}
	fmt.Println("Paillier公钥分发完毕")

	for {
		fmt.Println("waitng")
		time.Sleep(10 * time.Second)
	}
}
func Hello(conn net.Conn) { // 建立成功进行问好
	ConveyUtils.ConveyData(conn, []byte("成功建立连接，你好！\n"))
	ConveyUtils.ConveyData(conn, []byte("_over"))
	return
}
func WorkModel(conn net.Conn) (model int) {
LOOP:
	ConveyUtils.ConveyData(conn, []byte("您是候选人或投票人?\n1.候选人\n2.投票人\n3.公证人"))
	ConveyUtils.ConveyData(conn, []byte("_over"))
	strData := ConveyUtils.RecvStringFrom(conn)
	fmt.Println("strData:", strData)
	model, err := strconv.Atoi(strData)
	if err != nil {
		fmt.Println("err is ", err)
		ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入：\n"))
		ConveyUtils.ConveyData(conn, []byte("_over"))
		goto LOOP
	} else {
		fmt.Println("接收到来自用户的Model选择，进行下一步判断")
	}
	return model
}
