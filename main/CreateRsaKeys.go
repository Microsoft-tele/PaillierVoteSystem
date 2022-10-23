package main

import (
	"RemoteRouter/RSAUtils"
	"flag"
	"strconv"
)

func main() {
	var bit string
	flag.StringVar(&bit, "n", "1024", "RSA密钥位数")
	flag.Parse()
	intBit, _ := strconv.Atoi(bit)
	RSAUtils.GenerateRSAKey(intBit)
}
