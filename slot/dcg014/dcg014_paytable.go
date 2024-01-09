package dcg014

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

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
	//SymbolNA - Scatter NA Bonus
	NA = games.Symbol(-1)

	SYMBOL_COUNT = 16
)

var (
	SymbolList      = []games.Symbol{WW, H1, H2, H3, H4, H5, LA, LK, LQ, LJ, LT, LN, LE, SF, SB, NA}
	ScatterPosition = []games.Symbol{SF, SB, NA}
)

// 赔率表录入,注意对应关系
var (
	//yeah
	ngpayTable_unitbet100 = &games.PayTable{
		// WW
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// H1
		games.Pays{decimal.Zero, decimal.NewFromInt(10), decimal.NewFromInt(40), decimal.NewFromInt(80), decimal.NewFromInt(200), decimal.NewFromInt(1000)},
		// H2
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(30), decimal.NewFromInt(60), decimal.NewFromInt(120), decimal.NewFromInt(240)},
		// H3
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(50), decimal.NewFromInt(100), decimal.NewFromInt(200)},
		// H4
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(40), decimal.NewFromInt(100)},
		// H5
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// LA
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(50)},
		// LK
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(40)},
		// LQ
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
		// LJ
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
		// LT
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10)},
		// LN
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10)},
		// SE
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// SF
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// SB
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	}

	fgpayTable_unitbet100 = &games.PayTable{
		// WW
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// H1
		games.Pays{decimal.Zero, decimal.NewFromInt(10), decimal.NewFromInt(40), decimal.NewFromInt(80), decimal.NewFromInt(200), decimal.NewFromInt(1000)},
		// H2
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(30), decimal.NewFromInt(60), decimal.NewFromInt(120), decimal.NewFromInt(240)},
		// H3
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(50), decimal.NewFromInt(100), decimal.NewFromInt(200)},
		// H4
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(40), decimal.NewFromInt(100)},
		// H5
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// LA
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(50)},
		// LK
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20), decimal.NewFromInt(40)},
		// LQ
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
		// LJ
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
		// LT
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10)},
		// LN
		games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(5), decimal.NewFromInt(10)},
		// SE
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// SF
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
		// SB
		games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	}
)
