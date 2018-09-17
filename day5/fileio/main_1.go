package main
//标准的输入输出
import (
	"fmt"
	"os"
)

func main (){
	//标准输出设备
	//os.Stdout.Close()//关闭后，无法输出，默认时候时打开的
	fmt.Println("are you ok?")
	os.Stdout.WriteString("are you ok")

	var a int
	fmt.Println("请输入a：")
	fmt.Scan(&a)//从标准输入设备中读取内容，放入到a中
	fmt.Println("a=",a)
}







