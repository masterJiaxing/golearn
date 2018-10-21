package routers

import (
	"golearn/day9/beego/golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/orm", &controllers.MainController{}, "get:Orm")
}
