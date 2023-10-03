package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* Description:
* @Author Hollis
* @Create 2023/10/3 11:27
 */

type ParamLogin struct {
	UserName string `json:"username,omitempty" binding:"required" example:"test"`
	Password string `json:"password,omitempty" binding:"required" example:"123456"`
}

// LoginHandler 用户登录
//
//	@Summary		用户登录
//	@Description	用户登录
//	@Tags			User
//	@Accept			application/json
//	@Produce		application/json
//	@Param Login body ParamLogin false "用户名和密码"
//	@Success		200			{object}	_ResponseData
//	@Router			/login [post]
func LoginHandler(c *gin.Context) {
	p := new(ParamLogin)

	err := c.ShouldBindJSON(p)
	if err != nil {
		fmt.Printf("ShouldBindQuery err:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	if p.UserName == "test" && p.Password == "123456" {
		c.JSON(http.StatusOK, _ResponseData{
			Code: 200,
			Msg:  "login successful!",
		})
	} else {
		c.JSON(http.StatusOK, _ResponseData{
			Code: 401,
			Msg:  "wrong user name or password!",
		})
	}
}
