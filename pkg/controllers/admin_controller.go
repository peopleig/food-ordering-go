package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
