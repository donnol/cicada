package db

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
	namedStmt, err := _db.PrepareNamed(`SELECT *,
		COUNT(*) OVER() AS total
		FROM t_expense
		WHERE true

		AND CASE WHEN :id IS NOT NULL THEN
			id = :id
		ELSE true END

		ORDER BY 
			CASE :order
				WHEN 'pay' THEN pay
				ELSE id 
			END
		DESC

		LIMIT :limit

		OFFSET :offset
		`)
	if err != nil {
		log.Printf("%v, %v\n", err, namedStmt)
		return
	}
	err = namedStmt.Select(&expenses, map[string]interface{}{
		"order":  ep.Order,
		"limit":  ep.Limit,
		"offset": ep.Offset,
		"id":     ep.ID,
		"ids":    []int{1, 2, 3},
		// "user_id":    ep.UserID,
		// "pay":        ep.Pay,
		// "thing":      ep.Thing,
		// "created_at": ep.CreatedAt,
		// "created_on": ep.CreatedOn,
	})
	if err != nil {
		log.Printf("%v, %v", err, namedStmt)
		return
	}
	return
}
