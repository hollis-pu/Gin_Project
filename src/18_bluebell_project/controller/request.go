package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/30 10:59
 */

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录！")

// getCurrentUser 获取当前登录用户id（只能在controller层被调用）
func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
