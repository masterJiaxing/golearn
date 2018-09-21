package main

//Goexite的使用
import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("aaaaaaaaaaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbb")
	}()
	for {
	}
}

func test() {
	defer fmt.Println("cccccccccccccccc")

	runtime.Goexit() //终止所在协程
	fmt.Println("dddddddddddddd")
}
