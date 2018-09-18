package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	fmt.Println(1)
	go f1(out)
	go f1(out)
	fmt.Println(2)

	out <- 11
	out <- 22
	fmt.Println(3)
}