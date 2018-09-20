package main

//字符串常用函数
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("hellogo", "hello")) //检查字符串是否存在 返回布尔值
	fmt.Println(strings.Contains("hellogo", "adc"))

	fmt.Println(strings.Join([]string{"aa", "bb", "cc"}, ",")) //将切片转化为字符串sep为连接符

	fmt.Println(strings.Split("aa,bb,cc", ",")) //将s字符串按照sep分割返回slice

	fmt.Println(strings.Index("aweawf", "awf")) //查找字符串首次出现的位置，找不到返回-1

	fmt.Println(strings.Repeat("adv", 5)) //重复s字符串count次

	fmt.Println(strings.Replace("abcdefeg", "e", "xx", 2)) //在s字符串中，把old字符串替换为new字符串，n表示替换的次数，0为全部替换

	fmt.Println(strings.Trim("abcdefa", "a")) //在字符串的头部和尾部去掉含有cutset的字符

	fmt.Println(strings.Fields("a b c de f a")) //去除s字符串的空格符并且按照空格分割返回slice
}
