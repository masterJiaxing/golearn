package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"golearn/day9/beego/golang/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	var users []models.User
	id :=models.Create(&models.User{Username:"Username",Password:"pass"})
	//models.Update(&models.User{Id:int(id),Username:"Username",Password:"Update pass"})
	fmt.Println("Create new :"+strconv.Itoa(int(id)))
	models.Get(&users)
	//fmt.Printf("data = %+v",users[0].Person)
	c.Data["Users"] = users
	c.TplName = "index.tpl"
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