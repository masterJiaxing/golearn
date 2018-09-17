package main

import "fmt"
//匿名字段实现继承
type Persion struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Persion //只有类型，没有名字，匿名字段，集成了persion的成员
	id    int
	addr  string
}

func main(){
	//顺序初始化
	var s1 Student = Student{Persion{"mike",'m',19},1,"bj"}
	fmt.Println("s1=",s1)

	//自动推倒类型
	s2 := Student{Persion{"mike",'m',20},1,"bj"}
	fmt.Println(s2)
	//%+v,更详细的显示
	fmt.Printf("s2=%+v\n",s2)

	//指定成员初始化，没有初始化的自动赋值为0
	s3 := Student{id:1}
	fmt.Printf("s3=%+v\n",s3)

	//persion的值部分初始化
	s4 := Student{Persion:Persion{name:"mike"}}
	fmt.Printf("s4=%+v\n",s4)

	fmt.Println("-----------------------")
	//结构体匿名成员变量的操作方法
	//方式一
	s2.name = "zhangsan"
	s2.id   = 100
	s2.age   = 22
	s2.addr   = "sz"
	//方式二
	s2.Persion = Persion{"go", 'm', 18}
	fmt.Println(s2)
}
