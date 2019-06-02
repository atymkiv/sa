package transport

import (
	shop "github.com/atymkiv/sa/cmd/api/shop/service"
	"github.com/labstack/echo"
	"net/http"
)

type HTTP struct {
	svc shop.Service
}

// NewHTTP creates new shop http service
func NewHTTP(svc shop.Service, er *echo.Group) {
	h := HTTP{svc}
	// swagger:route GET /shop/{id} shop shopView
	// Shows shop by id.
	// responses:
	//  200: shopResp
	er.GET("/:id", h.getById)
	// swagger:route GET /shop/all shops shopsView
	// Shows all shops.
	// responses:
	//  200: shopsResp
	er.GET("/all", h.getAll)

	// swagger:route GET /shop/{id} shops addressGet
	//	---
	// summary: Gets address by shop ID
	// description: Gets address by shop ID
	// responses:
	//  200: addressResp
	er.GET("/addr/:id", h.getAddress)


}


func (h *HTTP) getAll(c echo.Context) error {
	shops, err := h.svc.ViewAll(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, shops)
}

type addressResp struct {
	City string
	Number string
	Street string
}

func (h *HTTP) getById(c echo.Context) error {
	id := c.Param("id")
	shop, err := h.svc.View(c, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, shop)
}

func (h *HTTP) getAddress(c echo.Context) error {
	id := c.Param("id")
	address, err := h.svc.GetAddressByShopId(c, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, addressResp{
		City:address.City,
		Street:address.Street,
		Number:address.Number,
	})
}