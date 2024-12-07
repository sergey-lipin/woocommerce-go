package woocommerce

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/sergey-lipin/woocommerce-go/entity"
)

type settingService service

func (s settingService) Groups() (items []entity.SettingGroup, err error) {
	resp, err := s.httpClient.R().Get("/settings")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &items)
	}
	return
}
