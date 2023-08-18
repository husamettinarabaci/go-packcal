package domain

import (
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type Calc struct {
	Item      int   `json:"item"`
	PackSizes []int `json:"pack_sizes"`
}

func (o Calc) ToJson() string {
	return tjson.ToJson(o)
}

func (a Calc) FromJson(i string) Calc {
	return tjson.FromJson[Calc](i)
}

func NewCalc(item int, packSizes []int) Calc {
	return Calc{
		Item:      item,
		PackSizes: packSizes,
	}
}

func (o Calc) IsEmpty() bool {
	return o.ToJson() == Calc{}.ToJson()
}

func (o Calc) IsNotEmpty() bool {
	return !o.IsEmpty()
}

func (o Calc) IsEqual(i Calc) bool {
	return o.ToJson() == i.ToJson()
}

func (o Calc) IsValid() error {
	if o.IsEmpty() {
		return ErrInvalidInput
	}
	if o.Item <= 0 {
		return ErrInvalidInput
	}
	if o.PackSizes == nil {
		return ErrInvalidInput
	}
	if len(o.PackSizes) == 0 {
		return ErrInvalidInput
	}
	for _, v := range o.PackSizes {
		if v <= 0 {
			return ErrInvalidInput
		}
	}
	return nil
}

var DefaultPackSizes = []int{250, 500, 1000, 2000, 5000}
