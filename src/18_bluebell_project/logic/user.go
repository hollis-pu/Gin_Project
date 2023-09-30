package logic

import (
	"Gin_Project/src/18_bluebell_project/dao/mysql"
	"Gin_Project/src/18_bluebell_project/models"
	"Gin_Project/src/18_bluebell_project/pkg/jwt"
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

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		UserName: p.UserName,
		PassWord: p.PassWord,
	}
	// 1.查询用户是否存在
	exist, err := mysql.CheckUserExist(p.UserName)
	if !exist {
		return "", ErrorUserNotExist
	}
	if err != nil {
		return "", err
	}
	// 2.判断密码是否正确
	// 传递的是指针（函数内修改的值能够反应到函数外面），就可以拿到user.UserID，用于生成JWT
	pass, err := mysql.CheckUserPassword(user)
	if !pass {
		return "", ErrorInvalidPassword
	}
	// 3.返回结果
	if err != nil {
		return "", err
	}
	// 4.生成JWT
	return jwt.GenToken(user.UserID, user.UserName)
}
