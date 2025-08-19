package cache

import (
	"encoding/json"
	"fmt"

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
		jsonData, err := json.Marshal(config.MenuCache)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return err
		}
		config.ByteMenuCache = jsonData
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
