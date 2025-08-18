package models

import (
	"database/sql"
)

func CheckForPendingBills(userId int) (bool, error) {
	var orderId int
	query := `SELECT order_id FROM Orders WHERE user_id = ? AND status IN ('payment_pending', 'preparing') LIMIT 1`
	err := DB.QueryRow(query, userId).Scan(&orderId)

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
