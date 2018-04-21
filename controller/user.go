package controllers

import (
	"github.com/astaxie/beego"
	"ingat_utang/service"
	"github.com/juusechec/jwt-beego"
	"time"
	"ingat_utang/http"
	"ingat_utang/util"
	_ "github.com/astaxie/beego/cache/memcache"
	"github.com/patrickmn/go-cache"
)

type UserController struct {
	Controller
}

type LoginController struct {
	beego.Controller
}

func (that *UserController) Get() {
	that.Data["json"] = http.Result(service.UserFindAll(), http.StatusOk)
	that.ServeJSON()
}

func (that *UserController) Post() {
	that.Data["json"] = http.Result(service.UserFindAll(), http.StatusOk)
	that.ServeJSON()
}

func (that *LoginController) Post() {
	username := that.GetString("username")
	password := that.GetString("password")

	user := service.FindByUsername(username, password)
	if username == user.Username && password == user.Password {
		et := jwtbeego.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 60, //Second
		}

		user.Token, _ = et.GetToken()
		util.UserCache().Set(user.Token, username, cache.DefaultExpiration)
		beego.Debug(util.UserCache().Get(user.Token))

		that.Data["json"] = http.Result(user, http.StatusOk)
	} else {
		that.Data["json"] = http.Result(user, http.StatusLoginError)
	}

	that.ServeJSON()
}

func (that *LoginController) Get() {
	token := that.Ctx.Request.Header.Get("Authorization")
	if _, found := util.UserCache().Get(token); found {
		util.UserCache().Delete(token)
	}

	that.Data["json"] = http.Result(nil, http.StatusOk)
	that.ServeJSON()
}
