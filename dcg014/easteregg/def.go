package easteregg

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

const (
	bgegg_default = games.Symbol(0) + iota
	bgegg_200
	bgMult_150
	bgMult_200
	bgMult_300
	bgMult_500
	bgAdd_5
	bgAdd_10
	bgAdd_15
)

var (
	unitbet      = decimal.NewFromInt(100)
	BgSymbolList = []games.Symbol{
		bgegg_default,
		bgegg_200,
		bgMult_150,
		bgMult_200,
		bgMult_300,
		bgMult_500,
		bgAdd_5,
		bgAdd_10,
		bgAdd_15,
	}

	PayTable = &games.Pays{
		decimal.NewFromInt(0),
		decimal.NewFromInt(200),
		decimal.NewFromFloat(1.5),
		decimal.NewFromFloat(2.0),
		decimal.NewFromFloat(3.0),
		decimal.NewFromFloat(5.0),
		decimal.NewFromInt(500),
		decimal.NewFromInt(1000),
		decimal.NewFromInt(1500),
	}
	//獎項分配
	Bg_EggWeightTable = [][]int{
		{0, 100},
		{0, 100},
		{10, 90},
		{40, 60},
		{60, 40},
	}
	Bg_EggNameTable = []int{int(bgegg_default), int(bgegg_200)}

	Bg_BonusWeightTable = [][]int{
		//400
		{0, 0, 0, 50, 0, 0, 50},
		//600
		{2, 8, 12, 28, 12, 30, 8},
		//800
		{15, 20, 5, 10, 5, 34, 11},
		//1000
		{16, 21, 10, 3, 40, 2, 8},
	}
	Bg_BonusNameTable = []int{
		int(bgMult_150),
		int(bgMult_200),
		int(bgMult_300),
		int(bgMult_500),
		int(bgAdd_5),
		int(bgAdd_10),
		int(bgAdd_15),
	}
)

func Symbol(i int) games.Symbol {
	return BgSymbolList[i]
}
