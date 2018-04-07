package ao

import (
	"cicada/server/ao/db"
)

// ExpenseParam 支出参数
type ExpenseParam = db.ExpenseParam

// ExpenseList 支出列表
func ExpenseList(ep ExpenseParam) (expenses []db.Expense, err error) {
	limit, offset, order := 10, 0, "pay"
	ep.CommonParam.Limit = limit
	ep.CommonParam.Offset = offset
	ep.CommonParam.Order = order
	// var id = 3
	// ep.ID = &id
	return db.ExpenseList(ep)
}
