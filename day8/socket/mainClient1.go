package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	//接受服务器回复的数据，新任务

	go func() {
		//从键盘输入内容，给服务器发送
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) //从键盘读取内容，放在str

			if err != nil {
				fmt.Println("os,Stdin,err=", err)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])
		}
	}()

	//切片缓冲
	buf := make([]byte, 1024*2)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}

		fmt.Println(string(buf[:n])) //打印接受到的数据内容
	}
}
