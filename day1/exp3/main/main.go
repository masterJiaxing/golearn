package main

import (
	"fmt"
	addPackage "golearn/day1/exp3/add"  /* _ "golearn/day1/array/add" 这样写的意思是只初始化不引用*/

)

func main (){
	addPackage.Test()
	fmt.Println(addPackage.Name)
	fmt.Println(addPackage.Age)
}