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
	if err := calcMapper.IsValid(); err != nil {
		return calcMapper.ToResponseObject(), err
	}
	calcMapper.Calculate()
	return calcMapper.ToResponseObject(), nil
}
