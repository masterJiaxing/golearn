package main

import "fmt"

//定义结构体
type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//定义一个结构体的普通变量
	var s Student2
	//操作成员，需要使用.
	s.id = 1
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Println("s=", s)

	//指针有合法的指向后，才能操作成员
	//定义一个普通的结构体变量
	var s1 Student2
	//在定义一个指针变量，保存s的地址
	var p1 *Student2
	p1 = &s1

	//通过指针操作成员p1.id 和(*p1).id完全等价，只能用.运算符

	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 20
	p1.addr = "bj"
	fmt.Println("p1=", p1)
	fmt.Println("s1=", s1)

	//2通过new 申请一个结构体
	p2 := new(Student2)
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2=", p2)
}
