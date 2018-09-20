package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	//赋值，如果已存在key值，修改内容
	fmt.Println(m1)

	m1[1] = "c++"
	m1[3] = "go" //追加，map底层自动扩容，和append类似
	fmt.Println(m1)

	//遍历 第一个返回值为key，第二个返回值为value，遍历是无序的
	for key, value := range m1 {
		fmt.Printf("%d====>%s\n", key, value)
	}

	//如何判断一个key是否存在
	//第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在就ok为true
	value, ok := m1[1]
	if ok == true {
		fmt.Println("m[1]=", value)
	} else {
		fmt.Println("key不存在")
	}

	//删除map里面的key
	//delete(m1, 1)//删除key为1的内容

	del(m1) //map作为函数参数，他是引入传递
	fmt.Println("m[1]=", m1)

}

func del(m map[int]string) {
	delete(m, 1)
}
