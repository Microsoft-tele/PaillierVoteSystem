package reference

import (
	"SockGo/RSAUtils"
	"SockGo/ShellUtils"
	"fmt"
)

func main() {
	//生成密钥对，保存到文件

	message := []byte("hello world")
	//加密
	// 获取rsa私钥进行签名
	home := ShellUtils.GetOutFromStdout("echo $HOME")[0]
	out := ShellUtils.GetOutFromStdout("ls " + home + "/rsa/keys/")
	fmt.Println("out:")
	for i, v := range out {
		fmt.Printf("[%d : %v]\n", i, v)
	}
	fmt.Println("Choose public key:")
	var op int
	fmt.Scanf("%d", &op)
	pubpath := home + "/rsa/keys/" + out[op]
	cipherText := RSAUtils.RSA_Encrypt(message, pubpath)
	fmt.Println("加密后为：", cipherText)
	//解密
	fmt.Println("Choose private key:")
	fmt.Scanf("%d", &op)
	pripath := home + "/rsa/keys/" + out[op]
	plainText := RSAUtils.RSA_Decrypt(cipherText, pripath)
	fmt.Println("解密后为：", string(plainText))

}
