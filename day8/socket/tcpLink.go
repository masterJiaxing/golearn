package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin) //从客户端读取
	for {
		input, _ := inputReader.ReadString('\n')    //以\n为分隔符
		trimmedInput := strings.Trim(input, "\r\n") //去掉、\r\n

		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput)) //往客户端写
		if err != nil {
			return
		}
	}

}
