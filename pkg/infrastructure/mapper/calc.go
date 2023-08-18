package infrastructure

import (
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type Calc struct {
	Item      int   `json:"item"`
	PackSizes []int `json:"pack_sizes"`
}

func (a Calc) ToJson() string {
	return tjson.ToJson(a)
}

func (e Calc) FromJson(i string) Calc {
	return tjson.FromJson[Calc](i)
}

func NewCalc(item int, packSizes []int) Calc {
	return Calc{
		Item:      item,
		PackSizes: packSizes,
	}
}

func FromCalcObject(calc mo.Calc) Calc {
	return NewCalc(
		calc.Item,
		calc.PackSizes,
	)
}

func (a Calc) IsValid() error {
	if a.Item <= 0 {
		return mo.ErrInvalidInput
	}
	if a.PackSizes == nil {
		return mo.ErrInvalidInput
	}
	if len(a.PackSizes) == 0 {
		return mo.ErrInvalidInput
	}
	for _, v := range a.PackSizes {
		if v <= 0 {
			return mo.ErrInvalidInput
		}
	}
	return nil
}
