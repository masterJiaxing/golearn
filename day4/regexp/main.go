package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "43.15 523 awfawef 1.54 1.2 6. 2.9 1111 5.55 1.24"

	//1)写一个正则 ， +匹配前一个字符1次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
	}

	//提取关键信息
	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1) //返回切片
	fmt.Println(result)

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := data[8:]

	s2 := data[:5]

	copy(s2, s1)
	fmt.Println(data)
}
