package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
* Description:
	单元测试，测试路由 /ping
* @Author Hollis
* @Create 2023/10/3 17:11
*/

func TestPingHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/ping"
	r.GET(url, PingHandler)

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)              // w.Code 响应码
	assert.Contains(t, w.Body.String(), "ok") // w.Body.String() 响应体
}
