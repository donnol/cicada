package model

import (
	"log"

	// 工具
	_ "cicada/server/go_server/util"

	"github.com/jmoiron/sqlx"
	// 数据库引擎
	_ "github.com/lib/pq"
)

// _db 数据库实例
var _db *sqlx.DB = func() *sqlx.DB {
	var connInfo = "postgres://jd:jd@localhost/cicada?sslmode=disable"
	db, err := sqlx.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}()

// CommonParam 通用参数
type CommonParam struct {
	Limit  int
	Offset int
	Order  string
}

func wrapTx(tx *sqlx.Tx, f func(tx *sqlx.Tx) error) error {
	err := f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
