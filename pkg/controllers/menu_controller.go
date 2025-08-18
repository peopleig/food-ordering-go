package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/peopleig/food-ordering-go/pkg/cache"
	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	role := r.Context().Value("role").(string)
	switch r.Method {
	case http.MethodGet:
		data := types.MenuData{
			Title:      "Menu",
			Items:      config.MenuCache,
			Categories: config.CategoryCache,
			Role:       role,
		}
		utils.RenderTemplate(w, "menu", data)

	case http.MethodPost:
		var order types.OrderRequest
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Incorrect Cart data sent", http.StatusBadRequest)
			return
		}
		if utf8.RuneCountInString(order.Special_instructions) > 50 {
			http.Error(w, "Too long. Make a shorter instruction", http.StatusBadRequest)
			return
		}
		if order.Order_type != "dine_in" && order.Order_type != "takeaway" {
			http.Error(w, "Incorrect order type", http.StatusBadRequest)
			return
		}
		table_number, err := strconv.Atoi(order.Table_number)
		if order.Order_type == "takeaway" {
			table_number = 0
		} else {
			if err != nil {
				fmt.Println(err)
				http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
				return
			}
		}
		if table_number > 20 || table_number < 0 {
			http.Error(w, "Table number out off bounds", http.StatusBadRequest)
		}
		cache.LoadMenu()
		err = models.CreateNewOrder(&order, table_number, user_id)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}
