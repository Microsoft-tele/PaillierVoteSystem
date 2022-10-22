package main

//var Voter VoteUtils.Voter

//var Ticket VoteUtils.BallotTicket

//func main() {
//
//	var Name string
//	fmt.Println("请输入您的姓名:")
//	scanf, err := fmt.Scanf("%s", &Name)
//	if err != nil {
//		fmt.Println("您的输入不合法:", err, scanf)
//		return
//	}
//	//privateKeyPath := home + "/rsa/keys/" + NowTime[0] + "_" + NowTime[1] + "_pri.pem" // 需要改进
//	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
//	dirList := ShellUtils.GetOutFromStdout("ls " + home + "/rsa/keys/")
//	for i, v := range dirList {
//		fmt.Printf("[%d : %v]\n", i, v)
//	}
//	fmt.Println("请输入您的RSA私钥路径选项[私钥不会上传至服务器，请放心]:")
//	var tmpPri int
//	fmt.Scanf("%d", &tmpPri)
//	RSAPrivateKeyPath := home + "/rsa/keys/" + dirList[tmpPri]
//	fmt.Println("您输入的路径是:", RSAPrivateKeyPath)
//
//	fmt.Println("请输入您的RSA公钥路径选项(注意两个密钥是匹配的，如不确定请重新生成，如果导致电子签名失效后果自负):")
//	var tmpPub int
//	fmt.Scanf("%d", &tmpPub)
//	RSAPublicKeyPath := home + "/rsa/keys/" + dirList[tmpPub]
//	fmt.Println("您输入的路径是:", RSAPublicKeyPath)
//
//	Voter.InitVoter(Name, RSAPrivateKeyPath, RSAPublicKeyPath)
//	fmt.Println(Voter)
//
//LOOP:
//	ServiceHostIp := "127.0.0.1"
//	ServiceHostPort := "8888"
//	address := ServiceHostIp + ":" + ServiceHostPort
//	conn, err := net.Dial("tcp", address)
//	if err != nil {
//		fmt.Println("建立连接失败：", err)
//		goto LOOP
//	}
//	fmt.Println("建立连接成功：", conn.RemoteAddr())
//	// 开始验证身份
//	write, err := conn.Write([]byte("Voter"))
//	if err != nil {
//		fmt.Printf("向[ %v ]发送失败 : err:[ %v ] : 发送的字节数:[ %v ]\n", conn.RemoteAddr(), err, write)
//		return
//	}
//
//	fmt.Println("开始与服务器通话：")
//
//	data := make([]byte, 0) // 进入对话模式
//	for {
//		buf := make([]byte, 1024)
//		fmt.Printf("等待[ %v ]的输入\n", conn.RemoteAddr())
//		read, err := conn.Read(buf)
//		if err != nil {
//			fmt.Printf("从[ %v ]接收失败 : err:[ %v ] : 接收的字节数:[ %v ]\n", conn.RemoteAddr(), err, read)
//			return
//		}
//		buf = buf[:read]
//		fmt.Printf("接收到的数据[byte]: [ %v ]\n", buf)
//		data = append(data, buf...)
//
//		str := string(buf)
//		fmt.Printf("接收到的数据[char]: [ %v ]\n", str)
//		if strings.Contains(str, "请等待公证人分发选票") {
//			fmt.Println("[本地]正在等待公证人分发选票")
//		} else if strings.Contains(str, "请开始接收选票") {
//			fmt.Println("[ 本地 ]:开始接收选票数据")
//			TicketJson := make([]byte, 1024*10)
//			read, err := conn.Read(TicketJson)
//			if err != nil {
//				fmt.Printf("接收的数据有误:[ %v ][ %v ]\n", err, read)
//				return
//			}
//			err = json.Unmarshal(TicketJson, &Ticket)
//			if err != nil {
//				fmt.Printf("转换为选票错误[ %v ]\n", err)
//			}
//		}
//	}
//}

//func RecvTicket(conn net.Conn) {
//	buf := make([]byte, 1024*10)
//	read, err := conn.Read(buf)
//	if err != nil {
//		fmt.Printf("从[ %v ]接收数据失败[ %v ][ %v ]\n", conn.RemoteAddr(), err, read)
//		return
//	}
//	buf = buf[:read]
//	fmt.Println("接收到的选片数据:", buf)
//
//	err = json.Unmarshal(buf, &Ticket)
//	if err != nil {
//		fmt.Println("反射选片失败：", err)
//		return
//	}
//	fmt.Println("反射选票成功")
//	fmt.Println(Ticket)
//}
