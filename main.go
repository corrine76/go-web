package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static-chat/", http.StripPrefix("/static-chat/", fs))

	http.HandleFunc("/", funcs.ChatDemoHandler)
	http.ListenAndServe(":8080", nil)
}
