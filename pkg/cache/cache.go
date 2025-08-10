package cache

import (
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
)

var MenuCache []types.Item
var MenuCacheLoaded bool

func LoadMenu() error {
	if !MenuCacheLoaded {
		items, err := models.GetAllItems()
		if err != nil {
			return err
		}
		MenuCache = items
		MenuCacheLoaded = true
	}
	return nil
}

func RefreshMenuCache() {
	MenuCache = nil
	MenuCacheLoaded = false
}
