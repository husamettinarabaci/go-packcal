package application

import (
	"context"

	as "github.com/husamettinarabaci/go-packcal/core/application/service"
	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
)

type CalcAdapter struct {
	service as.Service
}

func NewCalcAdapter(s as.Service) CalcAdapter {
	return CalcAdapter{
		service: s,
	}
}

func (a CalcAdapter) Calculate(ctx context.Context, calcRequest me.CalcRequest) (me.CalcResponse, error) {
	return a.service.Calculate(ctx, calcRequest)
}
