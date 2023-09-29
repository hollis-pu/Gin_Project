package mysql

import (
	"Gin_Project/src/18_bluebell_project/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 13:36
 */
var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}
func Close() {
	_ = db.Close()
}
