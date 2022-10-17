package ConveyUtils

import (
	"fmt"
	"net"
	"strings"
)

func ConveyData(conn net.Conn, data []byte) {
	n, err2 := conn.Write(data)
	if err2 != nil {
		fmt.Println("Write err is", err2)
	}
	fmt.Printf("向[ %v ]发送了[ %d ]个字节的数据:[%v]\n", conn.RemoteAddr(), n, string(data))
}
func PrintStringToSock(conn net.Conn, data string) {
	ConveyData(conn, []byte(data))
	ConveyData(conn, []byte("_over"))
}
func RecvStringFrom(conn net.Conn) string {
	data := RecvFrom(conn)
	strData := string(data)
	strData = strData[:len(strData)-1]
	return strData
}
func RecvFrom(conn net.Conn) (data []byte) { // 接收到 _over 结束本次接收
	data = make([]byte, 0)
	for {
		buf := make([]byte, 1024)
		fmt.Printf("等待[%v]的输入\n", conn.RemoteAddr())
		n, err1 := conn.Read(buf) //阻塞，直到客户端发送消息
		buf = buf[:n]
		fmt.Printf("接收到的数据[ %v ] 字节数:[ %d ]\n", buf, n)
		bufStr := string(buf)
		fmt.Printf("转换为字符串[ %v ] 字节数:[ %d ]\n", bufStr, n)
		if strings.Contains(bufStr, "_over") {
			fmt.Println("_over is", buf)
			data = append(data, buf...)
			break
		}
		if err1 != nil {
			fmt.Println("服务端Read err is", err1)
			break
		}
		data = append(data, buf...)
		//fmt.Print(string(buf[:n])) // n is real data read from conn
	}
	data = data[:len(data)-5]
	fmt.Println("Recv content:", string(data))
	return
}
