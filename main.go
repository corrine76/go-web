package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc("/", funcs.IndexHandler)
	// http.HandleFunc("/chat-login", funcs.IndexHandler)
	http.HandleFunc("/", funcs.ChatDemoHandler)
	http.ListenAndServe(":8080", nil)
}
