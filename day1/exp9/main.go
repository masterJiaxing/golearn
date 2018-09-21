package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*var a int = 0

	switch a {
		case 0:
		fmt.Printf("a is equals 0 ")
		fallthrough
	case 10,11:
		fmt.Println("a is equals 10")
	default:
		fmt.Println("a is equals default")
	}

	g := grad(1)
	fmt.Println(g)*/
	grad1()
}

func grad(score int) string { /*传入int 返回string*/
	g := ""

	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score:%d", score)) /*返回错误*/
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score < 100:
		g = "a"
	}
	return g
}

func grad1() {
	var n int
	n = rand.Intn(100)

	for {
		var input int
		fmt.Scanf("%d", &input)
		flag := false
		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("small")
		}

		if flag {
			break
		}
	}
}
