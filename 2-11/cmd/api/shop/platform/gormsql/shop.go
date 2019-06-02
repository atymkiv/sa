package gormsql

import (
	sa "github.com/atymkiv/sa/model"
	"github.com/atymkiv/sa/pkg/utl/gorm"

)

func NewShop(db *gorm.DB) *Shop {
	return &Shop{
		db:db,
	}
}

type Shop struct {
	db *gorm.DB
}
func (s *Shop) View(id string) (*sa.Shop, error) {
	shop := new(sa.Shop)
	if err := s.db.Where("shop_id = ?", id).Find(&shop).Error; err != nil {
		return nil, sa.ErrGeneric
	}
	return shop, nil
}

func (s *Shop) ViewAll() ([]sa.Shop, error) {
	var shops []sa.Shop
	if err := s.db.Find(&shops).Error; err != nil {
		return nil, sa.ErrGeneric
	}
	return shops, nil
}


func(s *Shop) GetAddress(id int)(*sa.Address, error) {
	address := new(sa.Address)
	if err := s.db.Where("id = ?", id).Find(&address).Error; err != nil {
		return nil, sa.ErrGeneric
	}

	return address, nil
}
