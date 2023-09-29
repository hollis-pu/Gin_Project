package models

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 15:48
 */

// ParamSignUp 定义注册的参数结构体
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"re_password" binding:"required,eqfield=PassWord"`
}

// ParamLogin 定义登录的参数结构体
type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}
