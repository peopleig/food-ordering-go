package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func BillHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	fmt.Println(user_id)
	var mybills []types.MyBills
	err := models.GetBills(user_id, &mybills)
	if err != nil {
		http.Error(w, "error finding your bills", http.StatusInternalServerError)
		return
	}
	data := types.BillData{
		Title:  "Bill",
		MyBill: mybills,
	}
	utils.RenderTemplate(w, "bill", data)
}

func BillPayerHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id").(int)
	var billpay types.BillPay
	err := json.NewDecoder(r.Body).Decode(&billpay)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = models.PaidbyUser(&billpay, user_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/menu", http.StatusSeeOther)
}
