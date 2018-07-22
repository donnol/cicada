package model

import (
	"log"
)

// Expense 支出
type Expense struct {
	ID        *int
	UserID    *int    `db:"user_id"`
	UserName  *string `db:"user_name"`
	Pay       *float64
	Thing     *string
	CreatedAt *string `db:"created_at"`
	CreatedOn *string `db:"created_on"`
	UpdatedAt *string `db:"updated_at"`
	Total     *int
	IDs       *[]int `db:"-"`
}

// ExpenseParam 参数
type ExpenseParam struct {
	CommonParam
	Expense
}

// ExpenseList 支出列表
func ExpenseList(ep ExpenseParam) (expenses []Expense, err error) {
	err = _db.Select(&expenses, `SELECT *,
		COUNT(*) OVER() AS total
		FROM t_expense
		WHERE true

		AND CASE WHEN $4::bigint IS NOT NULL THEN
			id = $4
		ELSE true END

		ORDER BY 
			CASE $3
				WHEN 'pay' THEN pay
				ELSE id 
			END
		DESC

		LIMIT $1

		OFFSET $2
		`,
		ep.Size,
		ep.Offset,
		ep.Order,
		ep.ID,
	)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	return
}
