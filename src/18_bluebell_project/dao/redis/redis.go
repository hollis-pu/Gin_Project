package redis

import (
	"Gin_Project/src/18_bluebell_project/settings"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 13:44
 */
var rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.Db,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	return
}
func Close() {
	_ = rdb.Close()
}
