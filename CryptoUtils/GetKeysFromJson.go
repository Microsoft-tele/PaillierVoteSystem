package CryptoUtils

import (
	"SockGo/ConveyUtils"
	"SockGo/FileUtils"
	"SockGo/ShellUtils"
	"SockGo/paillier"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

func GetKeysFromJson(conn net.Conn) (key *paillier.PrivateKey) {
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	dirList := ShellUtils.GetOutFromStdout("ls " + "/" + home + "/paillier/keys")
	for i, v := range dirList {
		data := fmt.Sprintf("[%d : %v]\n", i, v)
		ConveyUtils.ConveyData(conn, []byte(data)) // 接收返回信息
	}
	ConveyUtils.ConveyData(conn, []byte("_over"))

	choice := 0 //获取密钥文件
	fmt.Println("请输入您本次投票使用的密钥：")

	for {
		ConveyUtils.ConveyData(conn, []byte("请输入您本次投票使用的密钥:(Y/n)\n")) // 接收返回信息
		ConveyUtils.ConveyData(conn, []byte("_over"))
		data := ConveyUtils.RecvFrom(conn)
		strData := string(data)
		strData = strData[:len(strData)-1]
		fmt.Println("strData:", strData)
		choice, err := strconv.Atoi(strData)
		if err != nil {
			ConveyUtils.ConveyData(conn, []byte("您的输入不合法，请重新输入\n")) // 接收返回信息
			ConveyUtils.ConveyData(conn, []byte("_over"))
		} else {
			break
		}
		fmt.Println("Choice:", choice)
	}

	//_, err := fmt.Scanf("%d", &choice)
	//if err != nil {
	//	fmt.Println("Scanf err:", err)
	//	return
	//}
	filepath := home + "/paillier/keys/" + dirList[choice]
	fmt.Println(filepath)
	var PrivateKey *paillier.PrivateKey

	PrivateKeysSlice := FileUtils.ReadFileContent(filepath)

	err := json.Unmarshal([]byte(PrivateKeysSlice[0]), &PrivateKey)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println("成功从文件中恢复Paillier密钥对")
	return PrivateKey
}
