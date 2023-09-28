package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

/*
*
  - Description:
    使用go-redis库连接redis数据库
  - @Author Hollis
  - @Create 2023/9/28 10:32
*/
var rdb *redis.Client
var rdbCluster *redis.ClusterClient

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	return
}

// 哨兵模式
func initSentinel() {
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name",
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
}

// 连接到集群
func initCluster() {
	rdbCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},

		// 若要根据延迟或随机路由命令，请启用以下命令之一
		// RouteByLatency: true,
		// RouteRandomly: true,
	})
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("redis connect failed, err:%v\n", err)
		return
	}
	fmt.Println("redis connect successful!")
}
