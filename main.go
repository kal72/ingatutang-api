package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"ingat_utang/entity"
	"ingat_utang/controller"
)

func init() {
	// register model
	orm.RegisterModel(new(entity.User), new(entity.Debt))
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/ingatutang_db?charset=utf8", 30)
}

func main() {
	orm.Debug = true

	beego.Router("/users", &controllers.UserController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{})
	beego.Router("/debts", &controllers.DebtController{})
	beego.Run()
}
