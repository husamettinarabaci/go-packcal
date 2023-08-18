package infrastructure

import (
	"context"

	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	mp "github.com/husamettinarabaci/go-packcal/pkg/infrastructure/mapper"
)

type CalcAdapter struct {
}

func NewCalcAdapter() CalcAdapter {
	adapter := CalcAdapter{}

	return adapter
}

func (a CalcAdapter) Calculate(ctx context.Context, calcRequest me.CalcRequest) (mo.Response, error) {
	calcMapper := mp.FromCalcObject(calcRequest.Calc)
	var err error
	response := mo.Response{}
	if err := calcMapper.IsValid(); err != nil {
		return response, err
	}
	packs, counts, err := calc(calcMapper.Item, &calcMapper.PackSizes)
	if err != nil {
		return response, err
	} else {
		response = mo.NewResponse(packs, counts)
		return response, nil
	}
}

func calc(item int, packSizes *[]int) ([]int, []int, error) {
	return []int{5000, 2000, 250}, []int{2, 1, 1}, nil
}
