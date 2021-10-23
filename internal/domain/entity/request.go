package entity

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Jsonrpc string      `json:"jsonrpc" validate:"oneof=2.0" faker:"oneof: 2.0"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      uint64      `json:"id" validate:"required,gt=0" faker:"boundary_start=100, boundary_end=1000"`
}

func (r *Request) Validate() error {
	v := validator.New()
	if err := v.Struct(r); err != nil {
		return fmt.Errorf("invalid request")
	}
	switch r.Params.(type) {
	case []interface{}, map[string]interface{}:
		return nil
	default:
		return fmt.Errorf("invalid request")
	}
}

func FakeRequest(method string, params interface{}) Request {
	var req Request
	_ = faker.FakeData(&req)
	req.Method = method
	req.Params = params
	return req
}
