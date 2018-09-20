package main

import "fmt"

//此通道只能写，不能读
func product(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

//此channel只能读，不能写
func consumer(a <-chan int) {
	for num := range a {
		fmt.Println("num=", num)
	}
}

func main() {
	//创建一个双向管道
	ch := make(chan int)
	//创建一个单向channel ，只用于写int
	//var ch2 chan<- int

	//创建一个单向channel，只用于读
	//var ch3 <-chan int

	//生产者，生产数据，写入channel
	go product(ch)

	//消费者，从channel读取内容,打印
	consumer(ch)
}
