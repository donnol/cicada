package ao

import (
	"cicada/server/ao/db"
)

// ExpenseList 支出列表
func ExpenseList() (expenses []db.Expense, err error) {
	return db.ExpenseList()
}
