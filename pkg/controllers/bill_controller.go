package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func SingleBillHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, _ := strconv.Atoi(vars["order_id"])
	status := vars["status"]
	var isSame bool
	status, isSame = models.ConfirmOrderStatus(orderId, status)
	if !isSame {
		url := "/bill/" + status + "/" + vars["order_id"]
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
	var billOrder types.BillOrder
	var Contents []types.OrderContents
	var completeBill types.CompleteBill
	switch status {
	case "preparing", "payment_pending":
		err := models.GetSingleBill(orderId, &Contents, &billOrder)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "couldn't access DB", http.StatusInternalServerError)
			return
		}
	case "completed":
		err := models.GetFinalBill(orderId, &Contents, &completeBill)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "couldn't access DB", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "incorrect order status", http.StatusBadRequest)
		return
	}
	switch status {
	case "preparing":
		data := types.SingleBill{
			Title:    "Bill",
			Contents: Contents,
			Order:    billOrder,
		}
		utils.RenderTwoTemplates(w, "bill", "preparing", data)
	case "payment_pending":
		data := types.SingleBill{
			Title:    "Bill",
			Contents: Contents,
			Order:    billOrder,
		}
		utils.RenderTwoTemplates(w, "bill", "pay", data)
	case "completed":
		data := types.FinalBill{
			Title:    "Bill",
			Contents: Contents,
			Order:    completeBill,
		}
		utils.RenderTwoTemplates(w, "bill", "complete", data)
	default:
		http.Error(w, "incorrect order status", http.StatusBadRequest)
		return
	}
}
