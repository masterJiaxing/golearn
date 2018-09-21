package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego/session/mysql"
)

var (
	db orm.Ormer
)

type User struct {
	Id int
	Username string
	Password string
}

func init(){
	orm.Debug=true
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset")
	orm.RegisterModel(new(User))
	db = orm.NewOrm()
}

func ReadUser(users *[]User){
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}