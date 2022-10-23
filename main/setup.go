package main

import (
	"RemoteRouter/ShellUtils"
	"fmt"
)

func main() {
	// privateKeyPath := home + "/rsa/keys/" + NowTime[0] + "_" + NowTime[1] + "_pri.pem" // 需要改进
	home := ShellUtils.GetOutFromStdout("echo $HOME")
	comm := "mkdir " + home[0] + "/rsa"
	comm1 := "mkdir " + home[0] + "/rsa/keys"
	comm2 := "mkdir " + home[0] + "/paillier"
	comm3 := "mkdir " + home[0] + "/paillier/keys"
	fmt.Println(comm)
	fmt.Println(comm1)
	fmt.Println(comm2)
	fmt.Println(comm3)
	ShellUtils.GetOutFromStdout(comm)
	ShellUtils.GetOutFromStdout(comm1)
	ShellUtils.GetOutFromStdout(comm2)
	ShellUtils.GetOutFromStdout(comm3)
}
