package main

import "fmt"

func main() {
	const LENGTH int = 10
	const WIDTH = 5

	var area int

	area = LENGTH * WIDTH
	fmt.Println(area)

	const a, b, c = 1, false, "str"
	fmt.Println(a, b, c)

	const (
		Unknown = iota
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown, Female, Male)
}
