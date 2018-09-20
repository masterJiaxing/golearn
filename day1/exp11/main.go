package main

import "fmt"

/*规定函数输入和返回为int*/
func add(a, b int) int {
	return a + b
}

func main() {
	d := add(1, 2)

	c := add /*函数赋给变量*/
	d = c(4, 5)

	fmt.Println(d)
}
