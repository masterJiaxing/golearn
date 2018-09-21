package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//带有接收者，函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp=", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体
	var p2 Person
	//(&p2).SetInfo("go", 'f', 22)
	p2.SetInfo("go", 'f', 22)
	fmt.Println(p2)
}

//方法的基本类型不能是接口或指针
//错误演示 type pointer *int
type pointer int

func (tmp pointer) test() {

}

//不支持重载(名字相同但是不同参数的方法)
//错误演示
/*func (tmp pointer) test(n int){

}*/

type long int

func (tmp long) test() {

}
