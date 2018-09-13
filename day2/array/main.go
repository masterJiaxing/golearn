package main

import "fmt"

func main(){
	a := [6]int {1,2,3,4,6,5}
	//数组作为参数，是值传递
	//实参数组的每个元素给形参拷贝一份
	//形参的数组是实参数组的复制品
	modify(a)

	modify1(&a)
	fmt.Println("main a=", a)
}

func modify(a [6]int){
	a[0] = 66666
	fmt.Println("modify a=", a)
}

//p 指向实现数组a,他是指向数组，他是指针数组，p指向a的地址&a,*p等价于a, a[0] = (*p)[0]
func modify1(p *[6]int){
	(*p)[0] = 66666
	fmt.Println("modify p=", *p)
}