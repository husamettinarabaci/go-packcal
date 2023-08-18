package application

import (
	"context"

	mi "github.com/husamettinarabaci/go-packcal/core/domain/model/interface"
)

type LogPort interface {
	Log(ctx context.Context, source string, logData mi.Loggable)
}
