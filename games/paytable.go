package games

import "github.com/shopspring/decimal"

// PayTable - 賠付表結構
type PayTable []Pays

// CalcPayTable - 計算獎金
func (p PayTable) CalcPayTable(symbol int, count int) decimal.Decimal {
	return p[symbol][count]
}

// CalcPaysTable - 計算獎金 (多組)
func (p PayTable) CalcPaysTable(symbol int, count, muti int) decimal.Decimal {
	return p[symbol][count].Mul(decimal.NewFromInt(int64(muti)))
}

// Pays - 賠付結構
type Pays []decimal.Decimal

// CalcPays - 計算獎金
func (p Pays) CalcPays(count int) decimal.Decimal {
	return p[count]
}
