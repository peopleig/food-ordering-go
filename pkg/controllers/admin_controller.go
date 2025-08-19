package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/peopleig/food-ordering-go/pkg/cache"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	showMsg := r.URL.Query().Get("show")
	show := false
	message := ""
	if showMsg == "dishsuccess" {
		show = true
		message = "Dish Added Successfully!"
	}
	var items []types.Ordered
	err := models.GetAllOrderedItems(&items)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while accessing DB", http.StatusBadRequest)
		return
	}
	var orders []types.Order
	err = models.GetAllOrders(&orders)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while accessing DB", http.StatusBadRequest)
		return
	}
	var uausers []types.UnApprovedUser
	err = models.GetUnapprovedUsers(&uausers)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while accessing DB", http.StatusBadRequest)
		return
	}
	data := types.AdminData{
		Title:   "Admin",
		Items:   items,
		Orders:  orders,
		Uausers: uausers,
		Show:    show,
		Message: message,
		Role:    "admin",
	}
	utils.RenderTemplate(w, "admin", data)

}

func AdminApproveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	user_Id, _ := strconv.Atoi(userID)
	switch r.Method {
	case http.MethodPatch:
		err := models.ApproveUser(user_Id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error while accessing DB", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		err := models.RemoveUserRequest(user_Id)
		if err != nil {
			http.Error(w, "unable to decline access request", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

var decoder = schema.NewDecoder()

func AdminDishHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// In this I check for what the errorMsg is - which is probably something I'll do everywhere
		errorMsg := r.URL.Query().Get("error")
		showToast, message := utils.AddDishErrors(errorMsg)
		var categories []types.Categories
		err := models.GetAllCategories(&categories)
		if err != nil {
			http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
			return
		}
		data := types.GetAddDishData{
			Title:      "AddDish",
			Categories: categories,
			Error:      showToast,
			Message:    message,
			Role:       "admin",
		}
		utils.RenderTemplate(w, "add_dish", data)
	case http.MethodPost:
		r.ParseMultipartForm(10 << 20)
		//All the other fields I will check first
		var newDish types.NewDish
		if err := decoder.Decode(&newDish, r.PostForm); err != nil {
			http.Redirect(w, r, "/admin/dish?error=server", http.StatusSeeOther)
			return
		}
		if utf8.RuneCountInString(newDish.Description) > 100 {
			http.Redirect(w, r, "/admin/dish?error=len", http.StatusSeeOther)
			return
		}
		if newDish.Price > 9999 {
			http.Error(w, "price out of bounds!", http.StatusBadRequest)
			return
		}
		dishExists := utils.DishExists(newDish.DishName)
		if dishExists {
			http.Redirect(w, r, "/admin/dish?error=dish", http.StatusSeeOther)
			return
		}
		if newDish.SpiceLevel > 5 || newDish.SpiceLevel < -5 {
			http.Redirect(w, r, "/admin/dish?error=spice", http.StatusSeeOther)
			return
		}
		//Now the file
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
			return
		}
		defer file.Close()
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".webp" {
			http.Redirect(w, r, "/admin/dish?error=img", http.StatusSeeOther)
			return
		}
		milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
		url := "static/images/" + milli + ext
		dest, err := os.Create("./web/static/images/" + milli + ext)
		if err != nil {
			fmt.Println("Unable to create file: ", err)
			http.Redirect(w, r, "/admin/dish?error=server", http.StatusSeeOther)
			return
		}
		defer dest.Close()
		_, err = io.Copy(dest, file)
		if err != nil {
			fmt.Println("Unable to save file: ", err)
			http.Redirect(w, r, "/admin/dish?error=server", http.StatusSeeOther)
			return
		}
		err = models.AddDish(newDish, url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cache.RefreshMenuCache()
		_ = cache.LoadMenu()
		http.Redirect(w, r, "/admin?show=dishsuccess", http.StatusSeeOther)
	}
}

func AdminCategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var categories []types.Categories
		err := models.GetAllCategories(&categories)
		if err != nil {
			http.Error(w, "unable to fetch all categories", http.StatusInternalServerError)
			return
		}
		data := types.GetAddDishData{
			Title:      "Add Dish",
			Categories: categories,
		}
		utils.RenderTemplate(w, "add_dish", data)
	}
}

func AdminPaymentApproveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, _ := strconv.Atoi(vars["order_id"])
	err := models.ValidateBillPayment(orderId)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
