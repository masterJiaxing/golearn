package main

import "fmt"

//面向过程
func Add01(a,b int)int{
	return a+b
}

//面向对象 方法:给某个类型绑定一个函数
type long int
//temp 叫接收者，接收者就是传递一个参数
func (temp long) Add02(a long) long{
	return temp +a
}

func main(){
	var result int
	result = Add01(1,1)
	fmt.Println(result)
	//定义一个变量
	var a long = 2
	//调用方法格式变量名.函数（所需参数）
	r := a.Add02(3)
	fmt.Println(r)
}
