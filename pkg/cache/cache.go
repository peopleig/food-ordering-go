package cache

import (
	"github.com/peopleig/food-ordering-go/pkg/config"
	"github.com/peopleig/food-ordering-go/pkg/models"
	"github.com/peopleig/food-ordering-go/pkg/types"
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

func LoadCategory() error {
	if !config.CategoryCacheLoaded {
		var categories []types.Categories
		err := models.GetAllCategories(&categories)
		if err != nil {
			return err
		}
		config.CategoryCache = categories
		config.CategoryCacheLoaded = true
	}
	return nil
}

func RefreshMenuCache() {
	config.MenuCache = nil
	config.MenuCacheLoaded = false
}
