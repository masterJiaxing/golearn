package controllers

import (
	"github.com/astaxie/beego"
	models "golearn/day9/beego/pro/models"
)

type TestModelController struct {
	beego.Controller
}

type Users struct {
	Id int
	Username string
	Password string
}
/*func init(){
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8",30)
	orm.RegisterModel(new(Users))
}*/

func(c *TestModelController) Get(){


	/*
	新增数据
	o := orm.NewOrm()
	users := Users{Username:"张三",Password:"123"}
	id,_:=o.Insert(&users)
	c.Ctx.WriteString(fmt.Sprintf("新增数据id=%v",id))*/

	/*
	更新操作
	user := Users{Username:"111",Password:"123"}
	user.Id=6
	o := orm.NewOrm()
	bo,_:=o.Update(&user)
	c.Ctx.WriteString(fmt.Sprintf("新增数据id=%v",bo))*/

	/*
	查询
	o := orm.NewOrm()
	user := Users{Id:7}

	o.Read(&user)
	c.Ctx.WriteString(fmt.Sprintf("查询user:%v", user))*/

	/*
	Model新增数据
	user := models.Users{Username:"战士",Password:"12234234234"}
	models.AddUser(&user)
	c.Ctx.WriteString("call model success")*/

	var users []models.Users
	models.ReadUser(&users)

	c.Data["Users"] = users
	c.Data["len"] = len(users)

	c.TplName = "test.tpl"
}