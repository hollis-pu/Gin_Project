package main

import (
	"log"
	"net/http"
	"os"
)

/*
*
  - Description:
    默认logger的使用
	默认logger的缺点：
		缺少一个ERROR日志级别，这个级别可以在不抛出panic或退出程序的情况下记录错误。
		缺乏日志格式化的能力。
		不提供日志切割的能力。
  - @Author Hollis
  - @Create 2023/9/28 13:18
*/

// InitLogger01 设置log文件
func InitLogger01() {
	logFileLocation, _ := os.OpenFile("./src/14_zap/01.default_log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

// http请求
func simpleHttpGet01(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}

func main() {
	InitLogger01()
	simpleHttpGet01("www.google.com")
	simpleHttpGet01("https://www.google.com")
}
