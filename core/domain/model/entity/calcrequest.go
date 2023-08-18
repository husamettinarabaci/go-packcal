package domain

import (
	"github.com/google/uuid"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type CalcRequest struct {
	Id   uuid.UUID `json:"id"`
	Calc mo.Calc   `json:"calc"`
}

func (e CalcRequest) ToJson() string {
	return tjson.ToJson(e)
}

func (a CalcRequest) FromJson(i string) CalcRequest {
	return tjson.FromJson[CalcRequest](i)
}

func NewCalcRequest(id uuid.UUID, calc mo.Calc) CalcRequest {
	return CalcRequest{
		Id:   id,
		Calc: calc,
	}
}

func (o CalcRequest) IsEmpty() bool {
	return o.ToJson() == CalcRequest{}.ToJson()
}

func (o CalcRequest) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o CalcRequest) IsEqual(i CalcRequest) bool {
	return o.ToJson() == i.ToJson()
}

func (o CalcRequest) IsValid() error {
	if o.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if o.Id == uuid.Nil {
		return mo.ErrInvalidInput
	}
	if o.Calc.IsEmpty() {
		return mo.ErrInvalidInput
	}
	if err := o.Calc.IsValid(); err != nil {
		return err
	}
	return nil
}
