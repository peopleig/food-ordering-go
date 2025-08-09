package models

import (
	"fmt"

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
	query := `INSERT INTO Payment (order_id, tip_amount, discount_reward_points, amount_paid, payment_status) VALUES (?, ?, 0, ?, 'paid')`
	_, err := DB.Exec(query, billpay.OrderId, billpay.Tip, billpay.Tip*2)
	if err != nil {
		return err
	}
	query = `UPDATE Orders SET status = 'completed' WHERE order_id = ?`
	result, err := DB.Exec(query, billpay.OrderId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no order found with id %d", billpay.OrderId)
	}

	return nil
}
