package VoteUtils

import (
	"SockGo/ConveyUtils"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
)

type Candidate struct { // 候选人
	ID           string
	Name         string
	Introduction string //候选人自我介绍
}

func (c *Candidate) SetCandidateInfo(conn net.Conn) {
	fmt.Printf("+--------------------+----------------+\n")
	fmt.Printf("|     候选人姓名     | 候选人自我介绍 |\n")
	fmt.Printf("+--------------------+----------------+\n")
	fmt.Println("请输入对应信息：")

	ConveyUtils.ConveyData(conn, []byte("+--------------------+----------------+\n")) //
	ConveyUtils.ConveyData(conn, []byte("|     候选人姓名     | 候选人自我介绍 |\n"))             //
	ConveyUtils.ConveyData(conn, []byte("+--------------------+----------------+\n")) //
	ConveyUtils.ConveyData(conn, []byte("请输入对应信息：\n"))                                //
	ConveyUtils.ConveyData(conn, []byte("_over"))

	fmt.Printf("候选人姓名：")
	ConveyUtils.ConveyData(conn, []byte("候选人姓名：\n")) //
	ConveyUtils.ConveyData(conn, []byte("_over"))
	//var Name string
	// scanf, err := fmt.Scanf("%s", &Name)
	data := ConveyUtils.RecvFrom(conn)
	strData := string(data)
	strData = strData[:len(strData)-1] // 这个是姓名
	c.Name = strData

	//for {
	//	fmt.Printf("候选人自我介绍：")
	//	ConveyUtils.ConveyData(conn, []byte("候选人自我介绍：\n")) //
	//	ConveyUtils.ConveyData(conn, []byte("_over"))
	//	var Introduction string
	//	//scanf, err := fmt.Scanf("%s", &Introduction)
	//	if err != nil {
	//		return
	//	} else {
	//		c.Introduction = Introduction
	//		break
	//	}
	//}
	ConveyUtils.ConveyData(conn, []byte("候选人自我介绍：\n")) //
	ConveyUtils.ConveyData(conn, []byte("_over"))
	data = ConveyUtils.RecvFrom(conn)
	strData = string(data)
	strData = strData[:len(strData)] // 这个是自我介绍
	c.Introduction = strData

	b, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	c.ID = "Candidate_" + fmt.Sprintf("%s", b)
	ConveyUtils.ConveyData(conn, []byte("输入完毕：\n")) //
	ConveyUtils.ConveyData(conn, []byte("_over"))
}
