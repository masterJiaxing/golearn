package main

import (
	"errors"
	"fmt"
)

//异常处理error接口的使用

func main() {
	err1 := fmt.Errorf("%s", "this is normal err")
	fmt.Println(err1)

	err2 := errors.New("this is normal err2")
	fmt.Println(err2)

	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("result=", result)
		fmt.Println("err=", err)
	} else {
		fmt.Println("result=", result)
		fmt.Println("err=", err)
	}
}

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
		//err = fmt.Errorf("%s", "分母不能为0")//error的两种定义方式
	} else {
		result = a / b
	}
	return
}
