package presentation

import (
	"fmt"

	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
)

type CalcResponse struct {
	Packs []string `json:"packs"`
}

func (a CalcResponse) ToJson() string {
	return tjson.ToJson(a)
}

func (e CalcResponse) FromJson(i string) CalcResponse {
	return tjson.FromJson[CalcResponse](i)
}

func NewCalcResponse(packs []string) CalcResponse {
	return CalcResponse{
		Packs: packs,
	}
}

func (o CalcResponse) IsEmpty() bool {
	return o.ToJson() == CalcResponse{}.ToJson()
}

func FromResponseObject(response mo.Response) CalcResponse {
	var packs []string
	for i := 0; i < len(response.Packs); i++ {
		packs = append(packs, fmt.Sprintf("%d x %d", response.Counts[i], response.Packs[i]))
	}
	return NewCalcResponse(
		packs,
	)
}
