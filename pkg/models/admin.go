package models

import (
	"database/sql"
	"fmt"

	"github.com/peopleig/food-ordering-go/pkg/types"
)

func GetAllOrders(orders *[]types.Order) error {
	query := `SELECT order_id, user_id, order_type, table_number, status FROM Orders ORDER BY order_id;`
	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ordered types.Order
		err := rows.Scan(&ordered.OrderId, &ordered.UserId, &ordered.Order_type, &ordered.Table_number, &ordered.Status)
		if err != nil {
			return err
		}
		*orders = append(*orders, ordered)
	}

	return rows.Err()
}

func ApproveUser(user_id int) error {
	query := `UPDATE User SET approved = TRUE WHERE user_id = ?`
	result, err := DB.Exec(query, user_id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", user_id)
	}

	return nil
}

func RemoveUserRequest(userId int) error {
	query := `DELETE FROM User WHERE user_id = ?`
	_, err := DB.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func GetUnapprovedUsers(uausers *[]types.UnApprovedUser) error {
	query := `SELECT user_id, CONCAT(first_name, ' ', last_name) AS name, role FROM User WHERE approved = FALSE`
	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var uauser types.UnApprovedUser
		err := rows.Scan(&uauser.UserId, &uauser.Name, &uauser.Role)
		if err != nil {
			return err
		}
		*uausers = append(*uausers, uauser)
	}
	return rows.Err()
}

func GetAllCategories(categories *[]types.Categories) error {
	query := `SELECT * FROM Categories`
	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var category types.Categories
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		if err != nil {
			return err
		}
		*categories = append(*categories, category)
	}
	return rows.Err()
}

func AddDish(newDish types.NewDish, price float32, isVeg bool, url string, spiceLevel int) error {
	var categoryID int
	query := `SELECT category_id FROM Categories WHERE category_name = ?`
	err := DB.QueryRow(query, newDish.Category).Scan(&categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Category not found")
			return err
		} else {
			return err
		}
	}
	query = `INSERT INTO Items(item_name, category_id, price, description, item_image_url, is_veg, spice_level) 
	VALUES (?,?,?,?,?,?,?)`
	_, err = DB.Exec(query, newDish.DishName, categoryID, price, newDish.Description, url, isVeg, spiceLevel)
	if err != nil {
		return err
	}
	return nil
}
