package utils

import (
	"github.com/peopleig/food-ordering-go/pkg/types"
)

func CalculateTotal(cart []types.CartItem) (int, error) {
	total := 0
	for _, item := range cart {
		if item.Quantity <= 0 {
			return -1, nil
		}
		total += item.Price * item.Quantity
	}
	return total, nil
}
