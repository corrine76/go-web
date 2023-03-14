package funcs

import (
	"html/template"
	"log"
	"net/http"
)

type FormData struct {
	Name    string
	Email   string
	Message string
}

// ChatLogin 聊天登陆
func ChatLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderForm(w)
	} else if r.Method == "POST" {
		processForm(w, r)
	}
}

func renderForm(w http.ResponseWriter) {
	t, err := template.ParseFiles("../chatlogin.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	formData := FormData{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	// 重定向到新的URL
	http.Redirect(w, r, "/thankyou", http.StatusSeeOther)

	// 在这里可以将表单数据保存到数据库或发送电子邮件等操作

	// 输出表单数据到日志
	log.Printf("Name: %s, Email: %s, Message: %s", formData.Name, formData.Email, formData.Message)
}

// Chat 开始聊天
func Chat(w http.ResponseWriter, r *http.Request) {
}
