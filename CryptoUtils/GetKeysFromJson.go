package CryptoUtils

import (
	"SockGo/FileUtils"
	"SockGo/ShellUtils"
	"SockGo/paillier"
	"encoding/json"
	"fmt"
)

func GetKeysFromJson() (key *paillier.PrivateKey) {
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	dirList := ShellUtils.GetOutFromStdout("ls " + "/" + home + "/paillier/keys")
	for i, v := range dirList {
		fmt.Printf("[%d : %v]\n", i, v)
	}
	choice := 0 //获取密钥文件
	fmt.Println("请输入您本次加密传输的密钥：")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("Scanf err:", err)
		return
	}
	filepath := home + "/paillier/keys/" + dirList[choice]
	fmt.Println(filepath)
	var PrivateKey *paillier.PrivateKey

	PrivateKeysSlice := FileUtils.ReadFileContent(filepath)

	err = json.Unmarshal([]byte(PrivateKeysSlice[0]), &PrivateKey)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	return PrivateKey
}
