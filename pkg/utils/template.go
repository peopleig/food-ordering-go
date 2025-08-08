package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"web/templates/layout.html",
		"web/templates/"+tmplName+".html",
	)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error parsing templates", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
