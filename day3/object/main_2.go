package main

import "fmt"

//面向对象的同名字段和其他匿名字段

type mystr string //自定义类型，给一个类型改名

type Persion1 struct {
	name string
	sex  byte
	age  int
}

type Student1 struct {
	Persion1 //结构体匿名字段
	int      //基础类型的匿名字段
	mystr    //自定义类型
}

func main() {
	//声明结构体
	s := Student1{Persion1{"mike", 'm', 18}, 666, "go"}
	fmt.Printf("s=%+v\n", s)

	s.Persion1 = Persion1{"go", 'm', 22}
	fmt.Println(s.name, s.age, s.int, s.mystr)

	s.Persion1.age = 23
	s.Persion1.name = "c"
	fmt.Println(s.Persion1, s.int, s.mystr)

}
