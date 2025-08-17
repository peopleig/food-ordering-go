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
	switch err {
	case "signup":
		statusCode = "401: Unauthorized"
		message = "You have signed up asking for a position of authority<br>You will be given access after the admin grants it to you<br>Till then, you need to wait"
	case "login":
		statusCode = "401: Unauthorized"
		message = "<b>Cheeky cheeky</b><br>But you're still not approved<br>Wait while the admin does it"
	case "internal":
		statusCode = "500: Internal Server Error"
		message = "<b>Extremely Sorry!<b><br>Internal Server Error<br>Please go back and try again."
	case "chef":
		statusCode = "403: Forbidden"
		message = "This is the <b>Chef's</b> space<br>Not yours!<br>If you're interested though, feel free to apply for a job"
	case "admin":
		statusCode = "403: Forbidden"
		message = "This is the <b>Admin's</b> playground<br>Not yours!<br>Bye-bye now!"
	case "bill":
		statusCode = "403: Forbidden"
		message = "Not <b>your bill</b> now, is it?<br>Don't mess around here<br>Tata!"
	}
	data := types.ErrorPageData{
		Title:   "Error",
		Status:  statusCode,
		Message: message,
		Role:    "lost",
	}
	utils.RenderTemplate(w, "error", data)
}

func Custom404Handler(w http.ResponseWriter, r *http.Request) {
	var message template.HTML
	message = "<b>Lost?</b><br>How'd you land here?<br>Even we don't go here..."
	data := types.ErrorPageData{
		Title:   "Error",
		Status:  "404: Not Found",
		Message: message,
		Role:    "lost",
	}
	utils.RenderTemplate(w, "error", data)
}
