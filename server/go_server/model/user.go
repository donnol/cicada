package model

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// User 用户
type User struct {
	ID       int
	Phone    string
	Password string `json:"-"`
	Name     string // 注册时
}

// PhoneRegisterCode 获取注册码
func PhoneRegisterCode(phone string) (code string, err error) {
	tx := _db.MustBegin()
	err = wrapTx(tx, func() error {
		// 是否已注册
		var exist bool
		err := tx.Get(&exist, `SELECT EXISTS(SELECT FROM t_user
			WHERE phone = $1)
			`,
			phone,
		)
		if err != nil {
			return err
		}
		if exist {
			err = errors.New("phone have used")
			return err
		}

		// 生产随机数
		for i := 0; i < 6; i++ {
			n := rand.Intn(10)
			code += strconv.Itoa(n)
		}

		// 保存到数据库
		_, err = tx.Exec(`INSERT INTO t_phone_code
			(phone, code)
			VALUES($1, $2)
			`,
			phone,
			code,
		)
		if err != nil {
			log.Printf("%v\n", err)
			return err
		}
		return nil
	})
	if err != nil {
		return
	}

	return
}

// PhoneRegister 注册 TODO
func PhoneRegister(phone, code, password string) (u User, err error) {
	// 校验 phone, code

	// 生成随机名称

	// 保存用户

	return
}

// NameRegister 账号注册
func NameRegister(name, password string) (u User, err error) {
	password, err = hashPassword(password)
	if err != nil {
		return
	}

	var id int
	err = _db.Get(&id, `INSERT INTO t_user (name, password)
		VALUES($1, $2)
		RETURNING id
		`,
		name,
		password,
	)
	if err != nil {
		return
	}
	u.ID = id
	u.Name = name
	return
}

// 密码加密存储
func hashPassword(password string) (string, error) {
	cost := 16
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 密码校验
func verifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// NameLogin 账号登陆
func NameLogin(name, password string) (u User, err error) {
	err = _db.Get(&u, `SELECT id, name, password FROM t_user
		wHERE name = $1
		`,
		name,
	)
	if err == sql.ErrNoRows {
		err = errors.New("Don't exist user")
		return
	}
	if err != nil {
		return
	}
	if err = verifyPassword(u.Password, password); err != nil {
		return
	}
	return
}

// Logout 退出登陆
func Logout() (err error) {
	return
}

// ModifyUser 修改用户信息
func ModifyUser(id int, u User) (err error) {
	return
}
