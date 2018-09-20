package main

//定时重置
import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(6 * time.Second)
	timer.Reset(1 * time.Second) //重置延迟时间
	<-timer.C
	fmt.Println("end")
}
