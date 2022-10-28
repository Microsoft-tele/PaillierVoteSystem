package main

import(
	"fmt"
	"RemoteRouter/User"
)

func main(){
	fmt.Println("添加用户:")
	user := User.User{
		Username: "liweijun",
		Password: "liweijun",
		Email:    "123@gmail.com",
	}
	err := user.AddUser()
	if err != nil {
		fmt.Println("添加用户失败:", err)
		return
	}
}