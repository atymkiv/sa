package sa

import "github.com/jinzhu/gorm"

type (
	Shop struct {
		Shop_Id    int `json:"shop_id"`
		Address_id int `json:"address_id"`
		Size       int `json:"size"`
		Workers    int `json:"workers"`
	}
	Address struct {
		gorm.Model
		City   string `json:"city"`
		Street string `json:"street"`
		Number string `json:"number"`
	}
)
