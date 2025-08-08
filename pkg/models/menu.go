package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/peopleig/food-ordering-go/pkg/types"
	"github.com/peopleig/food-ordering-go/pkg/utils"
)

func GetAllItems(items *[]types.Item) error {
	query := `SELECT i.item_id, i.item_name, i.price, i.description, i.item_image_url, i.is_veg, i.spice_level, c.category_name
	FROM Items i JOIN Categories c ON i.category_id = c.category_id;`
	rows, err := DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var i types.Item
		if err := rows.Scan(&i.Item_id, &i.Item_name, &i.Price, &i.Description, &i.Item_img, &i.Is_veg, &i.Spice_level, &i.Category_name); err != nil {
			return err
		}
		*items = append(*items, i)
	}
	return rows.Err()
}

func CreateNewOrder(order *types.OrderRequest, user_id int) error {
	table_number, _ := strconv.Atoi(order.Table_number)
	if order.Order_type != "dine_in" && order.Order_type != "takeaway" {
		return fmt.Errorf("incorrect order_type")
	}
	if order.Order_type == "takeaway" {
		table_number = 0
	}
	total, _ := utils.CalculateTotal(order.Cart)
	query := `INSERT INTO Orders (user_id, instructions, order_type, table_number, status, total_cost) VALUES (?, ?, ?, ?, 'preparing', ?)`
	res, err := DB.Exec(query, user_id, order.Special_instructions, order.Order_type, table_number, total)
	if err != nil {
		return err
	}
	order_id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	values := []interface{}{}
	placeholders := []string{}

	for _, item := range order.Cart {
		placeholders = append(placeholders, "(?, ?, ?)")
		values = append(values, order_id, item.Item_id, item.Quantity)
	}

	query = fmt.Sprintf(`INSERT INTO Ordered_Items (order_id, item_id, quantity) VALUES %s`, strings.Join(placeholders, ","))
	_, err = DB.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}
