package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"golearn/day9/beego/golang/models"
	_ "golearn/day9/beego/golang/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var user models.User
	user.Id = 2
	_ =user.Find()
	fmt.Printf("%+v",user)
	c.Ctx.WriteString("1")
}
