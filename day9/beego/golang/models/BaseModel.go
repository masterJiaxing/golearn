package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)
type BaseModel struct {

}
//定义mysql连接
func init(){
	orm.Debug,_ = beego.AppConfig.Bool("mysqldebug")//是否开启调试模式

	dataSource := beego.AppConfig.String("mysqluser")+":"
	dataSource += beego.AppConfig.String("mysqlpass")+"@/"
	dataSource += beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Asia%2FShanghai"
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	orm.RegisterDataBase("default", "mysql", dataSource,10)
}
//遇到错误打印到控制台，没错误不打印
func (*BaseModel) ModelError(errMessage string, err error) string{
	if err != nil{
		beego.Error("【"+errMessage + "】 -> ",err)
		return errMessage
	}
	return ""
}

//查询单条记录
func (*BaseModel) baseFind(m *User) (error string) {
	o := orm.NewOrm() //NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存
	err := o.Read(m)
	if err == orm.ErrNoRows {
		return Base.ModelError("为查询到该条记录 -> ",err)
	} else if err == orm.ErrMissPK {
		return Base.ModelError("找不到主键 -> ",err)
	}
	return ""
}

//查询列表记录
func (*BaseModel) baseGet(m *User) (error string) {
	o := orm.NewOrm() //NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存
	err := o.Read(m)
	if err == orm.ErrNoRows {
		return Base.ModelError("为查询到该条记录 -> ",err)
	} else if err == orm.ErrMissPK {
		return Base.ModelError("找不到主键 -> ",err)
	}
	return ""
}






