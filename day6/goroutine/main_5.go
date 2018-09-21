package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) //指定1核运算 返回先前设置的cpu数
	fmt.Println("n=", n)

	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
