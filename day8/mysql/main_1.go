package main
//数据更新、删除
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//go get github.com/go-sql-driver/mysql  下载sql驱动
//go get github.com/jmoiron/sqlx  下载操作sql 的渠道
var Db *sqlx.DB
func init(){
	database ,err :=sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err !=nil{
		fmt.Println("open mysql failed",err)
		return
	}
	Db = database
}
func main() {
	r,err := Db.Exec("insert into person(username,sex,email) values(?,?,?)","stu1","man", "stu@qq.com")//该函数支持更新和删除操作的sql

	if err != nil{
		fmt.Println("exec failed", err)
		return
	}

	id,err :=r.LastInsertId()
	if err !=nil{
		fmt.Println("exec failed",err)
		return
	}
	fmt.Println("insert succ:",id)
}
