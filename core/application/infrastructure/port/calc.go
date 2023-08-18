package application

import (
	"context"

	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
)

type CalcPort interface {
	Calculate(ctx context.Context, calcRequest me.CalcRequest) (mo.Response, error)
}
