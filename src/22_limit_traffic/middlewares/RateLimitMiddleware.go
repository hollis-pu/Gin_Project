package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit" // 令牌桶
	ratelimit1 "go.uber.org/ratelimit"     // 漏桶
	"net/http"
	"time"
)

/**
* Description:
	限流中间件
* @Author Hollis
* @Create 2023/10/3 18:09
*/

// RateLimitMiddleware1 基于漏桶的限流中间件1
func RateLimitMiddleware1() func(ctx *gin.Context) {
	rl := ratelimit1.New(1) // 生成一个限流器，参数rate表示水滴的时间间隔（单位：Nanosecond）
	return func(c *gin.Context) {
		nextAvailableTime := rl.Take()
		// 2023.10.03 无论ratelimit1.New()函数中的参数rate设置多大的值，rl.Take()始终等于time.Now()。所以无法实现限流的功能，也不知道原因。
		fmt.Println("nextAvailableTime.Sub(time.Now()):", nextAvailableTime.Sub(time.Now()))
		if nextAvailableTime.Sub(time.Now()) > 0 { // 取水滴。Take()返回下一次可以访问的时间
			fmt.Println("rl.Take().Sub(time.Now()):", rl.Take().Sub(time.Now()))
			c.String(http.StatusOK, "Your visits are too frequent.")
			c.Abort()
			return
		}
		c.Next()
	}
}

// RateLimitMiddleware2 基于令牌桶的限流中间件2
func RateLimitMiddleware2(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit2.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "Your visits are too frequent.")
			c.Abort()
			return
		}
		c.Next()
	}
}
