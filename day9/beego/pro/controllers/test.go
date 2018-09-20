package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	c.Ctx.WriteString("this is test method get!")
}

func (c *TestController) Test() {
	c.Ctx.WriteString("this is test method post!")
}
