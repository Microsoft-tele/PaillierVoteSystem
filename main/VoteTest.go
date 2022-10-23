package main

import (
	"RemoteRouter/VoteUtils"
	"fmt"
)

func main() {
	//Notary := VoteUtils.Notary{} // 生成一个公证人
	//Notary.InitNotary()                                     // 初始化公证人
	BallotTicketMachine := VoteUtils.BallotOperateMachine{} // 建立一个选票器
	//Notary.Work(&BallotTicketMachine)                       // 公证人开始工作
	fmt.Println(BallotTicketMachine.BallotTickets) // 调试信息
}
