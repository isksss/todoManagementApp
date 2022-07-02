package router

import (
	"fmt"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, file ...string) {
	var files []string
	for _, file := range file {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)
}
