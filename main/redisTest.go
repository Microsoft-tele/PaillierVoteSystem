package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接失败:", err)
	}
	name := "李为君"
	do, err := conn.Do("Lpush", "name", name)
	if err != nil {
		fmt.Println("获取数据失败:", err)
		return
	}
	fmt.Println(do)
}
