package controllers

import (
	"github.com/juusechec/jwt-beego"
	"net/http"
	"github.com/astaxie/beego"
	http2 "goprint/http"
	"ingat_utang/util"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) Prepare() {
	token := c.Ctx.Request.Header.Get("Authorization")

	if _, found := util.UserCache().Get(token); !found {
		c.Ctx.Output.SetStatus(http2.StatusForbidden)
		c.Data["json"] = http2.Result(nil, http2.StatusForbidden)
		c.ServeJSON()
	}

	return
}

func CheckToken(r *http.Request) bool {
	token := r.Header.Get("Authorization")
	et := jwtbeego.EasyToken{}
	valid, _, _ := et.ValidateToken(token)
	return valid
}
