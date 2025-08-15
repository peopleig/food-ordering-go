package controllers

import (
	"html/template"
	"net/http"

	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func ErrorController(w http.ResponseWriter, r *http.Request) {
	var message template.HTML
	var statusCode string
	err := r.URL.Query().Get("error")
	if err == "signup" {
		statusCode = "401: Unauthorized"
		message = "You have signed up asking for a position of authority<br>You will be given access after the admin grants it to you<br>Till then, you need to wait"
	} else if err == "login" {
		statusCode = "401: Unauthorized"
		message = "<b>Cheeky cheeky</b><br>But you're still not approved<br>Wait while the admin does it"
	}
	data := types.ErrorPageData{
		Title:   "Error",
		Status:  statusCode,
		Message: message,
	}
	utils.RenderTemplate(w, "error", data)
}
