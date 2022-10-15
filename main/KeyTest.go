package main

import (
	"SockGo/CryptoUtils"
	"SockGo/paillier"
	"fmt"
	"math/big"
)

func main() {
	//CryptoUtils.CreateKeys(1024) // 向上层图形化界面提供调用接口
	PrivateKey := CryptoUtils.GetKeysFromJson()
	//fmt.Println(PrivateKey)

	m15 := new(big.Int).SetInt64(15)
	c15, err := paillier.Encrypt(&PrivateKey.PublicKey, m15.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	// Decrypt the number "15".
	d, err := paillier.Decrypt(PrivateKey, c15)
	if err != nil {
		fmt.Println(err)
		return
	}
	plainText := new(big.Int).SetBytes(d)
	fmt.Println("Decryption Result of 15: ", plainText.String())
}
