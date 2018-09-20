package main

//协程的创建
import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask() //创建一个协程
	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}
}
