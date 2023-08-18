package domain

import (
	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (a Service) IsCalcRequestEntityValid(calcRequest me.CalcRequest) error {
	return calcRequest.IsValid()
}
