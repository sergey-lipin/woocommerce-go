package woocommerce

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity/order"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

type orderNoteService service

type OrderNotesQueryParams struct {
	QueryParams
	Type string `url:"type"`
}

func (m OrderNotesQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Type, validation.When(m.Type != "", validation.In("any", "customer", "internal").Error("无效的类型"))),
	)
}

func (s orderNoteService) All(orderId int, params OrderNotesQueryParams) (items []order.Note, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	urlValues, _ := query.Values(params)
	var res []order.Note
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get(fmt.Sprintf("/orders/%d/notes", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
		}
	}
	return
}

func (s orderNoteService) One(orderId, noteId int) (item order.Note, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create order note

type CreateOrderNoteRequest struct {
	Note string `json:"note"`
}

func (m CreateOrderNoteRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Note, validation.Required.Error("内容不能为空")),
	)
}

func (s orderNoteService) Create(orderId int, req CreateOrderNoteRequest) (item order.Note, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().
		SetBody(req).
		Post(fmt.Sprintf("/orders/%d/notes", orderId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

func (s service) Delete(orderId, noteId int, force bool) (item order.Note, err error) {
	resp, err := s.httpClient.R().
		SetBody(map[string]string{
			"force": strconv.FormatBool(force),
		}).
		Delete(fmt.Sprintf("/orders/%d/notes/%d", orderId, noteId))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}
