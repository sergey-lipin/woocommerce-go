package woocommerce

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/sergey-lipin/woocommerce-go/entity"
)

type systemStatusService service

func (s systemStatusService) All() (item entity.SystemStatus, err error) {
	resp, err := s.httpClient.R().Get("/system_status")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
