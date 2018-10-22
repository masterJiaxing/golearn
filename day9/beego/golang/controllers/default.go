package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"golearn/day9/beego/golang/models"
<<<<<<< HEAD
	_ "golearn/day9/beego/golang/models"
=======
	"strconv"
>>>>>>> d0e64749b3187d3f30b70f07054bd5d786af0f21
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
<<<<<<< HEAD
	var user models.User
	user.Id = 2
	_ =user.Find()
	fmt.Printf("%+v",user)
	c.Ctx.WriteString("1")
=======

	var users []models.User
	id :=models.Create(&models.User{Username:"Username",Password:"pass"})
	//models.Update(&models.User{Id:int(id),Username:"Username",Password:"Update pass"})
	fmt.Println("Create new :"+strconv.Itoa(int(id)))
	models.Get(&users)
	//fmt.Printf("data = %+v",users[0].Person)
	c.Data["Users"] = users
	c.TplName = "index.tpl"
>>>>>>> d0e64749b3187d3f30b70f07054bd5d786af0f21
}

func (c *MainController) Orm(){
	var users []models.User
	models.Get(&users)
	for _, data := range users {
		fmt.Printf("data = %+v\n",data)
	}

/*	var persons []models.Person
	models.GetPerson(&persons)
	for _, data := range persons {
		fmt.Printf("data = %+v\n",data)
	}
*/
	c.Data["Users"] = users
	c.TplName = "index.tpl"
}