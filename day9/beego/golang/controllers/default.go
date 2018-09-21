package controllers

import (
	"github.com/astaxie/beego"
	"golearn/day9/beego/golang/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var users []models.User
	models.ReadUser(&users)
	c.Data["Users"] = users
	c.Data["len"] = len(users)
	c.TplName = "index.tpl"
}
