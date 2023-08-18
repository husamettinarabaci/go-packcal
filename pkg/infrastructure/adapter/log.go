package infrastructure

import (
	"context"
	"fmt"

	mi "github.com/husamettinarabaci/go-packcal/core/domain/model/interface"
	tconfig "github.com/husamettinarabaci/go-packcal/tool/config"
)

type LogAdapter struct {
}

func NewLogAdapter() LogAdapter {
	adapter := LogAdapter{}
	return adapter
}

func (a LogAdapter) Log(ctx context.Context, source string, logData mi.Loggable) {
	if tconfig.GetLogConfigInstance().Logger.Console {
		fmt.Println(source, " : ", logData.ToJson())
	}
}
