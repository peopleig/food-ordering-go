package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/peopleig/food-ordering-go/pkg/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	isPending, err := models.CheckForPendingBills(user_id)
	if err != nil {
		fmt.Println("DB error on logout check:", err)
		http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
		return
	} else if isPending {
		http.Redirect(w, r, "/bill?error=unpaid", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
