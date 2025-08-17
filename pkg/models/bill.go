package models

import (
	"fmt"
	"strings"

	"github.com/peopleig/food-ordering-go/pkg/types"
)

func GetBills(user_id int, mybills *[]types.MyBills) error {
	query := `SELECT order_id, status, total_cost FROM Orders WHERE user_id = ?`
	rows, err := DB.Query(query, user_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var mybill types.MyBills
		err = rows.Scan(&mybill.OrderId, &mybill.Status, &mybill.Price)
		if err != nil {
			return err
		}
		*mybills = append(*mybills, mybill)
	}

	return rows.Err()
}

func PaidbyUser(billpay *types.BillPay, user_id int) error {
	query := `SELECT total_cost FROM Orders WHERE order_id = ?`
	var initial_cost uint
	err := DB.QueryRow(query, billpay.OrderId).Scan(&initial_cost)
	fmt.Println(initial_cost)
	if err != nil {
		return err
	}
	query = `INSERT INTO Payment (order_id, tip_amount, discount_reward_points, amount_paid, payment_status) VALUES (?, ?, 0, ?, 'paid')`
	_, err = DB.Exec(query, billpay.OrderId, billpay.Tip, initial_cost+billpay.Tip)
	if err != nil {
		return err
	}
	fmt.Println(initial_cost + billpay.Tip)
	query = `UPDATE Orders SET status = 'completed' WHERE order_id = ?`
	_, err = DB.Exec(query, billpay.OrderId)
	if err != nil {
		return err
	}
	return nil
}

func ConfirmOrderStatus(orderId int, status string) (string, bool) {
	query := `SELECT status FROM Orders WHERE order_id = ?`
	var temp_status string
	_ = DB.QueryRow(query, orderId).Scan(&temp_status)
	return temp_status, strings.EqualFold(temp_status, status)
}

func GetSingleBill(orderId int, contents *[]types.OrderContents, billOrder *types.BillOrder) error {
	query := `SELECT o.order_id, o.user_id, o.instructions, o.status, o.table_number, o.order_type, o.total_cost, CONCAT(u.first_name, ' ', u.last_name) AS name 
	FROM Orders o
	JOIN User u ON o.user_id = u.user_id 
	WHERE o.order_id = ?`
	row := DB.QueryRow(query, orderId)
	err := row.Scan(&billOrder.OrderId, &billOrder.UserId, &billOrder.Instructions, &billOrder.Status, &billOrder.TableNumber, &billOrder.OrderType, &billOrder.TotalCost, &billOrder.UserName)
	if err != nil {
		return err
	}
	query = `SELECT i.item_name, i.price, oi.quantity 
	FROM Ordered_Items oi
	JOIN Items i ON oi.item_id = i.item_id
	WHERE oi.order_id = ?`
	rows, err := DB.Query(query, orderId)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var content types.OrderContents
		err := rows.Scan(&content.ItemName, &content.Price, &content.Quantity)
		if err != nil {
			return err
		}
		*contents = append(*contents, content)
	}
	return rows.Err()
}

func GetFinalBill(orderId int, contents *[]types.OrderContents, completeBill *types.CompleteBill) error {
	query := `SELECT o.order_id, o.user_id, o.instructions, o.status, o.table_number, o.order_type, o.total_cost, p.tip_amount, p.amount_paid, CONCAT(u.first_name, ' ', u.last_name) AS name
	FROM Orders o
	JOIN Payment p ON o.order_id = p.order_id
	JOIN User u ON o.user_id = u.user_id 
	WHERE o.order_id = ?`
	row := DB.QueryRow(query, orderId)
	err := row.Scan(&completeBill.OrderId, &completeBill.UserId, &completeBill.Instructions, &completeBill.Status, &completeBill.TableNumber, &completeBill.OrderType, &completeBill.TotalCost, &completeBill.TipPaid, &completeBill.AmtPaid, &completeBill.UserName)
	if err != nil {
		return err
	}
	query = `SELECT i.item_name, i.price, oi.quantity 
	FROM Ordered_Items oi
	JOIN Items i ON oi.item_id = i.item_id
	WHERE oi.order_id = ?`
	rows, err := DB.Query(query, orderId)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var content types.OrderContents
		err := rows.Scan(&content.ItemName, &content.Price, &content.Quantity)
		if err != nil {
			return err
		}
		*contents = append(*contents, content)
	}
	return rows.Err()
}
