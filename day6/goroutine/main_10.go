package main

import "fmt"

//range和close案例
func main() {
	closes()
	ranges()
}

//range 案例
func ranges() {
	ch := make(chan int) //创建一个无缓存区的channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往管道写数据
		}
		close(ch) //不加close管道不会关闭，一直等待接收导致死循环
	}()

	for num := range ch { //遍历管道内的数据
		fmt.Println("num=", num)
	}
}

//close 案例
func closes() {
	ch := make(chan int) //创建一个无缓存区的channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往管道写数据
		}
		close(ch) //不加close管道不会关闭，一直等待接收导致死循环
	}()

	for {
		//如果ok为true 说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num=", num)
		} else { //管道关闭
			break
		}
	}
}
