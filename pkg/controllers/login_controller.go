package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/peopleig/food-ordering-go/pkg/middleware"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := map[string]string{
			"Title": "Login",
		}
		utils.RenderTemplate(w, "login", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		login_type := r.FormValue("login_type")
		identifier := r.FormValue("identifier")
		password := r.FormValue("password")
		isValid, message := utils.CheckLoginTypeValidity(login_type, identifier)
		if !isValid {
			http.Error(w, message, http.StatusBadRequest)
		}
		if password == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing required fields"))
			return
			// http.Error(w, "Cannot have an empty password", http.StatusBadRequest)
		}
		user, errr := models.GetUserPwdatLogin(login_type, identifier)
		if errr != nil {
			if errr == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
		}

		if !middleware.CheckPasswordHash(password, user.Hash_pwd) {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}
		if !user.Approved {
			http.Error(w, "Cheeky cheeky. But you're still not approved", http.StatusUnauthorized)
		}
		user_id, _ := strconv.Atoi(user.User_id)

		token, err := utils.GenerateJWT(user_id, user.Role)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		})
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
		// switch user.Role {
		// case "admin":
		// 	http.Redirect(w, r, "/admin", http.StatusSeeOther)
		// case "chef":
		// 	http.Redirect(w, r, "/chef", http.StatusSeeOther)
		// case "customer":
		// 	http.Redirect(w, r, "/menu", http.StatusSeeOther)
		// }

	}
}
