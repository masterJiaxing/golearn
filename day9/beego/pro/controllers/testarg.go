package controllers

import (
	"github.com/astaxie/beego"
)

type User struct {
	Id int
	Username string `form:"username"`
	Age string `form:"age"`
	Email string `form:"email"`
}

type TestArgController struct {
	beego.Controller
}

/*
//参数解析
func (c *TestArgController) Get() {
	id := c.GetString("id")
	c.Ctx.WriteString(id)

	name := c.Input().Get("name")
	c.Ctx.WriteString(name)
}*/
func (c *TestArgController) Get() {
	c.TplName = "form.tpl"
}
func (c *TestArgController) Post(){
	u :=User{}
	if err := c.ParseForm(&u);err !=nil{

	}
	c.Ctx.WriteString("username:"+u.Username+"age:"+u.Age+"email:"+u.Email)
}