package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 22:03
 */

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data"`
}

// ResponseError 响应“错误”的信息
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithMsg 指定msg的响应信息
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 响应“成功”的信息
func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
