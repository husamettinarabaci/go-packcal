package presentation

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type CalcRequest struct {
	Item         int    `json:"item" form:"item" binding:"required"`
	PackSizes    []int  `json:"pack_sizes"`
	PackSizesStr string `json:"pack_sizes_str" form:"pack_sizes_str"`
}

func (a CalcRequest) ToJson() string {
	return tjson.ToJson(a)
}

func (e CalcRequest) FromJson(i string) CalcRequest {
	return tjson.FromJson[CalcRequest](i)
}

func NewCalcRequest(item int, packSizes []int) CalcRequest {
	return CalcRequest{
		Item:      item,
		PackSizes: packSizes,
	}
}

func (o CalcRequest) IsEmpty() bool {
	return o.ToJson() == CalcRequest{}.ToJson()
}

func (a CalcRequest) ToCalcRequestEntity() me.CalcRequest {
	if a.PackSizes == nil {
		a.PackSizes = mo.DefaultPackSizes
	}
	return me.NewCalcRequest(
		uuid.New(),
		mo.NewCalc(
			a.Item,
			a.PackSizes,
		),
	)
}

func (a *CalcRequest) FillPackSizes() {
	if a.PackSizesStr != "" {
		sizes := strings.Split(a.PackSizesStr, ",")
		a.PackSizes = make([]int, 0)
		for _, v := range sizes {
			if v != "" {
				size, err := strconv.Atoi(v)
				if err != nil {
					a.PackSizes = make([]int, 0)
					break
				}
				a.PackSizes = append(a.PackSizes, size)
			}
		}
	}
	if a.PackSizes == nil {
		a.PackSizes = mo.DefaultPackSizes
	}
	if len(a.PackSizes) == 0 {
		a.PackSizes = mo.DefaultPackSizes
	}
}

func (a *CalcRequest) IsValid() error {
	if a.IsEmpty() {
		return mo.ErrInvalidInput
	}
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
