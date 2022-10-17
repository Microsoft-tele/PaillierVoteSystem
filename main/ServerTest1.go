package main

import (
	"SockGo/ConveyUtils"
	"SockGo/CryptoUtils"
	"SockGo/VoteUtils"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"
	"time"
)

var MAX_CLINET_NUM = 150 // 包括投票人和被选举人的服务线程数
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

	// 是否扩大加入人数
	IsExpendMaxMember(conn)

	// 制作选票
	fmt.Println("开始制作选票：从公证人处获取所有参选人的姓名:")
	ConveyUtils.PrintStringToSock(conn, "开始制作选票，请输入所有候选人的姓名，人名中间以[:]（半角冒号）分割，例如 (李为君:何检涛)：\n")
	NameStr := ConveyUtils.RecvStringFrom(conn)
	NameList := strings.Split(NameStr, ":")
	fmt.Println("参选名单如下：")
	for i, v := range NameList {
		fmt.Printf("[%d : %v]\n", i, v)
	}

	ConveyUtils.PrintStringToSock(conn, "请输入制作的选票数量:\n")
	TicketNumStr := ConveyUtils.RecvStringFrom(conn)
	TicketNum, err := strconv.Atoi(TicketNumStr)
	if err != nil {
		fmt.Println("您的输入有误:", err)
	}

	var Tickets []VoteUtils.BallotTicket
	Tickets = make([]VoteUtils.BallotTicket, 0)
	for i := 0; i < TicketNum; i++ {
		Rand, err := rand.Int(rand.Reader, new(big.Int).SetInt64(999999999))
		if err != nil {
			fmt.Println("生成随机数失败")
		}
		ID := "Ticket_" + fmt.Sprintf("%v", Rand) // 如果这里有重复的，以后再解决吧

		tmpTick := VoteUtils.BallotTicket{
			ID:                ID,
			CandidateNum:      len(NameList),
			CandidateNameList: NameList,
			Option:            nil,
			Signature:         "",
		}
		Tickets = append(Tickets, tmpTick)
	}

	// 开始分发Paillier公钥
	PaillierPrivateKey := CryptoUtils.GetKeysFromJson(conn) // 制作公钥
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
		ConveyUtils.PrintStringToSock(v, "公证人发送：准备开始分发Paillier公钥:\n") // 必须自带换行
		// paillier to json
		PubKeyJson, err := json.Marshal(PaillierPrivateKey)
		if err != nil {
			fmt.Println("paillier转换失败:", err)
		}
		ConveyUtils.ConveyData(v, PubKeyJson)
		ConveyUtils.ConveyData(v, []byte("_over"))
		fmt.Printf("发送至[ %v ]完毕\n", v.RemoteAddr())
	}
	fmt.Println("Paillier公钥分发完毕")

	// 分发选票

	// 回收选票

	// 统计结果

	// 向所有人广播结果
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
	//ConveyUtils.ConveyData(conn, []byte("您是候选人或投票人?\n1.候选人\n2.投票人\n3.公证人"))
	//ConveyUtils.ConveyData(conn, []byte("_over"))
	ConveyUtils.PrintStringToSock(conn, "请选择您的身份:\n1.候选人\n2.投票人\n3.公证人\n")
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

func IsExpendMaxMember(conn net.Conn) {
	ConveyUtils.PrintStringToSock(conn, "是否修改加入投票的总人数:[目前默认最大加入人数是"+fmt.Sprintf("%d", MAX_CLINET_NUM)+"] (Y/n):\n") // 接收返回信息
	strData := ConveyUtils.RecvStringFrom(conn)
	fmt.Println("是否修改最大进入人数：", strData)
	if strData == "Y" || strData == "y" {
		for {
			ConveyUtils.ConveyData(conn, []byte("请输入修改的数字:\n")) // 接收返回信息
			ConveyUtils.ConveyData(conn, []byte("_over"))
			data1 := ConveyUtils.RecvFrom(conn)
			maxNumStr := string(data1)
			maxNumStr = maxNumStr[:len(maxNumStr)-1]
			maxNum, err := strconv.Atoi(maxNumStr)
			if err != nil {
				ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入:\n")) // 接收返回信息
				ConveyUtils.ConveyData(conn, []byte("_over"))
			} else {
				fmt.Println("用户输入的最大接入量：", maxNum)
				MAX_CLINET_NUM = maxNum
				break
			}
		}
	} else if strData == "N" || strData == "n" {
		fmt.Println("用户选择不修改最大加入人数")
	}
}
