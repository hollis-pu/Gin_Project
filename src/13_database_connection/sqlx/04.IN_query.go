package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

/**
* Description:
	使用sqlx.in进行IN查询
* @Author Hollis
* @Create 2023/9/28 9:59
*/

// QueryByIDs 根据给定ID查询，返回的结果会按照id进行升序排列
func QueryByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	fmt.Println("query:", query)
	fmt.Println("args:", args)
	err = db.Select(&users, query, args...)
	return
}

// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id)) // 将ids存入切片中
	}
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ",")) // 通过FIND_IN_SET指定结果集的排列顺序
	if err != nil {
		return
	}

	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func main() {
	ids := []int{3, 2, 10}
	queryData, err := QueryByIDs(ids)
	if err != nil {
		return
	}
	fmt.Println(queryData)

	queryData, err = QueryAndOrderByIDs(ids)
	if err != nil {
		return
	}
	fmt.Println(queryData)
}
