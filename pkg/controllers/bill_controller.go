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
	var mybills []types.MyBills
	err := models.GetBills(user_id, &mybills)
	if err != nil {
		http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
		return
	}
	data := types.BillData{
		Title:  "Bill",
		MyBill: mybills,
	}
	utils.RenderTemplate(w, "bill", data)
}

func GetMyBills(w http.ResponseWriter, r *http.Request) {
	unPaid := r.URL.Query().Get("error")
	user_id := r.Context().Value("user_id").(int)
	role := r.Context().Value("role").(string)
	show := false
	if unPaid == "unpaid" {
		show = true
	}
	var myBills []types.ShortBillForm
	err := models.GetShortBills(user_id, &myBills)
	if err != nil {
		http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
		return
	}
	data := types.ShortBillData{
		Title:      "Bills",
		ShortBills: myBills,
		Role:       role,
		Show:       show,
	}
	utils.RenderTemplate(w, "bills", data)
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
	orderIdString := strconv.Itoa(billpay.OrderId)
	redirectURL := "/bill/completed/" + orderIdString + "?show=success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"redirect": redirectURL,
	})
}

func SingleBillHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value("role").(string)
	vars := mux.Vars(r)
	orderId, _ := strconv.Atoi(vars["order_id"])
	status := vars["status"]
	showMsg := r.URL.Query().Get("show")
	show := false
	message := ""
	if showMsg == "success" {
		show = true
		message = "Payment successful!"
	}
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
			http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
			return
		}
	case "completed":
		err := models.GetFinalBill(orderId, &Contents, &completeBill)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/error?error=internal", http.StatusSeeOther)
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
			Role:     role,
		}
		utils.RenderTwoTemplates(w, "bill", "preparing", data)
	case "payment_pending":
		data := types.SingleBill{
			Title:    "Bill",
			Contents: Contents,
			Order:    billOrder,
			Role:     role,
		}
		utils.RenderTwoTemplates(w, "bill", "pay", data)
	case "completed":
		data := types.FinalBill{
			Title:    "Bill",
			Contents: Contents,
			Order:    completeBill,
			Show:     show,
			Message:  message,
			Role:     role,
		}
		utils.RenderTwoTemplates(w, "bill", "complete", data)
	default:
		http.Error(w, "incorrect order status", http.StatusBadRequest)
		return
	}
}
