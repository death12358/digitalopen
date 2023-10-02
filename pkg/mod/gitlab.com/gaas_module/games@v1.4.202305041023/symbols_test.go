package games_test

import (
	"log"
	"testing"

	"github.com/shopspring/decimal"
	"gitlab.com/gaas_module/games"
	"gitlab.com/gaas_module/games/slots"
)

func TestSymbols(t *testing.T) {
	reels := games.ReelStrips{
		{5, 3, 4},
		{0, 2, 2},
		{4, 1, 5},
		{7, 0, 5},
		{10, 8, 5},
	}

	reels = games.ReelStrips{
		{3, 5, 4},
		{3, 1, 5},
		{4, 5, 0},
		{5, 5, 3},
		{4, 4, 1},
	}
	bet := decimal.NewFromInt(1)
	result := slots.NewWinDetail(bet)
	symbols := reels[0].RemoveDuplicates()
	// log.Printf("symbols: %+v", symbols)

	for _, symbol := range symbols {
		count, way := reels.CalcSymbolsMatchFromLeft(symbol, 0)
		match := len(count) - 1
		log.Printf("symbol: %d, count: %d, way: %d, match: %d", symbol, count, way, match+1)
		pay := paytable.CalcPaysTable(symbol.Int(), match, way).Mul(bet)
		// log.Printf("pay: %s", pay.String())
		result.Add(pay)
		result.Wins = append(result.Wins, slots.NewWayWins(symbol, match, way, pay))
		log.Printf("result: %+v", result)
	}

}

var paytable = &games.PayTable{
	// WW
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// H1
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H2
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H3
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H4
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(10), decimal.NewFromInt(20)},
	// H5
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LA
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LK
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LQ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LJ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LT
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// LN
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(5), decimal.NewFromInt(8), decimal.NewFromInt(10)},
	// SE
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// SF
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// SB
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
}
