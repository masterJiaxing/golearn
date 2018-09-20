package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net listen err =", err)
		return
	}
	defer listenner.Close()

	//阻塞监听用户连接
	conn, err := listenner.Accept()
	if err != nil {
		fmt.Println("listenner accept err=", err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err =", err)
		return
	}

	fileName := string(buf[:n])

	//回复ok
	conn.Write([]byte("ok"))

	//调用保存文件的方法
	RecvFile(fileName, conn)
}

//接收文件
func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os create err =", err)
		return
	}
	buf := make([]byte, 1024*4)
	//接受多少写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接受完毕")
			} else {
				fmt.Println("conn read err =", err)
			}
			return
		}
		f.Write(buf[:n])
	}
}
