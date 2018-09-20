package main

//子协程的关闭   主线程关闭后，子协程会随之关闭
import (
	"fmt"
	"time"
)

func main() {
	//匿名方法
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i=", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
