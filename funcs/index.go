package funcs

import (
	"fmt"
	"html/template"
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
	pageData := BlogPage{
		Home:    "My Blog",
		About:   "about",
		Blog:    "blog",
		Contact: "contact",
	}

	// Parse the blog homepage template
	tmpl, err := template.ParseFiles("index2.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Execute the blog homepage template with the BlogPage data
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
