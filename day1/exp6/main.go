package main

import "fmt"

func main() {
	var a int = 50
	var b bool
	c := "a"

	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("90%%\n")
	fmt.Printf("%t\n", b)
	fmt.Printf("%b\n", 100)
	fmt.Printf("%f\n", 120.00)
	fmt.Printf("%q\n", "this is a test")
	fmt.Printf("%v\n", "this is a test")
	fmt.Printf("%x\n", 399922)
	fmt.Printf("%p\n", &a)

	str := fmt.Sprintf("a=%d", a)
	fmt.Printf("%q\n", str)
}
