package controllers

import (
	"encoding/json"
	"ingat_utang/model"
	"ingat_utang/service"
	"ingat_utang/http"
	"github.com/astaxie/beego"
)

type DebtController struct {
	beego.Controller
}

func (that *DebtController) Post() {
	that.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", that.Ctx.Request.Header.Get("Origin"))

	var debt entity.Debt
	json.Unmarshal(that.Ctx.Input.RequestBody, &debt)

	debt = service.DebtSave(debt)
	if debt.Id != 0 {
		that.Data["json"] = http.Result(debt, http.StatusSaveSuccess)
	} else {
		that.Data["json"] = http.Result(debt, http.StatusSaveFailed)
	}

	that.ServeJSON()
}

func (that *DebtController) Get() {
	var from, _ = that.GetInt("from")
	var to, _ = that.GetInt("to")
	debt := entity.Debt{}

	if from != 0 {
		debt.DebtFrom = &entity.User{Id:from}
	}else if from != 0 {
		debt.DebtTo = &entity.User{Id:to}
	}


	debt.DebtFrom = &entity.User{Id:1}
	that.Data["json"] = http.Result(service.DebtFind(debt), http.StatusOk)
	that.ServeJSON()
}
