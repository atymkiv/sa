package sa

import "github.com/jinzhu/gorm"

type (
	Good struct {
		gorm.Model
		Name        int     `json:"name"`
		Price       float64 `json:"price"`
		Category_id int     `json:"category_id"`
	}
	Category struct {
		gorm.Model
		Name int `json:"name"`
	}
)
