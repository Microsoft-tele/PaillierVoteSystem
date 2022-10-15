package main

import (
	"SockGo/ConveyUtils"
	"SockGo/paillier"
	"encoding/json"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端的数据
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Conn close err", err)
		}
	}(conn)
	/**
	接收公钥
	*/
	DataOfPublicKey := ConveyUtils.RecvFromClient(conn) //收到   '_over'  就停止接受的函数
	var PublicKey paillier.PublicKey
	err := json.Unmarshal(DataOfPublicKey, &PublicKey)
	if err != nil {
		fmt.Println("Json to publicKey err:", err)
		return
	}
	fmt.Println("Public key:", PublicKey)
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))
	/**
	接收公钥完毕，并向客户返回over
	*/

	/**
	接收要处理的图片数据
	*/
	FirstData := ConveyUtils.RecvFromClient(conn)
	fmt.Println("FirstData:", FirstData)
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))
	/**
	接收要处理的图片数据完毕，并向客户返回over
	*/

	/**
	接收要隐藏的信息
	*/
	SecondData := ConveyUtils.RecvFromClient(conn)
	fmt.Println("SecondData:", SecondData)
	/**
	接收要隐藏的信息完毕，并向客户返回over
	*/

	// 接收到所有数据， 进行计算
	ADDFirSec := paillier.AddCipher(&PublicKey, FirstData, SecondData)
	fmt.Println("Res is :", ADDFirSec)

	ConveyUtils.ConveyDataToService(conn, ADDFirSec)
	ConveyUtils.ConveyDataToService(conn, []byte("_over"))
}

func main() {
	fmt.Println("服务器开始监听!")
	listen, err := net.Listen("tcp", "0.0.0.0:8888") //阻塞
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}
	defer func(listen net.Listener) {
		err1 := listen.Close()
		if err1 != nil {
			fmt.Println("Listen close err = ", err1)
		}
	}(listen)
	for {
		fmt.Println("等待客户端请求服务----")
		conn, err2 := listen.Accept() //阻塞
		if err2 != nil {
			fmt.Println("Accept err = ", err2)
		} else {
			fmt.Println("Conn is", conn)
			go process(conn) // 多线程服务
		}
	}
}
