package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	// // 定义文件服务器
	// fs := http.FileServer(http.Dir("assets"))
	// http.Handle("/static-assets/", http.StripPrefix("/static-assets/", fs))

	// 请求首页
	http.HandleFunc("/", funcs.IndexHandler)
	// 请求详情
	// 请求代理跳转
	// 免责申明
	// 用户反馈
	// 用户注册
	// 隐私政策

	// 监听端口号
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
