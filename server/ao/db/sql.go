package db

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// columnOption 字段选项
type columnOption struct {
	field        string // 字段名
	fieldType    string // 字段类型
	isNull       bool   // 是否 null
	fieldDefault string // 默认值
	key          string // 键，如果是主键，一律自增 auto_increment
	extra        string // 其它
}

// 添加表或字段
func makeTable(tableName string, co ...columnOption) error {
	if len(co) == 0 {
		return errors.New("Empty columnOption")
	}

	sqlPrefix := fmt.Sprintf(`DROP TABLE IF EXISTS %s;
		CREATE TABLE %s(
		`, tableName, tableName)
	var colSQL string
	for i, single := range co {
		colSQL += fmt.Sprintf("%s %s ", single.field, single.fieldType)
		if single.isNull {
			colSQL += "NOT NULL "
		}
		colSQL += single.fieldDefault + " "
		colSQL += single.key
		if strings.Contains(strings.ToLower(single.key), "primary") {
			colSQL += " auto_increment"
		}
		if i != len(co)-1 {
			colSQL += ","
		}
	}
	sqlSuffix := `)engine=innodb DEFAULT charset=utf8mb4;`

	tableSQL := sqlPrefix + colSQL + sqlSuffix
	_, err := _db.Exec(tableSQL)
	if err != nil {
		log.Println(tableSQL)
		return err
	}
	return nil
}
