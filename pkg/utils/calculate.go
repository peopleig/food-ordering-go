package utils

import (
	"github.com/peopleig/food-ordering-go/pkg/types"
)

func CalculateTotal(cart []types.CartItem) (int, error) {
	total := 0
	for _, item := range cart {
		total += item.Price * item.Quantity

	}
	return total, nil
}
