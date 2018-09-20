package main

//接口的继承和转换
import "fmt"

type Humaner3 interface {
	//方法只有声明，没有实现，由别的类型实现
	sayhi()
}

type Personer interface {
	Humaner3
	sing(lrc string)
}

type Student3 struct {
	name string
	id   int
}

//Student 实现了sayhi()
func (tmp *Student3) sayhi() {
	fmt.Printf("Studet[%s,%d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student3) sing(lrc string) {
	fmt.Println("student 在唱歌：", lrc)
}
func main() {
	//定义一个借口类型的变量
	var i Personer
	s := &Student3{"mike", 666}
	i = s
	i.sayhi()
	i.sing("起来")

	var iPro Personer //超集
	iPro = &Student3{"msi", 555}
	var ii Humaner3

	ii = iPro //超集可以转换成子集

	ii.sayhi()
}
