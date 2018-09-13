package main

import "fmt"

type add_func func(int,int)int

func add(a,b int)int{
	return a+b
}

func sub(a,b int) int{
	return a-b
}

func operator(op add_func,a int,b int)int{
	return op(a,b)
}

/*多返回值*/
func div (a,b int)(int, int){
	return a/b,a%b
}

func eval(a,b int ,op string)(int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupport operation"+op)
	}
}
func main(){
	c := sub

	sum := operator(c, 5,6)

	fmt.Println(sum)

	q,p := div(9,4)
	fmt.Println(q)
	fmt.Println(p)

	if result,err := eval(4, 5, "1");err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println(result)
	}
}