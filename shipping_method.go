package woocommerce

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"

	"github.com/matewolf/woocommerce-go/entity"
)

type shippingMethodService service

// All list all shipping methods
func (s shippingMethodService) All() (items []entity.ShippingMethod, err error) {
	resp, err := s.httpClient.R().Get("/shipping_methods")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}

// One retrieve a shipping method
func (s shippingMethodService) One(id int) (item entity.ShippingMethod, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/shipping_methods/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
