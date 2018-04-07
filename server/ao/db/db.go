package db

import (
	"log"

	// 工具
	_ "cicada/server/util"

	// 数据库引擎
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// _db 数据库实例
var _db *sqlx.DB = func() *sqlx.DB {
	db, err := sqlx.Open("mysql", "jd:jd@/cicada")
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
