package models

/**
* Description:
* @Author Hollis
* @Create 2023/9/29 17:17
 */

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	PassWord string `db:"password"`
}
