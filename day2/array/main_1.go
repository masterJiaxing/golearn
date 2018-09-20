package main

import "fmt"

func main() {
	//定义一个数组，[10]int和[5]int是不同类型
	//【数字】这个数字作为数组元素的个数
	var a [10]int
	var b [5]int

	fmt.Printf("len(a)=%d,len(b)=%d\n", len(a), len(b))

	//注意：定义数组时指定的数据元素个数必须是常量
	// n := 10
	// var c [n]int // non-constant array bound n

	//操作数组元素，从0开始，最大下标到len()-1，下标可以是常量也可以是变量
	a[0] = 1
	i := 1
	a[i] = 2

	//赋值，每个元素
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	//打印
	//第一个返回下标第二个返回元素
	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}
}
