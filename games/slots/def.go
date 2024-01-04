package slots

import "digitalopen/games"

// ReelStripsDef - 轉輪表定義，用於盤面個數
type ReelStripsDef []int

// kWinLines - 幾條線
//var kWinLines int

// PayLine - 線圖結構
//type PayLine [][]int

// 優化以下程式可讀性
// General Symbols Definition - 常用獎圖定義
const (
	// SymbolWild - Scatter Wild
	WW = games.Symbol(0)
	// SymbolHighPay1 - High Pay 1
	H1 = games.Symbol(1)
	// SymbolHighPay2 - High Pay 2
	H2 = games.Symbol(2)
	// SymbolHighPay3 - High Pay 3
	H3 = games.Symbol(3)
	// SymbolHighPay4 - High Pay 4
	H4 = games.Symbol(4)
	// SymbolHighPay5 - High Pay 5
	H5 = games.Symbol(5)
	// SymbolLowPayA - Low Pay A
	LA = games.Symbol(6)
	// SymbolLowPayK - Low Pay K
	LK = games.Symbol(7)
	// SymbolLowPayQ - Low Pay Q
	LQ = games.Symbol(8)
	// SymbolLowPayJ - Low Pay J
	LJ = games.Symbol(9)
	// SymbolLowPayT - Low Pay Ten
	LT = games.Symbol(10)
	// SymbolLowPayN - Low Pay Nine
	LN = games.Symbol(11)
	// SymbolLowPayE - Low Pay Eight
	LE = games.Symbol(12)
	// SymbolFreeSpin - Scatter Free Spin
	SF = games.Symbol(13)
	// SymbolBonus - Scatter Bonus
	SB = games.Symbol(14)

	SYMBOL_COUNT = 15
)

var (
	SymbolList      = []games.Symbol{WW, H1, H2, H3, H4, H5, LA, LK, LQ, LJ, LT, LN, LE, SF, SB}
	ScatterPosition = []games.Symbol{SF, SB}
)
