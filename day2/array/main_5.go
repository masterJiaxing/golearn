package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//支持比较，比较2个数组类型要一致，只支持 == 和!=

	a :=[5]int{1,2,3,4,5}
	b :=[5]int{1,2,3,4,5}
	c :=[5]int{1,2,3}
	fmt.Println( "a==b", a==b)
	fmt.Println( "a==c", a==c)

	//同类型的数组可以赋值
	var d[5]int
	d = a
	fmt.Println("d=", d)

	//设置种子
	rand.Seed(time.Now().UnixNano()) //以当前系统时间作为种子
	for i:=0;i<5;i++{
		//产生随机数
		fmt.Println("rand=", rand.Int())//随机数
		fmt.Println("rand=", rand.Intn(100))//随机数
	}
}