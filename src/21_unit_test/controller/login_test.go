package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
* Description:
	单元测试，测试路由：/login
* @Author Hollis
* @Create 2023/10/3 16:54
*/

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/login"
	r.POST(url, LoginHandler)

	body := `{
		"username": "test",
    	"password": "123456"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)                             // w.Code 响应码
	assert.Contains(t, w.Body.String(), "login successful!") // w.Body.String() 响应体
}
