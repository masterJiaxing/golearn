package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego/session/mysql"
)

type Users struct {
	Id int
	Username string
	Password string
}

var (
	db orm.Ormer
)

func init(){
	orm.Debug=true//是否开启调试模式，开启会打印出sql语句
	orm.RegisterDataBase("default","mysql","root:root@/test?charset=utf8",30)

	orm.RegisterModel(new(Users))//注册模型

	db = orm.NewOrm()
}

//添加一个对象

func AddUser(users *Users)(int64 ,error){
	id,err :=db.Insert(users)
	return id,err
}

//读取一个对象
func ReadUser(users *[]Users){
	qb,_:=orm.NewQueryBuilder("mysql")
	qb.Select("*").From("users")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}