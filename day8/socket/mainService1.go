package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("err=", err)
			return
		}

		//处理用户发送的数据
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect successful")

	buf := make([]byte, 2048)
	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("errRead=", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		fmt.Printf("[%s]:%s\n", string(buf[:n]), "exit")
		if "exit" == string(buf[:n-2]) {
			fmt.Println(addr, "exit")
			return
		}

		//把数据转换成大写发送给对应的用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
