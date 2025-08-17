package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/schema"
	"github.com/peopleig/food-ordering-go/pkg/middleware"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
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
		var decoder = schema.NewDecoder()
		var loginData types.LoginData
		if err := decoder.Decode(&loginData, r.PostForm); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		isValid, message := utils.CheckLoginTypeValidity(loginData.LoginType, loginData.Identifier)
		if !isValid {
			http.Error(w, message, http.StatusBadRequest)
		}
		if loginData.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]string{
				"Title":   "Login",
				"Message": "Enter the Password!",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "login", data)
			return
		}
		user, errr := models.GetUserPwdatLogin(loginData.LoginType, loginData.Identifier)
		if errr != nil {
			if errr == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				data := map[string]string{
					"Title":   "Login",
					"Message": "Invalid Credentials",
					"Error":   "True",
				}
				utils.RenderTemplate(w, "login", data)
				return
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
		}

		if !middleware.CheckPasswordHash(loginData.Password, user.Hash_pwd) {
			w.WriteHeader(http.StatusUnauthorized)
			data := map[string]string{
				"Title":   "Login",
				"Message": "Incorrect Password",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "login", data)
			return
		}
		if !user.Approved {
			http.Redirect(w, r, "/error?error=login", http.StatusSeeOther)
			return
		}
		user_id, _ := strconv.Atoi(user.User_id)

		token, err := utils.GenerateJWT(user_id, user.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data := map[string]string{
				"Title":   "Login",
				"Message": "Sorry, Server Error. Please try again",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "login", data)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
		})
		switch user.Role {
		case "admin":
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		case "chef":
			http.Redirect(w, r, "/chef", http.StatusSeeOther)
		case "customer":
			http.Redirect(w, r, "/menu", http.StatusSeeOther)
		}
	}
}
