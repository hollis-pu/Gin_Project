package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* Description:
	在你代码中处理请求的接口函数（通常位于controller层）按如下方式写上注释
* @Author Hollis
* @Create 2023/10/3 10:51
*/

type _ResponseData struct {
	Code int    `json:"code"`           // 状态码
	Msg  string `json:"msg"`            // 状态信息
	Data any    `json:"data,omitempty"` // 数据
}

// PingHandler 测试网站路由
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, _ResponseData{
		Code: 200,
		Msg:  "ok",
		Data: nil,
	})
}
