package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peopleig/food-ordering-go/pkg/models"
)

func AllowChefAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user_id := r.Context().Value("user_id").(int)

		role := r.Context().Value("role").(string)
		fmt.Println(role)
		if role != "chef" {
			http.Error(w, "This is the chef's zone - Not for you!", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func AllowAdminAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user_id := r.Context().Value("user_id").(int)

		role := r.Context().Value("role").(string)
		fmt.Println(role)
		if role != "admin" {
			http.Error(w, "Admin's Playground. Not Yours. Bye bye!", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}

func AllowAdminandIdAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := r.Context().Value("user_id").(int)
		role := r.Context().Value("role").(string)
		vars := mux.Vars(r)
		orderId, _ := strconv.Atoi(vars["order_id"])
		userId, err := models.CheckForUser(orderId)
		if err != nil {
			http.Error(w, "error confirming user", http.StatusInternalServerError)
		}
		if role != "admin" && user_id != userId {
			http.Error(w, "Admin's Playground. Not Yours. Bye bye!", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
