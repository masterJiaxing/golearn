package main

import "fmt"

func main(){
	//声明一个map
	var m1 map[int]string
	fmt.Println(m1)
	fmt.Println(m1==nil)
	//对应map只有len没有cap
	fmt.Println("len=", len(m1))

	//通过make创建一个map
	m2 := make(map[int]string)
	fmt.Println("m2=", m2)
	fmt.Println("len=",len(m2))

	//通过make创建一个map，可以指定长度，如果没有数据会显示0
	m3 := make(map[int]string, 2)
	m3[1] = "mike"
	m3[2] = "c++"
	m3[3] = "go"

	fmt.Println("m3=", m3)
	fmt.Println("len=",len(m3))

	//键值是唯一的
	m4 := map[int]string{1:"mike",2:"c++",3:"go"}
	fmt.Println("m4=", m4)
	fmt.Println("len=",len(m4))
}
