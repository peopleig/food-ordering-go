package middleware

import (
	"fmt"
	"net/http"
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
