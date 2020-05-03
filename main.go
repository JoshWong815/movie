package main

import (
	"github.com/astaxie/beego/orm"
	_ "movie/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/mv?charset=utf8")
	beego.Run()
}

