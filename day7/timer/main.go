package main
//select的超时操作
import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func(){
		for {
			select{
				case num :=<- ch:
					fmt.Println("num=",num)
				case <-time.After(3*time.Second):
					fmt.Println("超时")
					quit <- true
			}
		}
	}()
	ch <- 1
fmt.Println("1")
	<- quit
	fmt.Println("程序结束")
}

