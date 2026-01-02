package handler

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base.html",
		"web/templates/pages/home.html",
	))

	data := map[string]string{
		"Title": "GoCore",
	}

	tmpl.Execute(w, data)
}
