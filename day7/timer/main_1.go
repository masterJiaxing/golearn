package main

//Timer的使用和实现延时功能
import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间2s，2s后往time通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间", time.Now())

	//2s后往timer.C 写数据，有数据后，就可以读取
	t := <-timer.C //channel 如果没有数据会阻塞
	fmt.Println("t=", t)

	//延迟功能
	var x = <-time.After(2 * time.Second) //定时2s，阻塞2s，2s后产生一个事件，往channel写内容
	fmt.Println("时间到", x)
}
