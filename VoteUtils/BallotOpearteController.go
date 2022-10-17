package VoteUtils

import (
	"SockGo/ConveyUtils"
	"fmt"
	"net"
	"strconv"
)

var CandidatesCnt int

type BallotOperateMachine struct { // 选票制作机器
	BallotTicketNum int
	BallotTickets   []BallotTicket // 等待收发选票
	CandidateNum    int
	VoterNum        int
}
type BallotMakeController interface {
	MakeBallotTickets(conn net.Conn, CandidateConnList *[]net.Conn) // 制作选票
	//DistributeBallots() // 分发选票
	//TakeBackBallots()   // 回收选票
	//StatisticResult()   // 统计结果
}

func (b *BallotOperateMachine) MakeBallotTickets(conn net.Conn, CandidateConnList *[]net.Conn) {
	// 选择手动录入或者等待选手提供
	fmt.Println("请选择从本地录入或者等待候选人提供信息：")
	fmt.Println("0.手动输入")
	fmt.Println("1.等待选手录入")
	ConveyUtils.ConveyData(conn, []byte("请选择从本地录入或者等待候选人提供信息：:\n")) //
	ConveyUtils.ConveyData(conn, []byte("0.手动输入:\n"))               //
	ConveyUtils.ConveyData(conn, []byte("1.等待选手录入:\n"))             //
	ConveyUtils.ConveyData(conn, []byte("_over"))
	var op int
	for {
		data := ConveyUtils.RecvFrom(conn)
		strData := string(data)
		strData = strData[:len(strData)-1]
		fmt.Println("strData:", strData)
		// scanf, err := fmt.Scanf("%d", &op)
		intData, err := strconv.Atoi(strData)
		if err != nil {
			ConveyUtils.ConveyData(conn, []byte("您的输入不合法，您的输入必须是1或0，请重新输入\n")) //
			ConveyUtils.ConveyData(conn, []byte("_over"))
		} else if intData == 1 || intData == 0 {
			op = intData
			break
		} else {
			fmt.Println("您的输入必须是1或0，请重新输入")
		}
	}
	if op == 0 {
		//Func1(conn, b) // 封装操作
		fmt.Println("请输入本次投票制作的选票数：注意一旦输入不可更改")
		ConveyUtils.ConveyData(conn, []byte("请输入本次投票制作的选票数：注意一旦输入不可更改\n")) //
		ConveyUtils.ConveyData(conn, []byte("_over"))
		//var TicketsCnt int
		for {
			data := ConveyUtils.RecvFrom(conn)
			strData := string(data)
			strData = strData[:len(strData)-1]
			fmt.Println("strData:", strData)
			// scanf, err := fmt.Scanf("%d", &op)
			intData, err := strconv.Atoi(strData)
			if err != nil {
				ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入\n")) //
				ConveyUtils.ConveyData(conn, []byte("_over"))
			} else {
				b.BallotTicketNum = intData // 修正本次投票的票数
				fmt.Println("BallotTicketNum is update:", intData)
				break
			}
		}
		fmt.Println("请输入本次参加选举的人数：注意一旦输入不可更改")
		ConveyUtils.ConveyData(conn, []byte("请输入本次参加选举的人数：注意一旦输入不可更改\n")) //
		ConveyUtils.ConveyData(conn, []byte("_over"))

		for {
			data := ConveyUtils.RecvFrom(conn)
			strData := string(data)
			strData = strData[:len(strData)-1]
			fmt.Println("strData:", strData)
			// scanf, err := fmt.Scanf("%d", &op)
			intData, err := strconv.Atoi(strData)
			if err != nil {
				ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入\n")) //
				ConveyUtils.ConveyData(conn, []byte("_over"))
			} else {
				fmt.Println("客户端发送：选举人数：", intData)
				b.CandidateNum = intData
				break
			}
		}
		var canList []Candidate
		for i := 0; i < b.CandidateNum; i++ {
			can := Candidate{}
			can.SetCandidateInfo(conn)
			canList = append(canList, can)
		}
		for i := 0; i < b.BallotTicketNum; i++ {
			BallotTick := BallotTicket{
				ID:            strconv.Itoa(i),
				CandidateNum:  CandidatesCnt,
				CandidateList: canList,
				Option:        nil,
				Signature:     "nil",
			}
			b.BallotTickets = append(b.BallotTickets, BallotTick)
		}

	} else if op == 1 {
		fmt.Println("从socket接收")
		// 通知所有人录入信息
		ConveyUtils.ConveyData(conn, []byte("准备通知所有选手进行信息录入\n")) //
		ConveyUtils.ConveyData(conn, []byte("_over"))
		// 向通知信道发送信息
		fmt.Println("当前在线选手数量为：", len(*CandidateConnList))
		for i := 0; i < len(*CandidateConnList); i++ {
			ConveyUtils.ConveyData((*CandidateConnList)[i], []byte("开始录入候选者信息：\n")) //
			ConveyUtils.ConveyData((*CandidateConnList)[i], []byte("_over"))
			go Func2((*CandidateConnList)[i], b)
		}
	}
}

func Func1(conn net.Conn, b *BallotOperateMachine) {
	fmt.Println("请输入本次投票制作的选票数：注意一旦输入不可更改")
	ConveyUtils.ConveyData(conn, []byte("请输入本次投票制作的选票数：注意一旦输入不可更改\n")) //
	ConveyUtils.ConveyData(conn, []byte("_over"))
	//var TicketsCnt int
	for {
		data := ConveyUtils.RecvFrom(conn)
		strData := string(data)
		strData = strData[:len(strData)-1]
		fmt.Println("strData:", strData)
		// scanf, err := fmt.Scanf("%d", &op)
		intData, err := strconv.Atoi(strData)
		if err != nil {
			ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入\n")) //
			ConveyUtils.ConveyData(conn, []byte("_over"))
		} else {
			b.BallotTicketNum = intData // 修正本次投票的票数
			fmt.Println("BallotTicketNum is update:", intData)
			break
		}
	}
	fmt.Println("请输入本次参加选举的人数：注意一旦输入不可更改")
	ConveyUtils.ConveyData(conn, []byte("请输入本次参加选举的人数：注意一旦输入不可更改\n")) //
	ConveyUtils.ConveyData(conn, []byte("_over"))

	for {
		data := ConveyUtils.RecvFrom(conn)
		strData := string(data)
		strData = strData[:len(strData)-1]
		fmt.Println("strData:", strData)
		// scanf, err := fmt.Scanf("%d", &op)
		intData, err := strconv.Atoi(strData)
		if err != nil {
			ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入\n")) //
			ConveyUtils.ConveyData(conn, []byte("_over"))
		} else {
			fmt.Println("客户端发送：选举人数：", intData)
			b.CandidateNum = intData
			break
		}
	}
	var canList []Candidate
	for i := 0; i < b.CandidateNum; i++ {
		can := Candidate{}
		can.SetCandidateInfo(conn)
		canList = append(canList, can)
	}
	for i := 0; i < b.BallotTicketNum; i++ {
		BallotTick := BallotTicket{
			ID:            strconv.Itoa(i),
			CandidateNum:  CandidatesCnt,
			CandidateList: canList,
			Option:        nil,
			Signature:     "nil",
		}
		b.BallotTickets = append(b.BallotTickets, BallotTick)
	}
}

func Func2(conn net.Conn, b *BallotOperateMachine) {
	var canList []Candidate
	for i := 0; i < b.CandidateNum; i++ {
		can := Candidate{}
		can.SetCandidateInfo(conn)
		canList = append(canList, can)
	}
	for i := 0; i < b.BallotTicketNum; i++ {
		BallotTick := BallotTicket{
			ID:            strconv.Itoa(i),
			CandidateNum:  CandidatesCnt,
			CandidateList: canList,
			Option:        nil,
			Signature:     "nil",
		}
		b.BallotTickets = append(b.BallotTickets, BallotTick)
	}
}
