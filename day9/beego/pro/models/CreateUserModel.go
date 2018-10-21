package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	"time"
)

var (
	db orm.Ormer
)

type User struct {
	Id int			`orm:"auto;pk"`//auto:自动增加pk:主键null:默认为null
	Username string `orm:"null;size(64)"` //size=varchar
	Password string `orm:"null;size(64)"`
	CreatedAt time.Time `orm:"auto_now;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
}

func init(){
	orm.Debug,_ = beego.AppConfig.Bool("mysqldebug")//是否开启调试模式
	TablePrefix := beego.AppConfig.String("mysqlfrefix")//获取表前缀
	dataSource := beego.AppConfig.String("mysqluser")+":"
	dataSource += beego.AppConfig.String("mysqlpass")+"@/"
	dataSource += beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Asia%2FShanghai"
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	orm.RegisterDataBase("default", "mysql", dataSource,10)
	orm.RegisterModelWithPrefix(TablePrefix, new(User))//注册带表前缀的表
	db = orm.NewOrm() //NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存

}
//自定义表名
func (u *User) TableName() string {
	TablePrefix := beego.AppConfig.String("mysqlprefix")//获取表前缀
	return TablePrefix+"user"
}

// 多字段索引
func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id"},
	}
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}











