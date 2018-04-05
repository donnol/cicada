package db

import (
	"database/sql"
)

// User 用户
type User struct {
	ID       int
	Phone    string
	Password string
	Name     string // 注册时
}

// ExistPhone 是否存在手机号
func ExistPhone(phone string) (exist bool, err error) {
	var userID int
	err = _db.Get(&userID, `SELECT id FROM t_user
		WHERE phone = ?
		`,
		phone,
	)
	if err == sql.ErrNoRows {
		exist = false
		err = nil
		return
	}
	if err != nil {
		return
	}
	exist = true
	return
}

// AddPhoneCode 添加手机号和验证码
func AddPhoneCode(phone, code string) (err error) {
	_, err = _db.Exec(`INSERT INTO t_phone_code
		(phone, code)
		VALUES(?, ?)
		`,
		phone,
		code,
	)
	if err != nil {
		return
	}

	return
}

// ExistPhoneCode 是否存在手机号和验证码
func ExistPhoneCode(phone, code string) (exist bool, err error) {
	var id int
	err = _db.Get(&id, `SELECT id FROM t_phone_code
		WHERE phone = ? AND code = ? AND used = false
		`,
		phone,
		code,
	)
	if err == sql.ErrNoRows {
		exist = false
		err = nil
		return
	}
	if err != nil {
		return
	}

	exist = true
	return
}
