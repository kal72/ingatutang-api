package service

import (
	"ingat_utang/model"
	"github.com/astaxie/beego/orm"
)

func DebtSave(dept entity.Debt) entity.Debt {
	o := orm.NewOrm()
	o.Insert(&dept)
	return dept
}

func DebtFind(param entity.Debt) []*entity.Debt {
	o := orm.NewOrm()
	var debts []*entity.Debt
	if param.DebtFrom != nil {
		o.QueryTable("debt").Filter("debt_from_id", param.DebtFrom.Id).All(&debts)
	} else if param.DebtTo != nil {
		o.QueryTable("debt").Filter("debt_to_id", param.DebtTo.Id).All(&debts)
	}else {
		o.QueryTable("debt").All(&debts)
	}

	return debts
}
