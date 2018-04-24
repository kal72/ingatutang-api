package controllers

import (
	"github.com/astaxie/beego"
	"ingat_utang/util"
	http2 "ingat_utang/http"
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

func (c *Controller) Get()  {
	c.Ctx.Output.SetStatus(http2.StatusForbidden)
	c.Data["json"] = http2.Result(nil, http2.StatusForbidden)
	c.ServeJSON()
}
