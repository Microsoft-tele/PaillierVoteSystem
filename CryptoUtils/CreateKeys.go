package CryptoUtils

import (
	"SockGo/ShellUtils"
	"SockGo/paillier"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func CreateKeys(bitNum int) (PrivateKey *paillier.PrivateKey) {
	seed := rand.Reader
	PrivateKey, err := paillier.GenerateKey(seed, bitNum)

	if err != nil {
		fmt.Println("Create keys err:", err)
	}
	marshal, err1 := json.Marshal(PrivateKey)
	if err1 != nil {
		fmt.Println("Keys to json err:", err1)
		return
	}

	fmt.Println("This is 公证人生成的 secret keys, don't explore them!")
	NowTime := strings.Split(time.Now().String(), " ")[:2]
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	filename := home + "/paillier/keys/" + NowTime[0] + "_" + NowTime[1] + ".json" // 需要改进
	fmt.Println(filename)

	file, err2 := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err2 != nil {
		fmt.Println("Open file err:", err2)
		return
	}
	defer func(file *os.File) {
		_, err := file.WriteString("\n")
		if err != nil {
			fmt.Println("Write \\n err:", err)
			return
		}
		err3 := file.Close()
		if err3 != nil {
			fmt.Println("Close file err:", err3)
		}
	}(file)
	_, err4 := file.Write(marshal)
	if err4 != nil {
		fmt.Println("Write marshal err:", err4)
		return
	}
	return
}
