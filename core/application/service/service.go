package application

import (
	"context"

	ip "github.com/husamettinarabaci/go-packcal/core/application/infrastructure/port"
	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
	mi "github.com/husamettinarabaci/go-packcal/core/domain/model/interface"
	ds "github.com/husamettinarabaci/go-packcal/core/domain/service"
)

type Service struct {
	domainService ds.Service
	calc          ip.CalcPort
	logger        ip.LogPort
}

func NewService(domainService ds.Service, calcPort ip.CalcPort, logger ip.LogPort) Service {
	return Service{
		domainService: domainService,
		calc:          calcPort,
		logger:        logger,
	}
}

func (a Service) Log(ctx context.Context, operationName string, logData mi.Loggable) {
	a.logger.Log(ctx, operationName, logData)
}

func (a Service) Calculate(ctx context.Context, calcRequest me.CalcRequest) (me.CalcResponse, error) {
	operationName := "ExecuteCalc"
	a.Log(ctx, operationName, calcRequest)
	if err := a.domainService.IsCalcRequestEntityValid(calcRequest); err != nil {
		return me.CalcResponse{
			Id: calcRequest.Id,
		}, err
	}
	response, err := a.calc.Calculate(ctx, calcRequest)
	calcResponse := me.NewCalcResponse(calcRequest.Id, response)
	a.Log(ctx, operationName, calcResponse)
	return calcResponse, err
}
