package main

//复制文件，客户端
import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//提示输入文件
	fmt.Println("请输入要传入的文件")
	var path string
	fmt.Scan(&path)

	//获取文件名 info.name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os state err=", err)
		return
	}

	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net dial err=", err)
		return
	}
	defer conn.Close()

	//给对方发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn write err=", err)
		return
	}

	//接收对方的回复，如果回复ok， 说明对方准备好 可以发送文件

	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err=", err)
		return
	}

	if "ok" == string(buf[:n]) {
		//发送文件内容
		SendFile(path, conn)
	}
}

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os open err=", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)
	//读取文件内容，读多少发多少
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.read err =", err)
			}
			return
		}
		//发送数据
		conn.Write(buf[:n]) //给服务器发送数据
	}

}
