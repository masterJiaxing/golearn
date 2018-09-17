package main

import "fmt"

type Student struct {
	id       int
	name     string
	sex      byte
	age      int
	addr     string
}

func main(){
	//顺序初始化，每个成员都要初始化
	var s1 Student = Student{1,"mike", 'm',18, "ah"}
	fmt.Println(s1)

	//指定成员初始化，没有初始化的成员，自动赋值为0
	var s2 Student = Student{name:"xiaohuang",addr:"bj"}
	fmt.Println(s2)

	//顺序初始化，每个成员都要初始化，别忘加&
	var  p1 *Student = &Student{1,"mike", 'm',18, "ah"}
	fmt.Println(p1)

	//指定成员初始化，没有初始化的成员，自动赋值为0，别忘加&
	p2 := &Student{name:"zhangsan",addr:"bj"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println(*p2)
}
