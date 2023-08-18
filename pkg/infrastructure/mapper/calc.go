package infrastructure

import (
	"sort"

	mo "github.com/husamettinarabaci/go-packcal/core/domain/model/object"
	tjson "github.com/husamettinarabaci/go-packcal/tool/json"
	tslice "github.com/husamettinarabaci/go-packcal/tool/slice"
)

var tSlice tslice.TSlice

type Calc struct {
	Item      int   `json:"item"`
	PackSizes []int `json:"pack_sizes"`
	Results   []int `json:"results"`
}

func (a Calc) ToJson() string {
	return tjson.ToJson(a)
}

func (e Calc) FromJson(i string) Calc {
	return tjson.FromJson[Calc](i)
}

func NewCalc(item int, packSizes []int) Calc {
	return Calc{
		Item:      item,
		PackSizes: packSizes,
	}
}

func FromCalcObject(calc mo.Calc) Calc {
	return NewCalc(
		calc.Item,
		calc.PackSizes,
	)
}

func (a Calc) IsValid() error {
	if a.Item <= 0 {
		return mo.ErrInvalidInput
	}
	if a.PackSizes == nil {
		return mo.ErrInvalidInput
	}
	if len(a.PackSizes) == 0 {
		return mo.ErrInvalidInput
	}
	for _, v := range a.PackSizes {
		if v <= 0 {
			return mo.ErrInvalidInput
		}
	}
	return nil
}

func (a Calc) ToResponseObject() mo.Response {
	var packs []int
	var counts []int
	resMap := make(map[int]int)
	for _, pack := range a.Results {
		resMap[pack] += 1
	}
	for k, v := range resMap {
		packs = append(packs, k)
		counts = append(counts, v)
	}

	return mo.NewResponse(
		packs,
		counts,
	)
}

func (a *Calc) Calculate() {
	sort.Sort(sort.Reverse(sort.IntSlice(a.PackSizes)))
	a.Results = a.CalculateCumulative(a.Item, a.PackSizes, []int{})
}

func (a Calc) CalculateCumulative(item int, packSizes []int, res []int) []int {
	//If the item is in the packSizes, return it
	if tSlice.Contains(packSizes, item) {
		return append(res, item)
	}
	remain := item

	if len(res) != 0 {
		temp := make([]int, len(packSizes))
		copy(temp, packSizes)
		sort.Ints(temp)
		lowValue, highValue := tSlice.FindEdgest(temp, remain)
		if remain >= lowValue && remain <= highValue {
			res = append(res, highValue)
			remain = remain - highValue
		} else if remain >= lowValue {
			res = append(res, lowValue)
			remain = remain - lowValue
		}
	}

	if remain > 0 {
		for _, packSize := range packSizes {
			if remain >= packSize {
				res = append(res, packSize)
				remain = remain - packSize
				return a.CalculateCumulative(remain, packSizes, res)
			}
		}
		if remain < packSizes[len(packSizes)-1] {
			res = append(res, packSizes[len(packSizes)-1])
		}

	}

	//If the sum of the results is in the packSizes, return it
	sumResults := tslice.TSlice{}.Sum(res)
	if tSlice.Contains(packSizes, sumResults) {
		res = []int{sumResults}
		return res
	}

	//If the closest value between the sum of the results and the order number is in the packSizes, return it
	closest := tSlice.FindClosest(a.Item, packSizes)
	if closest < sumResults && closest > a.Item {
		res = []int{closest}
		return res
	}

	return res
}
