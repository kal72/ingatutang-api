package service

import (
	"github.com/astaxie/beego/orm"
	"ingat_utang/entity"
	"github.com/astaxie/beego"
)

func UserSave(user entity.User) entity.User {
	o := orm.NewOrm()
	o.Insert(&user)
	return user
}

func UserFindAll() []*entity.User{
	o := orm.NewOrm()
	var users []*entity.User
	o.QueryTable("user").All(&users)
	return users
}

func FindByUsername(username string, password string) entity.User {
	o := orm.NewOrm()
	user := entity.User{Username: username}
	err := o.Read(&user, "username")
	if err != nil {
		beego.Error(err)
	}

	return user
}
