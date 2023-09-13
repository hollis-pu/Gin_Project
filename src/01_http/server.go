package main

import (
	"fmt"
	"net/http"
)

/**
* Description:
	一起简单的http的响应过程（适用内置包http实现）
* @Author Hollis
* @Create 2023/9/13 14:50
*/

// 处理HTTP请求的处理函数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 向客户端发送响应
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func main() {
	// 创建HTTP路由器
	mux := http.NewServeMux()

	// 注册处理函数
	mux.HandleFunc("/", helloHandler) // 当在浏览器访问http://localhost:9090/时，就会返回helloHandler函数中响应的内容

	// 启动HTTP服务器并监听端口
	port := 9090
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(serverAddr, mux)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
