package sa

import "github.com/jinzhu/gorm"

type (
	SupplyAndDemand struct {
		gorm.Model
		Month          int     `json:"month"`
		SupplyPerMonth float32 `json:"supply_per_month"`
		DemandPerMonth float32 `json:"demand_per_month"`
	}
)
