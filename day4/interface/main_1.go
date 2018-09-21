package main

//接口的定义和实现
import "fmt"

type Humaner interface {
	//方法只有声明，没有实现，由别的类型实现
	sayHi()
}

type Student struct {
	name string
	id   int
}

//Student实现此方法
func (tmp Student) sayHi() {
	fmt.Println(tmp)
	fmt.Printf("Student[%s,%d] sayHi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s,%d] sayHi\n", tmp.addr, tmp.group)
}

type Mystr string

func (tmp *Mystr) sayHi() {
	fmt.Printf("mystr[%s] sayHi\n", *tmp)
}

func main() {
	//定义接口类型变量
	var i Humaner
	s := Student{"mike", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str Mystr = "hello world"
	i = &str
	i.sayHi()
}
