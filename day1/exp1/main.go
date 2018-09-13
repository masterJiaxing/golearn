package main /*定义包名，必须在源文件中非注释的第一行指明这个文件属于哪个包入package main 表示一个可独立执行的程序，每个go应用程序都包含一个名为main的包*/

import "fmt" /*告诉go编译器 这个程序需要fmt包（函数或者其他元素可能会被使用），fmt包实现了格式化Io的函数*/

func main() { /*程序的开始执行函数，每个可执行的程序必须包含main方法，第一个执行的函数*/
	fmt.Println("Hello world") /*可以将字符串输入到控制台 后面ln是换行 \n*/
}
