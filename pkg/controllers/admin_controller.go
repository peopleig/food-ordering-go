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
			http.Error(w, "unable to fetch all categories", http.StatusInternalServerError)
			return
		}
		data := types.GetAddDishData{
			Title:      "AddDish",
			Categories: categories,
			Error:      showToast,
			Message:    message,
		}
		utils.RenderTemplate(w, "add_dish", data)
	case http.MethodPost:
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".webp" {
			http.Redirect(w, r, "/admin/dish?error=img", http.StatusSeeOther)
			return
		}
		milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
		gen_name := milli + header.Filename
		url := "static/images/" + gen_name
		dest, err := os.Create("./web/static/images/" + gen_name)
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
		var newDish types.NewDish
		if err := decoder.Decode(&newDish, r.PostForm); err != nil {
			http.Redirect(w, r, "/admin/dish?error=server", http.StatusSeeOther)
			return
		}
		dishExists := utils.DishExists(newDish.DishName)
		if dishExists {
			http.Redirect(w, r, "/admin/dish?error=dish", http.StatusSeeOther)
			return
		}
		var isVeg bool
		switch newDish.IsVeg {
		case "1":
			isVeg = true
		case "0":
			isVeg = false
		default:
			fmt.Println("Incorrect form data")
			http.Redirect(w, r, "/admin/dish?error=req", http.StatusSeeOther)
			return
		}
		price64, _ := strconv.ParseFloat(newDish.Price, 32)
		price32 := float32(price64)
		spiceLevel, _ := strconv.Atoi(newDish.SpiceLevel)
		err = models.AddDish(newDish, price32, isVeg, url, spiceLevel)
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
