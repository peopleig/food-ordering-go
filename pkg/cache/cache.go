package cache

import (
	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/models"
)

func LoadMenu() error {
	if !config.MenuCacheLoaded {
		items, err := models.GetAllItems()
		if err != nil {
			return err
		}
		config.MenuCache = items
		config.MenuCacheLoaded = true
	}
	return nil
}

func RefreshMenuCache() {
	config.MenuCache = nil
	config.MenuCacheLoaded = false
}
