package main

import (
	"fmt"
	"golearn/day3/packagevisibility/test"
)

func main() {
	s := test.Stu{1, "bj"}
	fmt.Println(s)
	test.Myfunc()
}
