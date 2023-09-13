package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
* Description:
	template的基本使用：1.先写一个模板文件；2.使用template.ParseFiles来解析模板；3.使用Execute来渲染模板（指定传入模板中的内容）。
	1. 模板文件通常定义为.tmpl和.tpl为后缀（也可以使用其他的后缀），必须使用UTF8编码。
	2. 模板文件中使用{{和}}包裹和标识需要传入的数据。
	3. 传给模板这样的数据就可以通过点号（.）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
	4. 除{{和}}包裹的内容外，其他内容均不做修改原样输出。
* @Author Hollis
* @Create 2023/9/13 17:23
*/

// 遇事不决，写注释
func sayHello01(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./01.hello.tmpl") // 请注意文件路径。编译的.exe文件应该和本源代码文件在同一个路径下才能正确运行。（请勿刻舟求剑）
	// 在终端中输入命令：>  cd .\src\03_template\
	// >  go build .\01.basic_usage.go
	// >  .\01.basic_usage.exe
	if err != nil {
		fmt.Println("Parse template failed, err=", err)
		return
	}
	// 渲染模板
	err = t.Execute(w, "golang")
	if err != nil {
		fmt.Println("Render template failed, err=", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello01)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http server start failed, err=", err)
		return
	}
}
