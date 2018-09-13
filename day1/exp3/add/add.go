package add

import "fmt"

var Name string = "hello world"
var Age int = 10

func init () {/*当其他包调用该包会自动执行init */
	Name = "hello go"
	Age  = 30
}


func Test(){
	fmt.Println("hello go")
}