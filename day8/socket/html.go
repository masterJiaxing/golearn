package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name  string
	Title string
	Age   string
}

func main() {
	t, err := template.ParseFiles("E:/go/goweb/golearn/day8/socket/i.html")

	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	p := Person{"Many", "个人网站", "15"}

	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("there was an error:", err)
	}
}
