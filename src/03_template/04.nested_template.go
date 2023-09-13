package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
* Description:
	嵌套模板
* @Author Hollis
* @Create 2023/9/13 21:12
*/

func nestedTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析嵌套的模板
	tmpl, err := template.ParseFiles("./04.t.tmpl", "04.ul.tmpl") // 外层的模板写在前面
	if err != nil {
		fmt.Println("Parse template failed, err=", err)
		return
	}

	// 渲染模板
	err = tmpl.Execute(w, "Jane")
	if err != nil {
		fmt.Println("Render template failed, err=", err)
		return
	}
}

func main() {
	http.HandleFunc("/", nestedTemplate)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server start failed, err=", err)
		return
	}
}
