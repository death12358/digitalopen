package tools

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// ConvertToDecimal 轉換為decimal(support int,int32,int64,float32,float64,string)
func ConvertToDecimal(value interface{}) (decimal.Decimal, error) {
	switch v := value.(type) {
	case int:
		return decimal.NewFromInt(int64(v)), nil
	case int32:
		return decimal.NewFromInt32(v), nil
	case int64:
		return decimal.NewFromInt(v), nil
	case float32:
		return decimal.NewFromFloat32(v), nil
	case float64:
		return decimal.NewFromFloat(v), nil
	case string:
		return decimal.NewFromString(v)
	default:
		return decimal.Decimal{}, fmt.Errorf("unsupported type: %T", value)
	}
}

// CalcDecimalDiv v1 / v2
func CalcDecimalDiv(v1, v2 interface{}) (decimal.Decimal, error) {
	v1D, err := ConvertToDecimal(v1)
	if err != nil {
		return decimal.Zero, err
	}
	v2D, err := ConvertToDecimal(v2)
	if err != nil {
		return decimal.Zero, err
	}

	if v2D.IsZero() {
		return decimal.Zero, err
	}

	return v1D.Div(v2D), nil
}

// CalcDecimalMul v1 * v2
func CalcDecimalMul(v1, v2 interface{}) (decimal.Decimal, error) {
	v1D, err := ConvertToDecimal(v1)
	if err != nil {
		return decimal.Zero, err
	}
	v2D, err := ConvertToDecimal(v2)
	if err != nil {
		return decimal.Zero, err
	}
	return v1D.Mul(v2D), nil
}
