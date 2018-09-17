package main

import "fmt"

func testa(){
	fmt.Println("aaaa")
}

func testb (x int){

	defer func(){//defer 当函数调用结束时执行的一个函数
		if err:= recover();err !=nil{//recover函数用于拦截panic错误，使程序可以继续往下运行  当未发生panic时recover返回nil
			fmt.Println(err)
		}
	}()

	fmt.Println("bbbbb")
	//panic("this is a panic test")//显示调用panic，会导致程序中断
	var a[10]int
	a[x] = 11//当x为22导致数组越界，产生一个panic导致系统崩溃
}

func testc(){
	fmt.Println("cccccccc")
}

func main(){
	testa()
	testb(222)
	testc()
}
