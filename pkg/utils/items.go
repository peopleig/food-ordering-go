package utils

import (
	"strings"

	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/types"
)

func CalculateTotal(cart []types.CartItem) (uint, error) {
	var total uint
	for _, item := range cart {
		total += config.MenuCache[item.Item_id].Price * (item.Quantity)
	}
	return total, nil
}

func DishExists(dish_name string) bool {
	for _, dish := range config.MenuCache {
		if strings.EqualFold(dish.Item_name, dish_name) {
			return true
		}
	}
	return false
}
