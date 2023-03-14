package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	// http.HandleFunc("/", funcs.IndexHandler)
	// http.HandleFunc("/chat-login", funcs.IndexHandler)
	http.HandleFunc("/", funcs.ChatDemoHandler)
	http.ListenAndServe(":8080", nil)
}
