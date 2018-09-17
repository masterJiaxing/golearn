package main
//指针类型和普通类型的方法集
import "fmt"

type Person1 struct{
	name     string
	sex      byte
	age      int
}

func (p Person1) SetInfoPointer(){
	fmt.Println("SetInfoPointer")
}

func (p Person1) SetInfoValue(){
	fmt.Println("SetInfoValue")
}

func main(){
	//指针变量，他能调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person1{"mike", 'f',18}
	p.SetInfoPointer()

	//内部做了转化，先把指针p转换成*p然后再调用
	p.SetInfoValue()

	p1 := Person1{"mike", 'm',22}
	//内部做了转换，先把p转换成&p然后再调用
	p1.SetInfoPointer()
	p1.SetInfoValue()
}

