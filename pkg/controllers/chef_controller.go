package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func ChefHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	switch r.Method {
	case http.MethodGet:
		var items []types.Ordered
		err := models.GetAllOrderedItems(&items)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error while accessing DB", http.StatusBadRequest)
			return
		}
		data := types.OrdersData{
			Title:  "Chef",
			Items:  items,
			UserId: user_id,
		}
		utils.RenderTemplate(w, "chef", data)
	case http.MethodPatch:
		var assign types.ChefAssignRequest

		err := json.NewDecoder(r.Body).Decode(&assign)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err = models.AssignToChef(&assign)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to update", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})

	case http.MethodPost:
		var done types.ChefAssignRequest
		err := json.NewDecoder(r.Body).Decode(&done)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		err = models.DoneByChef(&done)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		toUpdate, err := models.CheckCompletion(done.OrderID)
		if err != nil {
			fmt.Println(err)
		}
		if toUpdate {
			err := models.UpdateOrderStatus(done.OrderID)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "failed to update order status", http.StatusInternalServerError)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
	}
}
