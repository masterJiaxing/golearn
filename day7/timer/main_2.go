package main

//ticker的介绍
import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(1 * time.Second)

	i := 0
	for {
		var x = <-timer.C
		i++
		fmt.Println("x=", x)
		fmt.Println("i=", i)
		if i == 5 {
			timer.Stop() //停止定时器
			break
		}
	}
}
