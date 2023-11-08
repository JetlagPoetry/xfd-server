package cache

import (
	"sync"
	"xfd-backend/database/db/enum"
)

var categoryLock sync.RWMutex
var categoryCache map[int32]*Category

func GetCategory() map[int32]*Category {
	categoryLock.RLock()
	defer categoryLock.RUnlock()
	return categoryCache
}

func GetCategoryByID(id int32) *Category {
	categoryLock.RLock()
	defer categoryLock.RUnlock()
	return categoryCache[id]
}

func SetCategory(new map[int32]*Category) {
	categoryLock.Lock()
	categoryCache = new
	defer categoryLock.Unlock()
}

type Category struct {
	ID               int32
	Name             string
	ParentCategoryID int32
	SubCategoryIDs   []int32
	Level            enum.GoodsCategoryLevel
	Image            string
}
