package main

//import (
//	"RemoteRouter/ConveyUtils"
//	"RemoteRouter/CryptoUtils"
//	"RemoteRouter/ShellUtils"
//	"RemoteRouter/VoteUtils"
//	"RemoteRouter/paillier"
//	"bufio"
//	"fmt"
//	"net"
//	"os"
//	"strconv"
//	"sync"
//)
//
//// var MAX_CLINET_NUM = 150 // 包括投票人和被选举人的服务线程数
//var lock sync.Mutex
//
////var NotaryNum = 0
////var Info = 0
//
//func main() {
//	InitWork()
//}
//func InitWork() {
//	fmt.Println("服务器正在监听：")
//	listen, err := net.Listen("tcp", "0.0.0.0:8888")
//	defer func(listen net.Listener) {
//		err := listen.Close()
//		if err != nil {
//			fmt.Println("关闭监听失败，请查明原因：", err)
//		}
//	}(listen)
//	if err != nil {
//		fmt.Println("建立监听失败，请查明原因:", err)
//	}
//
//	ClinetCnt := 0
//	// var Info chan bool // 第三方公证人进行广播的通道
//
//	var ConnListOfCandidates []net.Conn // 候选者列表
//	var ConnListOfVoters []net.Conn     // 投票者列表
//	var ConnOfNotary net.Conn           // 第三方公证人
//
//	ConnListOfCandidates = make([]net.Conn, 0)
//	ConnListOfVoters = make([]net.Conn, 0)
//
//	for ClinetCnt < MAX_CLINET_NUM { // 多线程处理
//		conn, err := listen.Accept()
//		if err != nil {
//			fmt.Println("建立连接失败，请查明原因:", err)
//		} else {
//			go ChildProcess(conn, &ConnListOfCandidates, &ConnListOfVoters, &ConnOfNotary) // 开启子线程，不影响主线程继续连接用户
//			// ConnList = append(ConnList, conn) // 将所有加入系统法的连接存储起来，方便后续第三方公证人分发资料
//			ClinetCnt++
//		}
//	}
//	fmt.Println("达到默认并发限制上限，请重新设置参数并重启系统")
//}
//
//func RecvRSAPublicKeyFile(conn net.Conn, filePath string) {
//	fmt.Println("从用户处接收RSA公钥")
//	ConveyUtils.ConveyData(conn, []byte("请上传您的RSA公钥以获得投票资格证:\n")) // 这个地方需要把所有用户的RAS公钥存好
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//	data := ConveyUtils.RecvFrom(conn)
//	fmt.Println("接收到的RSA公钥文件是：", data)
//	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
//	//out := ShellUtils.GetOutFromStdout("ls " + home + "/rsa/keys/")
//	savePath := home + "/serverData/RSAPubKey/" + filePath + "_pub.pem"
//	f, err := os.OpenFile(savePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
//	if err != nil {
//		fmt.Println("打开保存文件失败:", err)
//		return
//	}
//	writer := bufio.NewWriter(f)
//	for i := 0; i < len(data); i++ {
//		if data[i] == 0 {
//			break
//		}
//		err := writer.WriteByte(data[i])
//		if err != nil {
//			fmt.Println("写入文件失败:", err)
//			return
//		}
//	}
//	writer.Flush()
//	f.Close()
//	fmt.Println("保存用户公钥成功！")
//}
//
//func ChildProcess(conn net.Conn, Candidates *[]net.Conn, Voters *[]net.Conn, Notary *net.Conn) {
//	Hello(conn)
//	model := WorkModel(conn)
//	fmt.Println("Model is int", model)
//	if model == 1 { // 候选人模式
//		lock.Lock()
//		*Candidates = append(*Candidates, conn) // 做好记录
//		fmt.Println("当前在线人数：", len(*Candidates))
//		lock.Unlock()
//		Model_1(conn) // 进入候选人处理函数
//
//	} else if model == 2 { // 投票人模式
//		//lock.Lock()
//		*Voters = append(*Voters, conn) // 做好记录
//		//lock.Unlock()
//		// 获取投票人的基本信息
//		Model_2(conn) // 进入投票人的处理
//
//	} else if model == 3 { // 公证人模式
//		fmt.Println("有客户选择了公证人模式，请注意！！！")
//		if NotaryNum < 3 {
//			lock.Lock()
//			*Notary = conn
//			lock.Unlock()
//			NotaryNum++
//			Model_3(conn, Candidates) //进入公证人的处理
//
//		} else {
//			ConveyUtils.ConveyData(conn, []byte("本次投票已经产生第三方监控员，您无权限加入管理员"))
//			ConveyUtils.ConveyData(conn, []byte("_over"))
//			return
//		}
//	}
//}
//
////func Hello(conn net.Conn) { // 建立成功进行问好
////	ConveyUtils.ConveyData(conn, []byte("成功建立连接，你好！"))
////	ConveyUtils.ConveyData(conn, []byte("_over"))
////	return
////}
//
////func WorkModel(conn net.Conn) (model int) {
////LOOP:
////	ConveyUtils.ConveyData(conn, []byte("您是候选人或投票人?\n1.候选人\n2.投票人\n3.公证人"))
////	ConveyUtils.ConveyData(conn, []byte("_over"))
////	data := ConveyUtils.RecvFrom(conn)
////	strData := string(data)
////	strData = strData[:len(strData)-1]
////	fmt.Println("strData:", strData)
////	model, err := strconv.Atoi(strData)
////	if err != nil {
////		fmt.Println("err is ", err)
////		ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入：\n"))
////		ConveyUtils.ConveyData(conn, []byte("_over"))
////		goto LOOP
////	} else {
////		fmt.Println("接收到来自用户的Model选择，进行下一步判断")
////	}
////	return model
////}
//
//func Model_1(conn net.Conn) {
//
//	Hello(conn)
//	ConveyUtils.ConveyData(conn, []byte("请等待第三方公证人的指令:\n")) // 这个地方只接受姓名即可
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//
//	// 等待公证人操作
//}
//func Model_2(conn net.Conn) {
//	ConveyUtils.ConveyData(conn, []byte("请输入您的基本信息以获得投票资格证:\n")) // 这个地方只接受姓名即可
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//	data := ConveyUtils.RecvFrom(conn)
//	Name := string(data)
//	fmt.Println("接收到投票人的姓名:", Name[:len(Name)-1])
//	RecvRSAPublicKeyFile(conn, Name[:len(Name)-1]) // 把用户公钥按照每个人的名字存入本地文件进行记录
//}
//func Model_3(conn net.Conn, CandidateConnList *[]net.Conn) {
//	// 选择是否修改最大加入人数
//
//	// 生成一个公证人
//	Notary := NotaryInitWork(conn)
//
//	for { // 发送菜单给客户
//		ConveyUtils.ConveyData(conn, []byte("功能菜单:\n"))    //
//		ConveyUtils.ConveyData(conn, []byte("1.制作选票:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("2.分发选票:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("3.统计结果:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("4.查看结果:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("5.查看密文:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("6.验证签名:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("7.退出程序:\n"))  //
//		ConveyUtils.ConveyData(conn, []byte("8.调试:\n"))    //
//		ConveyUtils.ConveyData(conn, []byte("请输入您的选择:\n")) //
//		ConveyUtils.ConveyData(conn, []byte("_over"))
//
//		data := ConveyUtils.RecvFrom(conn)
//		strData := string(data)
//		strData = strData[:len(strData)-1]
//		fmt.Println("strData:", strData)
//		numDataOfChoice, err := strconv.Atoi(strData)
//		if err != nil {
//			fmt.Println("To int err:", err)
//		}
//
//		//-------------------------------------------------------------------------
//		//-------------------------------------------------------------------------
//		// 生成菜单
//		BallotMakeMachine := VoteUtils.BallotOperateMachine{ // 创建选票制作器
//			BallotTicketNum: 0,
//			BallotTickets:   nil,
//			CandidateNum:    0,
//			VoterNum:        0,
//		}
//		switch numDataOfChoice {
//		case 1:
//			fmt.Println("用户选择制作选票")
//			// 通知
//			// 阻塞至所有候选者完成信息录入
//
//			Notary.Work(&BallotMakeMachine, conn, CandidateConnList) // 有对全频道广播的能力
//
//		case 2:
//			fmt.Println("用户选择分发选票")
//		case 3:
//			fmt.Println("用户选择统计结果")
//		case 4:
//			fmt.Println("用户选择查看结果")
//		case 5:
//			fmt.Println("用户选择查看密文")
//		case 6:
//			fmt.Println("用户选择验证签名")
//		case 7:
//			fmt.Println("用户选择验退出程序")
//			NotaryNum-- // 把第三方线程数量减1
//			return
//		case 8:
//			fmt.Println("产生调试信息：")
//			fmt.Println(BallotMakeMachine)
//		default:
//			fmt.Println("您的输入不合法，请重新输入：")
//		}
//	}
//}
//
//func NotaryInitWork(conn net.Conn) VoteUtils.Notary {
//	ConveyUtils.ConveyData(conn, []byte("是否生成新的Paillier密钥对:(Y/n)\n")) // 接收返回信息
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//	data := ConveyUtils.RecvFrom(conn)
//	strData := string(data)
//	strData = strData[:len(strData)-1]
//
//	var paillierKey *paillier.PrivateKey
//	if strData == "Y" || strData == "y" {
//		fmt.Println("用户选择生成新的Paillier密钥对")
//		paillierKey = CryptoUtils.CreateKeys(1024)
//	} else if strData == "N" || strData == "n" {
//		fmt.Println("用户选择不生成新的Paillier密钥对")
//		paillierKey = CryptoUtils.GetKeysFromJson(conn)
//	}
//
//	Notary := VoteUtils.Notary{
//		ID:                 "1",
//		Name:               "Notary",
//		PaillierPublicKey:  paillierKey.PublicKey,
//		PaillierPrivatekey: paillierKey,
//		RsaPublicKey:       nil,
//	}
//	fmt.Println(Notary)
//	fmt.Println("第三方公证人的初始化操作完成！！！")
//	ConveyUtils.ConveyData(conn, []byte("第三方公证人的初始化操作完成！！！\n")) // 不接收返回信息
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//	return Notary
//}
