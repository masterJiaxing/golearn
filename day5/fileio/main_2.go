package main

//文件的写入操作
import (
	"fmt"
	"os"
)

//文件的写入操作

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//使用延迟关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		//格式化字符串
		buf = fmt.Sprintf("i=%d\n", i)
		fmt.Println("buf=", buf)

		n, err := f.WriteString(buf) //字符写入文件
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("n=", n)
	}
}

func main() {
	path := "demo.txt"
	WriteFile(path)
}
