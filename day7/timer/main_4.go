package main

//介绍select语句
import "fmt"

func main() {
	ch := make(chan int)    //数据通道
	quit := make(chan bool) //程序是否结束

	//消费者，从channel读取内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num=", num)
		}
		quit <- true
	}()
	//生产者，产生数字，写入channel
	bonacci(ch, quit)
}

func bonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据流动   select每个case语句里必须是一个IO操作
		select {
		case ch <- x: //只写 如果成功写入数据，则进行该case处理语句
			x, y = y, x+y
		case flag := <-quit: //只读 如果成功读取数据则进行case处理语句
			fmt.Println("flag=", flag)
			return
		default:
			//如果上面都没有成功则进入default处理流程
		}
	}
}
