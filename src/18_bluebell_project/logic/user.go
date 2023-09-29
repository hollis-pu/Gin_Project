package logic

import (
	"Gin_Project/src/18_bluebell_project/dao/mysql"
	"Gin_Project/src/18_bluebell_project/models"
	"Gin_Project/src/18_bluebell_project/pkg/snowflake"
	"errors"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 15:09
 */

// 存放业务逻辑的代码

var (
	ErrorUserExist       = errors.New("用户已存在！")
	ErrorUserNotExist    = errors.New("用户不存在！")
	ErrorInvalidPassword = errors.New("密码错误！")
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	exist, err := mysql.CheckUserExist(p.UserName)
	if exist {
		return ErrorUserExist
	}
	if err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 3.构造一个User实例
	user := &models.User{
		UserID:   userID,
		UserName: p.UserName,
		PassWord: p.PassWord,
	}
	// 4.保存到数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (err error) {
	// 1.查询用户是否存在
	exist, err := mysql.CheckUserExist(p.UserName)
	if !exist {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 2.判断密码是否正确
	pass, err := mysql.CheckUserPassword(p.UserName, p.PassWord)
	if !pass {
		return ErrorInvalidPassword
	}
	// 3.返回结果
	return err
}
