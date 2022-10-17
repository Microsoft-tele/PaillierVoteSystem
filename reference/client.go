package reference

//func main() {
//	ServiceHostIp := "192.168.1.103"
//	ServiceHostPort := "8888"
//	address := ServiceHostIp + ":" + ServiceHostPort
//	conn, err := net.Dial("tcp", address)
//	if err != nil {
//		fmt.Println("Client dial err =", err)
//		return
//	}
//	fmt.Printf("连接成功：%v 服务器地址是：%v\n", conn, conn.RemoteAddr().String())
//
//	// 获取本地公钥准备传送
//	PrivateKey := CryptoUtils.GetKeysFromJson()   // 从文件中获取私钥和公钥
//	PublicKey := PrivateKey.PublicKey             // 创建单独的公钥，准备向服务器发送
//	PublicKeyJson, err := json.Marshal(PublicKey) // 转换成json准备在网络中进行发送
//	if err != nil {
//		fmt.Println("PublicKey to json err:", err)
//	}
//
//	ConveyUtils.ConveyData(conn, PublicKeyJson) //向服务器发送公钥
//	ConveyUtils.ConveyData(conn, []byte("_over"))
//
//	ConveyUtils.RecvOver(conn)
//}
