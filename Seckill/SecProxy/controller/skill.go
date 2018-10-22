package controller

import "github.com/astaxie/beego"

type SkillController struct {
	beego.Controller
}

func (p *SkillController) Seckill(){
	p.Data["json"] = "sec Kill"
	p.ServeJSON()
}

func(p *SkillController) SeckInfo(){
	p.Data["json"] = "sec info"
	p.ServeJSON()
}
