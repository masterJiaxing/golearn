package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")
	listenner, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("net listen err =", err)
		return
	}
	defer listenner.Close()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("accept failed, err =", err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}

		fmt.Println(string(buf[0:n]))
	}
}
