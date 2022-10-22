package main

import (
	"SockGo/VoteUtils"
	"encoding/json"
	"fmt"
	"net"
)

var MAX_CLINET_NUM = 150 // 包括投票人和被选举人的服务线程数
var NotaryNum = 0
var ClinetCnt = 0
var CandidatesConns *[]net.Conn
var VotersConns *[]net.Conn
var Candidatas []VoteUtils.Candidate

var Notary VoteUtils.Notary

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

	CandidatesConns = new([]net.Conn) // 开辟内存空间
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
	// 开始验证身份
	buf := make([]byte, 1024)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("从[ %v ]接收失败 : err:[ %v ] : 接收的字节数:[ %v ]\n", conn.RemoteAddr(), err, read)
		return
	}
	buf = buf[:read]
	Identity := string(buf)
	fmt.Printf("接收到的数据: [ %v ]\n", Identity)
	if Identity == "Candidate" {
		fmt.Printf("[ %v ]选择候选人模式:\n", conn.RemoteAddr())
		*CandidatesConns = append(*CandidatesConns, conn) // 候选人进入队列
		ModelOfCandidate(conn)
	} else if Identity == "Voter" {
		fmt.Printf("[ %v ]选择投票人模式:\n", conn.RemoteAddr())
		*VotersConns = append(*VotersConns, conn) // 投票人进入队列
		ModelOfVoter(conn)
	} else if Identity == "Notary" {
		fmt.Printf("[ %v ]选择公证人模式:\n", conn.RemoteAddr())
		if NotaryNum < 1 { // 仅允许一个管理员登陆
			NotaryNum++
		}
		ModelOfNotary(conn)
		NotaryNum--
	}
}

func ModelOfCandidate(conn net.Conn) {
	fmt.Println("进入候选人模式:")
	write, err := conn.Write([]byte("请上传您的信息:"))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}
	fmt.Printf("向[ %v ]发送[ %v ]字节:", conn.RemoteAddr(), write)
	can := VoteUtils.Candidate{}
	buf := make([]byte, 1024)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("从[ %v ]接收失败 : err:[ %v ] : 接收的字节数:[ %v ]\n", conn.RemoteAddr(), err, read)
		return
	}
	buf = buf[:read] // 接收JSON数据
	fmt.Printf("接收到的JSON:\n[%v]\n", buf)
	err = json.Unmarshal(buf, &can)
	if err != nil {
		fmt.Println("候选人对象反射失败：", err)
		return
	}
	fmt.Println("候选者对象:")
	fmt.Println(can)
	Candidatas = append(Candidatas, can)
	write, err = conn.Write([]byte("成功收到您的信息，请等待结果公示"))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}
	fmt.Printf("向[ %v ]发送[ %v ]字节:", conn.RemoteAddr(), write)
}

func ModelOfVoter(conn net.Conn) {
	fmt.Println("进入投票人模式:")
	write, err := conn.Write([]byte("请等待公证人分发选票:"))
	if err != nil {
		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}
	fmt.Printf("向[ %v ]发送[ %v ]字节:", conn.RemoteAddr(), write)
	fmt.Println([]byte("请等待公证人分发选票:"))

}
func ModelOfNotary(conn net.Conn) {
	fmt.Println("进入公证人模式:")
	write, err := conn.Write([]byte("请确认所有候选人已经成功上传信息，所有投票人已经设置好个人信息，输入【1】开始上传公证人对象，服务器开始制作选票,并分发给所有投票人"))
	if err != nil {
		fmt.Println(write)
		fmt.Printf("向[ %v ]写入数据失败:", conn.RemoteAddr())
		return
	}
	buf := make([]byte, 10)
	n, _ := conn.Read(buf)
	buf = buf[:n]
	fmt.Println("公证人输入:", string(buf))
	if string(buf) == "1\n" {
		fmt.Println("准备开始接收公证人对象：")
	}
	buf = make([]byte, 1024*10)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("从[ %v ]读取数据失败:", conn.RemoteAddr())
		return
	}
	buf = buf[:read]
	fmt.Printf("从[ %v ]读取数据[ %v ]\n", conn.RemoteAddr(), buf)

	err = json.Unmarshal(buf, &Notary)
	if err != nil {
		fmt.Println("Notary反射失败:", err)
		return
	}
	fmt.Println("Notary对象反射成功")
	fmt.Println(Notary)

	write, err = conn.Write([]byte("请等待投票结果:"))
	if err != nil {
		fmt.Printf("向[ %v ]发送数据失败[ %v ][ %v ]\n", conn.RemoteAddr(), err, write)
		return
	}

	BallotOperateMachine := VoteUtils.BallotOperateMachine{
		BallotTicketNum: 0,
		BallotTickets:   nil,
	}
	Notary.Work(&BallotOperateMachine, Candidatas, VotersConns)
}
