package main
//Gosched的使用
import (
	"fmt"
	"runtime"
)

func Ins(){
	for i:=0;i<5;i++{
		fmt.Println("go")
	}
}

func main(){
	go Ins()

	for i:=0;i<2;i++{
		//让出时间片，先让别的协程执行，它执行完，再回来执行这里的代码
		runtime.Gosched()
		fmt.Println("hello")
	}
}
