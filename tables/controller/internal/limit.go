package internal

import (
	"math"
)

const (
	limitValIndex = 0
)

type ProbabilityLimit struct {
	Upper int
	Lower int
}

// // CreateProbabilityLimit 建立機率限制
// func CreateProbabilityLimit(inData DataMap) ProbabilityLimit {
// 	pL := ProbabilityLimit{
// 		Upper: math.MaxInt,
// 		Lower: -1,
// 	}

// 	rT := reflect.TypeOf(pL)

// 	numFields := rT.NumField()
// 	// 透過迴圈獲取每個欄位的名稱
// 	for i := 0; i < numFields; i++ {
// 		field := rT.Field(i)
// 		key := strings.ToUpper(field.Name)
// 		if data, ok := inData.ParseRowDataToInt(key); ok && len(data) > 0 {
// 			limitVal := int64(data[limitValIndex])
// 			if limitVal != -1 {
// 				rV := reflect.ValueOf(&pL).Elem()
// 				fieldV := rV.FieldByName(field.Name)
// 				fieldV.SetInt(int64(data[limitValIndex]))
// 			}
// 		}
// 	}

// 	return pL
// }

func (pl *ProbabilityLimit) IsOutOfLimit(inVal int) bool {
	if pl.Upper <= pl.Lower {
		return false
	}
	return inVal < pl.Lower || inVal > pl.Upper
}

func (pl ProbabilityLimit) CalculateMergedProbabilityLimit(upper int) *ProbabilityLimit {
	newProbabilityLimit := ProbabilityLimit{
		Upper: pl.Upper,
		Lower: pl.Lower,
	}

	if upper == -1 {
		upper = math.MaxInt
	}
	if upper < pl.Upper {
		newProbabilityLimit.Upper = upper
	}

	return &newProbabilityLimit
}
