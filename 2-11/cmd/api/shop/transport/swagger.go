package transport

import (
	sa "github.com/atymkiv/sa/model"
)

// Shops feed model response
// swagger:response shopsResp
type swaggShopsResponse struct {
	//in:body
	Body struct{
		Shops []sa.Shop
	}
}


// swagger:response addressResp
type swaggAddressResponse struct {
	//in:body
	Body struct{
		City string
		Number string
		Street string

	}
}
