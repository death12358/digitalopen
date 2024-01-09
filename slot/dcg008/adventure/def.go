package adventure

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

const (
	bgAdv_default = games.Symbol(0) + iota
	bgAdv_crys50
	bgAdv_crys100
	bgAdv_crys200
	bgAdv_crys300
	bgAdv_crys500
)

var (
	unitbet               = decimal.NewFromInt(100)
	bgAdventureSymbolList = []games.Symbol{
		bgAdv_default,
		bgAdv_crys50,
		bgAdv_crys100,
		bgAdv_crys200,
		bgAdv_crys300,
		bgAdv_crys500,
	}
	// bgAdventureSymbolName = []string{
	// 	"0",
	// 	"crystal50",
	// 	"crystal100",
	// 	"crystal200",
	// 	"crystal300",
	// 	"crystal500"}
	// bgAdventureSymbolMapping = map[string]games.Symbol{
	// 	"0":          bgAdv_default,
	// 	"crystal50":  bgAdv_crys50,
	// 	"crystal100": bgAdv_crys100,
	// 	"crystal200": bgAdv_crys200,
	// 	"crystal300": bgAdv_crys300,
	// 	"crystal500": bgAdv_crys500,
	// }

	AdventurePayTable = &games.Pays{
		decimal.NewFromInt(0),
		decimal.NewFromInt(50),
		decimal.NewFromInt(100),
		decimal.NewFromInt(200),
		decimal.NewFromInt(300),
		decimal.NewFromInt(500),
	}
	//獎項分配
	AdvPWTable = [][]int{
		{0, 30, 25, 15, 10, 20},
		{5, 25, 15, 10, 15, 30},
		{15, 20, 15, 10, 20, 20},
		{20, 10, 10, 10, 25, 25},
		{20, 5, 5, 5, 5, 60},
		{20, 5, 5, 5, 5, 60},
	}
	BgAdvPOTable = []int{int(bgAdv_default), int(bgAdv_crys50), int(bgAdv_crys100), int(bgAdv_crys200), int(bgAdv_crys300), int(bgAdv_crys500)}

	// AdventurFl1PointGame = weights.NewGames(
	// 	bgAdvPointWeightTable[0],
	// 	bgAdvPointObjectTable)

)

func Symbol(i int) games.Symbol {
	return bgAdventureSymbolList[i]
}
