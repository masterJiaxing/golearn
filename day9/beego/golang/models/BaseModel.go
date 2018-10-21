package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego/session/mysql"
)

type BaseModel struct {
}

func init(){
	orm.Debug=true
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8",30)
}
