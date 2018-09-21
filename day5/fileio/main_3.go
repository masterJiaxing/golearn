package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件的读取

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //每次读取2k大小
	//n代表从文件读取内容的长度
	n, err := f.Read(buf)
	if err != nil && err != io.EOF { //EOF 文件出错同时没有读到结尾
		fmt.Println("err=", err)
	}
	fmt.Println("buf=", string(buf[:n]))
}

func main() {
	path := "demo.txt"
	ReadFile(path)
	ReadFileLine(path)
}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()
	//新建一个缓冲区，把内容放入缓冲区中
	r := bufio.NewReader(f)

	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err=", err)
		}
		fmt.Println("buf=#%s#\n", string(buf))
	}
}
