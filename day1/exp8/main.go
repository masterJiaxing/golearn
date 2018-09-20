package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "day1/exp8/a.txt"
	//读取文件
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}
