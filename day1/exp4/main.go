package main

import "fmt"

/*全局变量简写*/
var (
	aa = 22
	ss = "kkk"
	bb = true
)

/*多个变量单独声明*/
var a1, b, c, d int

func variableZeroValue() {
	//变量的声明格式：var<变量名称> <变量类型>
	var a int
	var b string
	var c bool

	//变量的赋值格式：<变量名称> = <表达式>
	a = 10

	fmt.Printf("%d %q %v", a, b, c) /*Printf %d打印数字 %q打印一个引用*/
}

func variableInitValue() {
	var a, b int = 2, 3
	var s string = "db"
	fmt.Println(a, b, s)
}

/*類型推斷*/
func variableInitTypeValue() {
	var a, b, s, c = 2, 3, "bd", true
	fmt.Println(a, b, s, c)
}

/*简写形式*/
func variableShorter() {
	a, b, s, c := 2, 3, "db", true
	fmt.Println(a, b, s, c)
}
func main() {
	variableZeroValue()
	variableInitValue()
	variableInitTypeValue()
	variableShorter()
}
