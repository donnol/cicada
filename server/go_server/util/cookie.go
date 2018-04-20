package util

import (
	"strconv"
)

// CookieUser 获取用户 ID
func CookieUser(value string) (userID int, err error) {
	userID, err = strconv.Atoi(value)
	if err != nil {
		return
	}
	return
}
