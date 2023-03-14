package main

import (
	"goweb/funcs"
	"net/http"
)

func main() {
	http.HandleFunc("/", funcs.IndexHandler)
	http.HandleFunc("/chat-login", funcs.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
