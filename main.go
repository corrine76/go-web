package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	// 定义文件服务器
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static-assets/", http.StripPrefix("/static-assets/", fs))

	// 定义请求接口
	http.HandleFunc("/", funcs.IndexHandler)

	// 监听端口号
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
