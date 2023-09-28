package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
*
  - Description:
    增删改查操作
  - @Author Hollis
  - @Create 2023/9/27 20:38
*/

// 1.单行查询：使用db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。
func queryRowData() {
	sqlStr := "select id,name,age from user where id=?"
	var user User
	row := db.QueryRow(sqlStr, 1) // 得到一行的记录
	// 要确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := row.Scan(&user.Id, &user.Name, &user.Age) // 将查询结果的列和结构体字段进行映射赋值（根据查询出来的列的顺序进行依次赋值）
	// 也可以直接链式调用: err := db.QueryRow(sqlStr, 1).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("user:", user) // user: {1 Tom 25}
}

// 2.多行查询：使用db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
func queryMultiRowData() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

// 3.插入数据：插入、更新和删除操作都使用Exec方法。
func insertRowData() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38) // 第一个返回值类型为Result，保存了执行结果的信息（自增ID、影响行数）
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 4.更新数据
func updateRowData() {
	sqlStr := "update user set age=? where name = ?"
	ret, err := db.Exec(sqlStr, 30, "Tom")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 5.删除数据
func deleteRowData() {
	sqlStr := "delete from user where name = ?"
	ret, err := db.Exec(sqlStr, "王五")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	queryRowData()
	queryMultiRowData()
	insertRowData()
	updateRowData()
	deleteRowData()
}
