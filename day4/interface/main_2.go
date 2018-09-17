package main
//多态的实现
import "fmt"

type humaner interface{
	//方法只有声明，没有实现，由别的类型实现
	sayHi()
}

type student struct{
	name   string
	id     int
}

//Student实现此方法
func (tmp student) sayHi(){
	fmt.Println(tmp)
	fmt.Printf("Student[%s,%d] sayHi\n", tmp.name,tmp.id)
}

type teacher struct{
	addr string
	group string
}

func (tmp *teacher) sayHi(){
	fmt.Printf("Teacher[%s,%d] sayHi\n",tmp.addr,tmp.group)
}

type mystr string
func (tmp *mystr) sayHi(){
	fmt.Printf("mystr[%s] sayHi\n", *tmp)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以有不同表现，多态
func whoSayHi(i humaner){
	i.sayHi()
}

func main(){
	//定义接口类型变量
	//var i humaner
	s := student{"mike", 666}


	t := &teacher{"bj", "go"}


	var str mystr = "hello world"
//多态的实现1
	whoSayHi(s)
	whoSayHi(t)
	whoSayHi(&str)

	//多态的实现2   创建切片
	x := make([]humaner,3)
	x[0] = s
	x[1] = t
	x[2] = &str
	//第一个返回下标，第二个返回下表对应的值
	for _,i :=range x{
		i.sayHi()
	}
}

