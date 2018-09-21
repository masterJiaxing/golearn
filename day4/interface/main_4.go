package main

//空接口的介绍
import "fmt"

func xxx(...interface{}) {

}

func main() {
	//空接口是万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)

	var i1 interface{} = struct {
		X int
	}{1}
	fmt.Println(i1)

	i = "abc"
	fmt.Println("i=", i)
	xxx("aaa")
}
