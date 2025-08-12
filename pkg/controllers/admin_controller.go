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
	}
	utils.RenderTemplate(w, "admin", data)

}

func AdminApproveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]
	user_Id, _ := strconv.Atoi(userID)
	err := models.ApproveUser(user_Id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while accessing DB", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

var decoder = schema.NewDecoder()

func AdminDishHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".webp" {
			http.Error(w, "Invalid image file. Only PNG/JPG/WEBP allowed", http.StatusBadRequest)
			return
		}

		milli := strconv.FormatInt(time.Now().UnixMilli(), 10)
		gen_name := milli + header.Filename
		url := "static/images/" + gen_name
		dest, err := os.Create("./web/static/images/" + gen_name)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to create file", http.StatusInternalServerError)
			return
		}
		defer dest.Close()
		_, err = io.Copy(dest, file)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
		var newDish types.NewDish
		if err := decoder.Decode(&newDish, r.PostForm); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var isVeg bool
		switch newDish.IsVeg {
		case "1":
			isVeg = true
		case "0":
			isVeg = false
		default:
			http.Error(w, "Incorrect form data", http.StatusBadRequest)
			return
		}
		dishExists := utils.DishExists(newDish.DishName)
		if dishExists {
			http.Error(w, "dish with such a name already exists", http.StatusBadRequest)
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
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
