package main

import "fmt"

func main(){
	array := []int{0,1,2,3,4,5,6,7,8,9}
	//[low:heigh:max] 取下标从low开始的元素,len=heigh-low, cap = max-low

	s1 := array[:]//[0:len(array):cap(array)]不指定容量和长度一样
	fmt.Println("s1=",s1)
	fmt.Printf("len=%d,cap=%d\n",len(s1), cap(s1))
	//操作某个元素，和数组操作方式一样
	data := array[1]
	fmt.Println("data=", data)

	s2 := array[3:6:7]
	fmt.Println("s2=", s2)
	fmt.Printf("len=%d,cap=%d\n",len(s2), cap(s2))

	s3 := array[0:6]
	fmt.Println("s3=", s3)
	fmt.Printf("len=%d,cap=%d\n",len(s3), cap(s3))

	s4 := array[3:]
	fmt.Println("s4=", s4)
	fmt.Printf("len=%d,cap=%d\n",len(s4), cap(s4))


	a := []int{0,1,2,3,4,5,6,7,8,9}

	s11 := a[2:5]
	s11[1] = 666
	fmt.Println("s11=",s11)
	fmt.Println("a=",a)

	s22 := s11[2:7]
	s22[2] = 777
	fmt.Println("s22=",s22)
	fmt.Println("a=",a)

}
