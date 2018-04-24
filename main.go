package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"ingat_utang/model"
	"ingat_utang/controller"
)

func init() {
	// register model
	orm.RegisterModel(new(entity.User), new(entity.Debt))
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		"user=baotpcfatibenj password=e0691cc86a656edcfdb4f2bb7fe3e93cf6527b1e88b90490ca64dce75d73d2b2 host=1ec2-107-22-168-211.compute-1.amazonaws.com port=5432 dbname=d93t5tq2lfk3t1 sslmode=disable");

	orm.RunSyncdb("default", false, true)
	//orm.RegisterDataBase("default", "pq", "postgres://baotpcfatibenj:e0691cc86a656edcfdb4f2bb7fe3e93cf6527b1e88b90490ca64dce75d73d2b2@ec2-107-22-168-211.compute-1.amazonaws.com:5432/d93t5tq2lfk3t1", 30)
}

func main() {
	orm.Debug = true

	beego.Router("/", &controllers.Controller{})
	beego.Router("/users", &controllers.UserController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{})
	beego.Router("/debts", &controllers.DebtController{})
	beego.Run()
}
