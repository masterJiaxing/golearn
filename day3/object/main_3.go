package main

import "fmt"

//结构指针类型匿名字段


type Persion2 struct {
	name string
	sex  byte
	age  int
}

type Student2 struct {
	*Persion2 //结构体匿名字段
	id int
	addr string
}

func main(){
	//声明结构体
	s := Student2{&Persion2{"mike", 'm',18},666,"go"}
	fmt.Printf("s=%+v\n", s)


	//先定义变量
	var s2 Student2
	s2.Persion2 = new (Persion2)//new 也是地址引入
	s2.name = "go"
	s2.sex='m'
	s2.age=18
	s2.id=222
	s2.addr="sz"
	fmt.Println(s2.name,s2.sex,s2.age,s2.id,s2.addr)

}
