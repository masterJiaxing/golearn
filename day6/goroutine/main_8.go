package main

//无缓冲区的channel案例
import (
	"fmt"
)

func main() {
	//创建一个无缓冲区管道
	ch := make(chan int, 0)

	//len(ch)缓冲区剩余的个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程：i=%d\n", i)
			ch <- i //往chan写内容
			fmt.Println("子协程 channel data")
		}
	}()

	//延时
	//time.Sleep(2*time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println("--------")
		num := <-ch //读取管道内容，没有内容前，阻塞
		fmt.Println("num=", num)
	}
}

func e(a *[]string, b string) {
	*a = append(*a, b)
}
