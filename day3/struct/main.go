package main

import "fmt"

//定义结构体
type Student3 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var s1 Student3 = Student3{1, "mike", 'm', 18, "ah"}
	var s2 Student3 = Student3{1, "mike", 'm', 18, "ah"}
	var s3 Student3 = Student3{2, "mike", 'm', 18, "ah"}

	//结构体可以比较
	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s1==s3", s1 == s3)
	//同类型的2个结构体变量可以相互赋值
	var tmp Student3
	tmp = s3
	fmt.Println(tmp)

	fmt.Println("传入钱：", tmp)
	test1(tmp)
	fmt.Println("传入后test1：", tmp)
	test2(&tmp)
	fmt.Println("传入后test2：", tmp)
}

func test1(s Student3) { //值传递，形参无法改变实参
	s.id = 666
	fmt.Println("test01:", s)
}

func test2(s *Student3) { //地址传递（引用传递），形参可以改变实参
	s.id = 666
	fmt.Println("test01", s)
}
