package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
*
  - Description:
    模板语法的进一步介绍
  - @Author Hollis
  - @Create 2023/9/13 17:52
*/

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello02(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./02.user.html")
	if err != nil {
		fmt.Println("Parse template failed, err=", err)
		return
	}

	user01 := UserInfo{"Tom", "male", 26}
	// 渲染模板
	//err = t.Execute(w, user01) // 传入结构体变量
	//if err != nil {
	//	fmt.Println("Render template failed, err=", err)
	//	return
	//}

	// 通过渲染模板传递多个变量
	book01 := map[string]any{
		"name":   "Go语言从入门到精通", // map的key首字母是可以小写的
		"author": "零壹快学",
		"price":  79,
	}
	hobbyList := []string{"sing", "read", "study"}
	err = t.Execute(w, map[string]any{ // 这里通过map传入多个变量
		"user":  user01,
		"book":  book01,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Println("Render template failed, err=", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello02)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server start failed, err=", err)
		return
	}
}
