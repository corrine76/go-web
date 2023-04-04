package funcs

import (
	"html/template"
	"log"
	"net/http"
)

// BlogPage represents the data to be used in the blog homepage template
type BlogPage struct {
	Title   string
	Home    string
	About   string
	Blog    string
	Contact string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 页面数据
	pageData := BlogPage{
		Home:    "My Blog",
		About:   "about",
		Blog:    "blog",
		Contact: "contact",
	}

	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		log.Println("Error parsing template", err)
		return
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Println("Error executing template", err)
		return
	}
}
