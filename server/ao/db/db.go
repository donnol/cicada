package db

import (
	"log"

	// 数据库引擎
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 数据库实例
var db *sqlx.DB = func() *sqlx.DB {
	db, err := sqlx.Open("mysql", "jd:jd@/cicada")
	if err != nil {
		log.Fatal(err)
	}
	return db
}()
