package domain

import (
	"testing"

	"github.com/google/uuid"
	me "github.com/husamettinarabaci/go-packcal/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
)

func Test_IsCalcRequestEntityValid(t *testing.T) {

	var tests = []struct {
		name        string
		calcRequest me.CalcRequest
		want        error
	}{
		{name: "empty", calcRequest: me.CalcRequest{}, want: mo.ErrInvalidInput},
		{name: "non_values", calcRequest: me.CalcRequest{Id: uuid.New(), Calc: mo.Calc{}}, want: mo.ErrInvalidInput},
		{name: "valid", calcRequest: me.CalcRequest{Id: uuid.New(), Calc: mo.Calc{Item: 250, PackSizes: []int{250, 500, 1000, 2000, 5000}}}, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NewService().IsCalcRequestEntityValid(tt.calcRequest)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
