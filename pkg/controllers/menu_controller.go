package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
		err := cache.LoadMenu()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to access the menu", http.StatusInternalServerError)
			return
		}
		var myBills []types.MyBills
		err = models.GetBills(user_id, &myBills)
		if err != nil {
			http.Error(w, "error finding your bills", http.StatusInternalServerError)
			return
		}
		data := types.MenuData{
			Title: "Menu",
			Bills: myBills,
			Items: config.MenuCache,
		}
		fmt.Println(user_id, role)
		utils.RenderTemplate(w, "menu", data)

	case http.MethodPost:
		var order types.OrderRequest
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Incorrect Cart data sent", http.StatusBadRequest)
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
				http.Error(w, "Error in parsing table number data", http.StatusInternalServerError)
				return
			}
		}
		if order.Order_type == "takeaway" {
			table_number = 0
		}
		fmt.Printf("Got order: %+v\n", order)
		cache.LoadMenu()
		err = models.CreateNewOrder(&order, table_number, user_id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error in Creating Order", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}
