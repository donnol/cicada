package ao

import (
	"errors"
	"math/rand"
	"strconv"

	"cicada/server/ao/db"
)

// RegisterCode 获取注册码
func RegisterCode(phone string) (code string, err error) {
	// 是否已注册
	var exist bool
	if exist, err = db.ExistPhone(phone); err != nil {
		return
	}
	if exist {
		err = errors.New("phone have used")
		return
	}

	// 生产随机数
	for i := 0; i < 6; i++ {
		n := rand.Intn(10)
		code += strconv.Itoa(n)
	}

	// 保存到数据库
	err = db.AddPhoneCode(phone, code)
	if err != nil {
		return
	}

	return
}

// Register 注册 TODO
func Register(phone, code, password string) (u db.User, err error) {
	// 校验 phone, code

	// 生成随机名称

	// 保存用户

	return
}

// LoginCode 获取动态码
func LoginCode(phone string) (code string, err error) {
	return
}

// LoginWithCode 动态码登陆
func LoginWithCode(phone, code string) (u db.User, err error) {
	return
}

// LoginWithPassword 密码登陆
func LoginWithPassword(phone, password string) (u db.User, err error) {
	return
}

// Logout 退出登陆
func Logout() (err error) {
	return
}

// ModifyUser 修改用户信息
func ModifyUser(id int, u db.User) (err error) {
	return
}
