package main

import "fmt"
/*
append 在原切片的末尾添加元素
*/
func main(){
	s1 := []int{}
	fmt.Printf("len=%d,cap=%d\n",len(s1), cap(s1))
	fmt.Println("s1=", s1)

	//在原切片的末尾添加元素
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)

	fmt.Printf("len=%d,cap=%d\n",len(s1), cap(s1))
	fmt.Println("s1=", s1)

	s2 := []int{1,2,3}

	fmt.Println("s2=", s2)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	s2 = append(s2, 5)
	fmt.Println("s2=",s2)

	fmt.Println("--------------------")

	//append函数会智能底层数组的容量增长，一旦超过底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据
	s := make([]int, 0,1) //容量为1
	fmt.Println(cap(s))
	oldCap := cap(s)
	for i:=0;i<20;i++{
		s = append(s, i)
		if newCap := cap(s);oldCap < newCap{
			fmt.Printf("cap:%d====>%d\n", oldCap, newCap)
			oldCap = newCap
		}
	}


		var a uint = 60
		var b uint = 13
		var c uint = 0

		c = a & b
		fmt.Printf("第一行 - c 的值为 %d\n", c )

		c = a | b       /* 61 = 0011 1101 */
		fmt.Printf("第二行 - c 的值为 %d\n", c )

	fmt.Printf("姓名：%q\n","陈家兴")
	fmt.Printf("性别：%q\n","男")
	fmt.Printf("手机号码：%d\n",13735885491)
}
