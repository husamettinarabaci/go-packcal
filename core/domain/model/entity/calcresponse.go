package domain

import (
	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type CalcResponse struct {
	Id       uuid.UUID   `json:"id"`
	Response mo.Response `json:"response"`
}

func (e CalcResponse) ToJson() string {
	return tjson.ToJson(e)
}

func (a CalcResponse) FromJson(i string) CalcResponse {
	return tjson.FromJson[CalcResponse](i)
}

func NewCalcResponse(id uuid.UUID, response mo.Response) CalcResponse {
	return CalcResponse{
		Id:       id,
		Response: response,
	}
}

func (o CalcResponse) IsEmpty() bool {
	return o.ToJson() == CalcResponse{}.ToJson()
}

func (o CalcResponse) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o CalcResponse) IsEqual(i CalcResponse) bool {
	return o.ToJson() == i.ToJson()
}

func (o CalcResponse) IsValid() error {
	if o.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if o.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if o.Response.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if err := o.Response.IsValid(); err != nil {
		return err
	}
	return nil
}
