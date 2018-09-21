package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
	return
}

func main() {
	a := 5
	b := 10
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
