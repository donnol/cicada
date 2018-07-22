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
	Size   int
	Offset int
	Order  string
}

// 如果 f 参数的类型是 func(tx *sqlx.Tx)，在执行时，需要将 tx 传入，f 函数里会不会有改变 tx 参数的可能呢？
func wrapTx(tx *sqlx.Tx, f func() error) error {
	err := f()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
