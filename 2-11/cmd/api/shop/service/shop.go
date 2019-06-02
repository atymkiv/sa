package shop

import (
	sa "github.com/atymkiv/sa/model"
	"github.com/atymkiv/sa/pkg/utl/messages"
	"github.com/labstack/echo"
	"log"
)
const TOPIC = "shops"
type Service interface {
	View(echo.Context, string)(*sa.Shop, error)
	ViewAll(echo.Context) ([]sa.Shop, error)
	GetAddressByShopId(echo.Context, string) (*sa.Address, error)
}

func New(sdb DB, nats *messages.Service) *Shop {
	return &Shop{sdb: sdb, nats:nats}
}

type Shop struct {
	sdb DB
	nats *messages.Service
}

type DB interface {
	View(id string)(*sa.Shop, error)
	ViewAll() ([]sa.Shop, error)
	GetAddress(id int)(*sa.Address, error)
}

func (s *Shop) View(c echo.Context, id string) (*sa.Shop, error) {
	shop, err := s.sdb.View(id)
	if err != nil {
		return nil, err
	}
	if err := s.nats.PushMessage(shop, TOPIC); err != nil {
		log.Printf("failed pushing user into nuts; err: %v", err)
		return nil, err
	}
	return shop, nil
}

func (s *Shop) ViewAll(c echo.Context) ([]sa.Shop, error) {
	shops, err := s.sdb.ViewAll()
	if err != nil {
		return nil, err
	}
	return shops, nil
}

func(s *Shop) GetAddressByShopId(ctx echo.Context, id string) (*sa.Address, error) {
	shop, err := s.sdb.View(id)
	if err != nil{
		return nil, err
	}
	address, err := s.sdb.GetAddress(shop.Address_id)
	if err != nil {
		return nil, err
	}
	return address, nil
}
