package mysql

import (
	"Gin_Project/src/18_bluebell_project/models"
	"crypto/md5"
	"encoding/hex"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 15:19
 */

// 把每一步数据库操作封装成函数，使logic层根据业务需求调用

const secret = "hollis"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (b bool, err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.PassWord = encryptPassword(user.PassWord)
	// 执行SQL语句入库
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, user.PassWord)
	return
}

// encryptPassword 对密码进行MD5加密
func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}

// CheckUserPassword 验证用户名和密码是否正确
func CheckUserPassword(username string, password string) (b bool, err error) {
	var user models.User
	password = encryptPassword(password) // 对输入的密码进行加密转换
	sqlStr := "select username,password from user where username=?"
	if err := db.Get(&user, sqlStr, username); err != nil { // 将查询结果存入结构体变量中
		return false, err
	}
	return user.PassWord == password, nil
}
