package routers

import (
	"github.com/astaxie/beego/context"
	"golearn/day9/beego/pro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//固定路由
    //beego.Router("/", &controllers.MainController{})
    //基础路由
    beego.Get("/1", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
    beego.Router("/", &controllers.MainController{},"get:Get;post:Test")
    beego.Router("/test", &controllers.TestController{},"get:Get;post:Test")
    beego.Router("/test/arg", &controllers.TestArgController{},"get:Get;post:Post")
    //测试orm
    beego.Router("/test/orm", &controllers.TestModelController{},"get:Get;post:Post")
}
