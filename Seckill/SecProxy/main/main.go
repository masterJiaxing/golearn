package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "golearn/Seckill/SecProxy/router"
)



func main(){
	err := initConfig()
	if err !=nil{
		panic(err)
		return
	}

	err = initSec()
	if err !=nil{
		panic(err)
		return
	}
	fmt.Println("运行成功")

	beego.Run()
}