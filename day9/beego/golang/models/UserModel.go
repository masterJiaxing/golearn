package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego/session/mysql"
)
const ModelName = "User"
type Person struct {
	Id int
	Username string
	Sex string
	Email string
	User *User `orm:"rel(one)"`

}
type User struct {
	Id int
	Username string
	Password string
	Person *Person `orm:"reverse(one)"`
}



func init(){
	orm.RegisterModel(new(Person),new(User))
}

func Find(users *User){
	o := orm.NewOrm()
	err := o.Read(users)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}
}

func Get(users *[]User){
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable(ModelName)
	qs.All(users)
}

func GetPerson(persons *[]Person){
	o := orm.NewOrm()
	var qs orm.QuerySeter
	qs = o.QueryTable("Person")
	qs.All(persons)
}

func Create(user *User) int64{
	o := orm.NewOrm()
	id, err :=o.Insert(user)
	if err !=nil{
		fmt.Println(err.Error())
	}
	return id
}

func Update(user *User) int64{
	o := orm.NewOrm()
	id, err :=o.Update(user)
	if err !=nil{
		fmt.Println(err.Error())
	}
	return id
}

