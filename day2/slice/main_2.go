package main

import "fmt"

//切片和数组的区别以及创建
func main() {
	//切片和数组的区别
	//数组里面[]的长度固定是一个常量，数组不能修改长度，len和cap固定
	a := [5]int{}

	fmt.Printf("len=%d,cap=%d\n", len(a), cap(a))

	s := []int{}
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 13) //切片添加值
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))

	//切片创建
	//1自动推倒类型，同时初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1=", s1)
	//借助make 创建切片 (切片类型，长度，容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2))

	//借助make 创建切片  没有指定容量 容量和长度一致(切片类型，长度，没有指定容量)
	s3 := make([]int, 5)
	s3[4] = 1
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3))

	s4 := make([]string, 5)
	s4 = append(s4, "w")
	fmt.Println(s4)
	fmt.Printf("len=%d,cap=%d\n", len(s4), cap(s4))
}
