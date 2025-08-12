package utils

import (
	"strings"

	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/types"
)

func CalculateTotal(cart []types.CartItem) (float32, error) {
	var total float32
	for _, item := range cart {
		if item.Quantity <= 0 {
			return -1, nil
		}
		total += config.MenuCache[item.Item_id].Price * float32(item.Quantity)
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
