package http

import (
	"ingat_utang/entity"
)

const StatusOk  = 200
const StatusForbidden  = 401
const StatusSaveSuccess = 1010
const StatusSaveFailed = 1011
const StatusLoginError = 1020
var statuses = []string{200:"OK", 401:"Permission denied", 1010:"Save Success", 1011:"Save failed", 1020:"Login gagal, check username and password"}

func Result(v interface{}, status int) entity.Base {
	return entity.Base{status,statuses[status], v}
}