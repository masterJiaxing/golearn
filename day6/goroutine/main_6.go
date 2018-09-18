package main
//channel入门以及实现同步
import (
	"fmt"
	"time"
)
var ch = make(chan int)//创建管道
func Printer(str string){
	for _,data := range str{
		fmt.Printf("%c",data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1(){
	Printer("hello")
	ch <- 888
}

func person2(){
	<- ch //从管道里读取数据如果没有数据它就会阻塞
	Printer("world")

}

func main(){
	go person1()
	go person2()
	for{}
}


