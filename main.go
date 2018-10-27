package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "storeserv/routers"
)

func init() {
	beego.BConfig.CopyRequestBody = true
	orm.RegisterDataBase("default", "mysql",
		"root:ddddd@/store?charset=utf8", 30)
}

func main() {
	beego.Run()
}
