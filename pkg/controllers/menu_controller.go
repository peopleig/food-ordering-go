package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	role := r.Context().Value("role").(string)
	switch r.Method {
	case http.MethodGet:
		var items []types.Item
		err := models.GetAllItems(&items)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Unable to access the menu", http.StatusInternalServerError)
			return
		}
		data := types.MenuData{
			Title: "Menu",
			Items: items,
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
		fmt.Printf("Got order: %+v\n", order)
		err = models.CreateNewOrder(&order, user_id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error in Creating Order", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}
