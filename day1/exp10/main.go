package main

import "fmt"

func main() {
	Print(6)
	forDie()
}

func Print(n int) {
	for i := 0; i < n+1; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

func forDie() {
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("", i)
		i++
	}

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
