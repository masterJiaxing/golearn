package router

import (
	"github.com/astaxie/beego"
	"golearn/Seckill/SecProxy/controller"
)

func init(){
	beego.Router("/seckill", &controller.SkillController{}, "*:Seckill")
	beego.Router("/seckinfo", &controller.SkillController{}, "*:SeckInfo")
}