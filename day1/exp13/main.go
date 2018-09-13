package main

import "fmt"
/*
多个int值相加
*/
func add(a int, arg... int) int{
	var sum int = a

	for i:=0;i<len(arg);i++{
		sum+= arg[i]
	}
	return sum;
}
/*
字符拼接   这里的返回值意思是返回一个类型是string的result
*/
func concat(a string, arg... string)(result string){
	result = a
	for i :=0;i<len(arg);i++{
		result += arg[i]
	}
	return result
}
func main(){
	sum := add(10, 2,3,4,5)
	fmt.Println(sum)

	res := concat("hello", "," ,"world")
	fmt.Println(res)
}
