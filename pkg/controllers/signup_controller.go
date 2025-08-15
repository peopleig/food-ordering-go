package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/peopleig/food-ordering-go/pkg/middleware"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data := map[string]string{
			"Title": "Signup",
		}
		utils.RenderTemplate(w, "signup", data)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		var decoder = schema.NewDecoder()
		var new_user types.NewUser
		if err := decoder.Decode(&new_user, r.PostForm); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if new_user.Email == "" && new_user.Mobile == "" {
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]string{
				"Title":   "Signup",
				"Message": "Atleast one of mobile/email has to be entered",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		if new_user.First_name == "" || new_user.Last_name == "" || new_user.Role == "" || new_user.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]string{
				"Title":   "Signup",
				"Message": "Cannot have empty name/role/password fields",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		isValid, message := utils.CheckSignupFormValidity(new_user.Email, new_user.Mobile)
		if !isValid {
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		new_user.Password, err = middleware.HashPassword(new_user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data := map[string]string{
				"Title":   "Signup",
				"Message": "Sorry! Server error. Please Try again!",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		var user_id int64
		isValid, message, user_id, err = models.CreateNewUser(&new_user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data := map[string]string{
				"Title":   "Signup",
				"Message": "Sorry! Server error. Please Try again!",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		if !isValid {
			w.WriteHeader(http.StatusBadRequest)
			data := map[string]string{
				"Title":   "Signup",
				"Message": message,
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		if new_user.Role != "customer" {
			http.Error(w, "You haven't yet been granted approval from the admin for this role. Please wait till then", http.StatusUnauthorized)
			return
		}

		token, err := utils.GenerateJWT(int(user_id), new_user.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			data := map[string]string{
				"Title":   "Signup",
				"Message": "Sorry! Server error. Please Try again!",
				"Error":   "True",
			}
			utils.RenderTemplate(w, "signup", data)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token",
			Value:    token,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
		http.Redirect(w, r, "/menu", http.StatusSeeOther)

	}
}
