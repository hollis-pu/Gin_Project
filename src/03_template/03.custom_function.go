package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
* Description:
	自定义模板和自定义模板的函数
* @Author Hollis
* @Create 2023/9/13 19:47
*/

func customTemplate(w http.ResponseWriter, r *http.Request) {
	myFunc := func(name string) (string, error) {
		return name + " is beautiful！", nil
	}
	tmpl := template.New("03.use_custom_func.html") // 创建一个名字为03.use_custom_func的模板对象，名字一定要与模板的名字能对应上。
	// 告诉模板，我们定义了一个自定义函数cusFunc，可以在模板文件中使用该函数
	// 这个动作要在解析模板之前
	tmpl.Funcs(template.FuncMap{
		"cusFunc": myFunc,
	})

	// 解析模板
	_, err := tmpl.ParseFiles("./03.use_custom_func.html") // 注意：这里要使用自定义的模板对象来解析
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
	http.HandleFunc("/", customTemplate)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server start failed, err=", err)
		return
	}
}
