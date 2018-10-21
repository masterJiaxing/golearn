package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	Username string
	Password string

}
var Base BaseModel //全局baseModel
func init(){
	TablePrefix := beego.AppConfig.String("mysqlprefix")//获取表前缀
	orm.RegisterModelWithPrefix(TablePrefix, new(User))//注册带表前缀的表
	Base = BaseModel{}//初始化baseModel
}

func (m *User) Find() (error string) {
	return Base.baseFind(m)
}

func (m *User) Get() (error string) {
	return Base.baseGet(m)
}








