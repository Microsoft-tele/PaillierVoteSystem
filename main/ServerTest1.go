package main

//var MAX_CLINET_NUM = 150 // 包括投票人和被选举人的服务线程数
//var NotaryNum = 0
//var ClinetCnt = 0
//var CandidatesConns *[]net.Conn
//var VotersConns *[]net.Conn

//func main() {
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
//	CandidatesConns = new([]net.Conn)
//	VotersConns = new([]net.Conn)
//
//	for ClinetCnt < MAX_CLINET_NUM { // 多线程处理
//		conn, err := listen.Accept()
//		if err != nil {
//			fmt.Println("建立连接失败，请查明原因:", err)
//		} else {
//			ClinetCnt++
//			go Process(conn)
//		}
//	}
//	fmt.Println("达到默认并发限制上限，请重新设置参数并重启系统")
//}

//func Process(conn net.Conn) {
//	Hello(conn)
//	model := WorkModel(conn)
//	fmt.Println("Model is int", model)
//	if model == 1 { // 候选人模式
//		*CandidatesConns = append(*CandidatesConns, conn) // 候选人进入队列
//		Model_1_TEST(conn)                                // 进入候选人处理函数
//
//	} else if model == 2 { // 投票人模式
//		*VotersConns = append(*VotersConns, conn) // 投票人进入队列
//		Model_2_TEST(conn)
//
//	} else if model == 3 { // 公证人模式
//		fmt.Println("有客户选择了公证人模式，请注意！！！")
//		if NotaryNum < 3 {
//			Model_3_TEST(conn)
//		} else {
//			ConveyUtils.PrintStringToSock(conn, "本次投票已经产生第三方监控员，您无权限加入管理员:\n")
//			return
//		}
//	}
//}
//
//func Model_1_TEST(conn net.Conn) {
//	ConveyUtils.PrintStringToSock(conn, "进入候选人模式界面:\n")
//}
//
//func Model_2_TEST(conn net.Conn) {
//	ConveyUtils.PrintStringToSock(conn, "进入投票人模式界面:\n")
//
//}
//
//func Model_3_TEST(conn net.Conn) {
//	ConveyUtils.PrintStringToSock(conn, "进入公证人模式界面:\n")
//	// 是否扩大加入人数
//	IsExpendMaxMember(conn)
//	// 制作选票
//	Tickets := MakeTickets(conn)
//
//	for {
//		ConveyUtils.PrintStringToSock(conn, "是否继续等待其他候选人或投票人:10秒后自动停止等待:(1/0)\n")
//		data := ConveyUtils.RecvStringFrom(conn)
//		intData, err := strconv.Atoi(data)
//		if err != nil {
//			ConveyUtils.PrintStringToSock(conn, "您的输入不合法\n")
//		}
//		if intData == 1 {
//			continue
//		} else if intData == 0 {
//			break
//		} else {
//			fmt.Println("您的输入不合法，请重新输入")
//		}
//	} // 选择是否等地其他人进入系统
//
//	// 分发选票
//	Distribute(conn, Tickets)
//
//	// 回收选票
//
//	// 统计结果
//
//	// 向所有人广播结果
//	for {
//		fmt.Println("waitng")
//		time.Sleep(10 * time.Second)
//	}
//}
//
//func Hello(conn net.Conn) { // 建立成功进行问好
//	ConveyUtils.ConveyData(conn, []byte("成功建立连接，你好！\n"))
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//	return
//}
//
//func WorkModel(conn net.Conn) (model int) {
//LOOP:
//	//ConveyUtils.ConveyData(conn, []byte("您是候选人或投票人?\n1.候选人\n2.投票人\n3.公证人"))
//	//ConveyUtils.ConveyData(conn, []byte("_over"))
//	//ConveyUtils.PrintStringToSock(conn, "请选择您的身份:\n1.候选人\n2.投票人\n3.公证人1234\n")
//	conn.Write([]byte("请选择您的身份:\n1.候选人\n2.投票人\n3.公证人1234\n_over"))
//	strData := ConveyUtils.RecvStringFrom(conn)
//	fmt.Println("strData:", strData)
//	model, err := strconv.Atoi(strData)
//	if err != nil {
//		fmt.Println("err is ", err)
//		ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入：\n"))
//		ConveyUtils.ConveyData(conn, []byte("_over"))
//		goto LOOP
//	} else {
//		fmt.Println("接收到来自用户的Model选择，进行下一步判断")
//	}
//	return model
//}
//
//func IsExpendMaxMember(conn net.Conn) {
//	//ConveyUtils.PrintStringToSock(conn, "是否修改加入投票的总人数:[目前默认最大加入人数是"+fmt.Sprintf("%d", MAX_CLINET_NUM)+"] (Y/n):\n") // 接收返回信息
//	conn.Write([]byte("是否修改加入投票的总人数:[目前默认最大加入人数是" + fmt.Sprintf("%d", MAX_CLINET_NUM) + "] (1/0):\n_over"))
//	//strData := ConveyUtils.RecvStringFrom(conn)
//	buf := make([]byte, 10)
//	n, _ := conn.Read(buf)
//	buf = buf[:n]
//	intBuf, _ := strconv.Atoi(string(buf))
//	fmt.Println("是否修改最大进入人数：", intBuf)
//	if intBuf == 1 {
//		for {
//			ConveyUtils.ConveyData(conn, []byte("请输入修改的数字:\n")) // 接收返回信息
//			ConveyUtils.ConveyData(conn, []byte("_over"))
//			data1 := ConveyUtils.RecvFrom(conn)
//			maxNumStr := string(data1)
//			maxNumStr = maxNumStr[:len(maxNumStr)-1]
//			maxNum, err := strconv.Atoi(maxNumStr)
//			if err != nil {
//				ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入:\n")) // 接收返回信息
//				ConveyUtils.ConveyData(conn, []byte("_over"))
//			} else {
//				fmt.Println("用户输入的最大接入量：", maxNum)
//				MAX_CLINET_NUM = maxNum
//				break
//			}
//		}
//	} else if intBuf == 0 {
//		fmt.Println("用户选择不修改最大加入人数")
//	}
//}
//
//func MakeTickets(conn net.Conn) []VoteUtils.BallotTicket {
//	// 开始分发Paillier公钥
//	PaillierPrivateKey := CryptoUtils.GetKeysFromJson(conn) // 制作公钥
//	fmt.Println(PaillierPrivateKey)
//
//	fmt.Println("开始制作选票：从公证人处获取所有参选人的姓名:")
//	//ConveyUtils.PrintStringToSock(conn, "开始制作选票，请输入所有候选人的姓名，人名中间以[:]（半角冒号）分割，例如 (李为君:何检涛)：\n")
//	conn.Write([]byte("开始制作选票，请输入所有候选人的姓名，人名中间以[:]（半角冒号）分割，例如 (李为君:何俭涛)：\n_over"))
//	NameStr := ConveyUtils.RecvStringFrom(conn)
//	NameList := strings.Split(NameStr, ":")
//	fmt.Println("参选名单如下：")
//	for i, v := range NameList {
//		fmt.Printf("[%d : %v]\n", i, v)
//	}
//
//	ConveyUtils.PrintStringToSock(conn, "请输入制作的选票数量:\n")
//	TicketNumStr := ConveyUtils.RecvStringFrom(conn)
//	TicketNum, err := strconv.Atoi(TicketNumStr)
//	if err != nil {
//		fmt.Println("您的输入有误:", err)
//	}
//
//	var Tickets []VoteUtils.BallotTicket
//	Tickets = make([]VoteUtils.BallotTicket, 0)
//	for i := 0; i < TicketNum; i++ {
//		Rand, err := rand.Int(rand.Reader, new(big.Int).SetInt64(999999999))
//		if err != nil {
//			fmt.Println("生成随机数失败")
//		}
//		ID := "Ticket_" + fmt.Sprintf("%v", Rand) // 如果这里有重复的，以后再解决吧
//
//		tmpTick := VoteUtils.BallotTicket{
//			ID:                ID,
//			CandidateNum:      len(NameList),
//			CandidateNameList: NameList,
//			Option:            nil,
//			Signature:         "",
//			PaillierPublicKey: PaillierPrivateKey.PublicKey,
//			RSAPublicKey:      nil,
//		}
//		Tickets = append(Tickets, tmpTick)
//	}
//	fmt.Println("已经成成的选票如下：")
//	fmt.Println("Tickets:", Tickets)
//	return Tickets
//}
//
//func Distribute(conn net.Conn, Tickets []VoteUtils.BallotTicket) {
//	// 将选票序列化然后发送给投票人
//	TicketsJson := make([][]byte, 0)
//	for i := 0; i < len(Tickets); i++ {
//		ticketJson, err := json.Marshal(Tickets[i])
//		if err != nil {
//			fmt.Println("选票序列化失败")
//		}
//		TicketsJson = append(TicketsJson, ticketJson)
//	}
//	fmt.Println("TicketJson:", TicketsJson)
//	fmt.Println("序列化选票:", Tickets)
//	// 分发选票
//	//ConveyUtils.PrintStringToSock(conn, "是否开始分发选票:(1/0)\n")
//	conn.Write([]byte("是否开始分发选票:(1/0)\n_over"))
//	data := ConveyUtils.RecvStringFrom(conn)
//	intData, err := strconv.Atoi(data)
//	if err != nil {
//		fmt.Println("转换失败")
//	}
//	for {
//		if intData == 1 {
//			break
//		} else {
//			continue
//		}
//	}
//	for i, v := range *VotersConns {
//		fmt.Println("Conn info:", v.RemoteAddr())
//		fmt.Println("开始向上述地址传递选票")
//		//ConveyUtils.PrintStringToSock(v, "开始接收选票\n")
//		v.Write([]byte("是否开始接收选票\n_over"))
//		buf := make([]byte, 1)
//		v.Read(buf)
//		//ConveyUtils.ConveyData(v, TicketsJson[i]) // 这里是没有判断选票和选民人数是否匹配的，后期优化再说
//		v.Write(TicketsJson[i])
//		v.Write([]byte("_over"))
//		//v.Write([]byte("_over"))
//		fmt.Println("TicketJson:", TicketsJson[i])
//		//ConveyUtils.ConveyData(v, []byte("_over")) // 这里是没有判断选票和选民人数是否匹配的，后期优化再说
//	}
//	for _, v := range *VotersConns {
//		v.Write([]byte("分发结束\n_over"))
//	}
//}
