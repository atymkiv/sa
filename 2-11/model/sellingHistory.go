package sa

import "github.com/jinzhu/gorm"

type (
	SellingHistory struct {
		gorm.Model
		Shop_id  int `json:"shop_id"`
		Good_id  int `json:"good_id"`
		Month    int `json:"month"`
		Quantity int `json:"quantity"`
	}
)
