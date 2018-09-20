package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true) //
	//第二个参数表示要追加的数，第三个为制定的10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcd")
	fmt.Println("slice=", string(slice))

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println(str)

	str = strconv.Itoa(666)
	fmt.Println("str=", str)

	//字符串转换成其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag=", flag)
	} else {
		fmt.Println("err=", err)
	}

	//把字符串转换成整形
	a, _ := strconv.Atoi("10")
	fmt.Println(a)
}
