package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 asc tac"
	//1)解析规则，他会解析正则，如果成功返回解析器
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("regexp err")
	}

	//2)根据规则提取关键信息
	result := reg1.FindAllStringSubmatch(buf, 2) //n为-1时匹配所有
	fmt.Println(result)
}
