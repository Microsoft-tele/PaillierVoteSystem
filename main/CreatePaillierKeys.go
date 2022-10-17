package main

import (
	"SockGo/CryptoUtils"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var bitNumStr string
	flag.StringVar(&bitNumStr, "n", "1024", "Paillier密钥的位数")
	flag.Parse()
	bitNum, err := strconv.Atoi(bitNumStr)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	fmt.Println(bitNum)

	fmt.Println("生成新的Paillier密钥对:")
	CryptoUtils.CreateKeys(bitNum)
}
