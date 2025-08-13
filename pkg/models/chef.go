package models

import (
	"github.com/peopleig/food-ordering-go/pkg/types"
)

func GetAllOrderedItems(items *[]types.Ordered) error {
	query := `SELECT oi.order_id, oi.item_id, i.item_name, oi.quantity, u.user_id AS chef_id, CONCAT(u.first_name, ' ', u.last_name) AS chef_name, o.instructions,o.order_type
        FROM Ordered_Items oi 
		JOIN Items i ON oi.item_id = i.item_id 
		JOIN Orders o ON oi.order_id = o.order_id
        JOIN User u ON oi.chef_id = u.user_id
		WHERE oi.dish_complete = false
        ORDER BY oi.order_id, oi.item_id;`

	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ordered types.Ordered
		err := rows.Scan(&ordered.OrderId, &ordered.ItemId, &ordered.ItemName, &ordered.Quantity, &ordered.ChefId, &ordered.ChefName, &ordered.Instructions, &ordered.Order_type)
		if err != nil {
			return err
		}
		ordered.Assigned = !(ordered.ChefId == 1)
		*items = append(*items, ordered)
	}

	return rows.Err()
}

func AssignToChef(assign *types.ChefAssignRequest) error {
	query := `UPDATE Ordered_Items SET chef_id = ? WHERE order_id = ? AND item_id = ? AND chef_id = 1;`
	_, err := DB.Exec(query, assign.ChefID, assign.OrderID, assign.ItemID)
	if err != nil {
		return err
	}
	return nil
}

func DoneByChef(done *types.ChefAssignRequest) error {
	query := `UPDATE Ordered_Items SET dish_complete = TRUE WHERE order_id = ? AND item_id = ? AND chef_id = ?`
	_, err := DB.Exec(query, done.OrderID, done.ItemID, done.ChefID)
	if err != nil {
		return err
	}
	return nil
}

func CheckCompletion(orderId int) (bool, error) {
	query := `SELECT dish_complete FROM Ordered_Items WHERE order_id = ?`
	rows, err := DB.Query(query, orderId)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	answer := true
	for rows.Next() {
		var this bool
		err := rows.Scan(&this)
		if err != nil {
			return false, err
		}
		answer = answer && this
	}
	return answer, nil
}

func UpdateOrderStatus(orderId int) error {
	query := `UPDATE Orders SET status='payment_pending' WHERE order_id = ?`
	_, err := DB.Exec(query, orderId)
	if err != nil {
		return err
	}
	return nil
}
