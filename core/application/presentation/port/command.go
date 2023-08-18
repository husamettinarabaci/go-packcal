package application

import (
	"context"

	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
)

type CommandPort interface {
	Calculate(ctx context.Context, calcRequest me.CalcRequest) (me.CalcResponse, error)
}
