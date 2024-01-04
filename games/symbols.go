package games

import "strconv"

// TODO: 撰寫測試
// Symbol - 獎圖結構
type Symbol int

// Int - 轉換為整數
func (s Symbol) Int() int {
	return int(s)
}

// String - 轉換為字串
func (s Symbol) String() string {
	return strconv.Itoa(int(s))
}

// ToSymbol - 轉換為符號
func ToSymbol(s string) Symbol {
	i, _ := strconv.Atoi(s)
	return Symbol(i)
}

// ToSymbolArray - 轉換為符號陣列
func ToSymbolArray(sa []string) []Symbol {
	var result []Symbol
	for _, s := range sa {
		result = append(result, ToSymbol(s))
	}
	return result
}

// // ToSymbol - 轉換為符號
// func ToSymbol(s string) Symbol {
// 	i, _ := strconv.Atoi(s)
// 	return Symbol(i)
// }

// ToSymbols - 轉換為符號陣列
func ToSymbols(sa []string) []Symbol {
	var result []Symbol
	for _, s := range sa {
		result = append(result, ToSymbol(s))
	}
	return result
}
