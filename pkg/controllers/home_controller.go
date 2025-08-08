package controllers

import (
	"net/http"

	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Home",
	}
	utils.RenderTemplate(w, "home", data)
}
