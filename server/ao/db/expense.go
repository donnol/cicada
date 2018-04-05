package db

// Expense 支出
type Expense struct {
	ID        int
	UserID    int `db:"user_id"`
	Pay       float64
	Thing     string
	CreatedAt string `db:"created_at"`
	CreatedOn string `db:"created_on"`
}

// ExpenseList 支出列表
func ExpenseList() (expenses []Expense, err error) {
	err = _db.Select(&expenses, `SELECT * FROM t_expense
		`,
	)
	if err != nil {
		return
	}
	return
}
