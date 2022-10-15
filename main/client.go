package main

import (
	"SockGo/ConveyUtils"
	"SockGo/CryptoUtils"
	"SockGo/paillier"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
)

func main() {
	ServiceHostIp := "192.168.1.103"
	ServiceHostPort := "8888"
	address := ServiceHostIp + ":" + ServiceHostPort
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Client dial err =", err)
		return
	}
	fmt.Printf("连接成功：%v 服务器地址是：%v\n", conn, conn.RemoteAddr().String())

	// 获取本地公钥准备传送
	PrivateKey := CryptoUtils.GetKeysFromJson()   // 从文件中获取私钥和公钥
	PublicKey := PrivateKey.PublicKey             // 创建单独的公钥，准备向服务器发送
	PublicKeyJson, err := json.Marshal(PublicKey) // 转换成json准备在网络中进行发送
	if err != nil {
		fmt.Println("PublicKey to json err:", err)
	}

	ConveyUtils.ConveyDataToService(conn, PublicKeyJson) //向服务器发送公钥
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))

	ConveyUtils.RecvOver(conn)

	m15 := new(big.Int).SetInt64(15) // 生成测试数据
	c15, err := paillier.Encrypt(&PrivateKey.PublicKey, m15.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	m20 := new(big.Int).SetInt64(20)
	c20, err := paillier.Encrypt(&PrivateKey.PublicKey, m20.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}

	ConveyUtils.ConveyDataToService(conn, c15)
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))

	ConveyUtils.RecvOver(conn)

	ConveyUtils.ConveyDataToService(conn, c20)
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))

	AddRes := ConveyUtils.RecvFromClient(conn)

	decrypt, err := paillier.Decrypt(PrivateKey, AddRes)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}
	fmt.Println("Result of 15 + 20 after decryption: ",
		new(big.Int).SetBytes(decrypt).String()) // 150
}
