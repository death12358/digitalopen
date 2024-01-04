package lightning

import (
	"games"

	"github.com/shopspring/decimal"

	weights "games/weight"
)

const (

	// Lightning - 獎圖定義
	bg_default = games.Symbol(0) + iota
	bg_fg01
	bg_fg05
	bg_fg1
	bg_fg2
	bg_fg5
	bg_fg10
	bg_bg
	bg_fg
)

var (
	bgLightningSymbolList = []games.Symbol{
		bg_default,
		bg_fg01,
		bg_fg05,
		bg_fg1,
		bg_fg2,
		bg_fg5,
		bg_fg10,
		bg_bg}

	bgSymbolName = []string{
		"0",
		"fg0.1",
		"fg0.5",
		"fg1",
		"fg2",
		"fg5",
		"fg10",
		"bg",
		"fg",
	}

	bgSymbolMapping = map[string]games.Symbol{
		"fg0.1": bg_fg01,
		"fg0.5": bg_fg05,
		"fg1":   bg_fg1,
		"fg2":   bg_fg2,
		"fg5":   bg_fg5,
		"fg10":  bg_fg10,
	}

	bgLightningPayTable = games.Pays{
		decimal.NewFromFloat(0.0),
		decimal.NewFromFloat(0.1),
		decimal.NewFromFloat(0.5),
		decimal.NewFromFloat(1.0),
		decimal.NewFromFloat(2.0),
		decimal.NewFromFloat(5.0),
		decimal.NewFromFloat(10.0),
		decimal.NewFromFloat(0),
		decimal.NewFromFloat(0),
	}

	//8
	bg_LightningWeightTable8 = []int{900, 90, 10}
	bg_LightningObjectTable8 = []int{int(bg_default), int(bg_fg), int(bg_bg)}
	LightningGame_8          = weights.NewGames(
		bg_LightningWeightTable8,
		bg_LightningObjectTable8,
	)

	bg_LightningFGWeightTable8 = []int{1, 7, 35, 15, 5, 1}
	bg_LightningFGObjectTable8 = []int{int(bg_fg01), int(bg_fg05), int(bg_fg1), int(bg_fg2), int(bg_fg5), int(bg_fg10)}
	LightningFreeGame_8        = weights.NewGames(
		bg_LightningFGWeightTable8,
		bg_LightningFGObjectTable8,
	)

	//18
	bg_LightningWeightTable18 = []int{880, 100, 20}
	bg_LightningObjectTable18 = []int{int(bg_default), int(bg_fg), int(bg_bg)}
	LightningGame_18          = weights.NewGames(
		bg_LightningWeightTable18,
		bg_LightningObjectTable18,
	)

	bg_LightningFGWeightTable18 = []int{1, 24, 30, 10, 5, 1}
	bg_LightningFGObjectTable18 = []int{int(bg_fg01), int(bg_fg05), int(bg_fg1), int(bg_fg2), int(bg_fg5), int(bg_fg10)}
	LightningFreeGame_18        = weights.NewGames(
		bg_LightningFGWeightTable18,
		bg_LightningFGObjectTable18,
	)

	//38
	bg_LightningWeightTable38 = []int{860, 110, 30}
	bg_LightningObjectTable38 = []int{int(bg_default), int(bg_fg), int(bg_bg)}
	LightningGame_38          = weights.NewGames(
		bg_LightningWeightTable38,
		bg_LightningObjectTable38,
	)

	bg_LightningFGWeightTable38 = []int{5, 24, 35, 10, 5, 1}
	bg_LightningFGObjectTable38 = []int{int(bg_fg01), int(bg_fg05), int(bg_fg1), int(bg_fg2), int(bg_fg5), int(bg_fg10)}
	LightningFreeGame_38        = weights.NewGames(
		bg_LightningFGWeightTable38,
		bg_LightningFGObjectTable38,
	)
	//68
	bg_LightningWeightTable68 = []int{830, 130, 40}
	bg_LightningObjectTable68 = []int{int(bg_default), int(bg_fg), int(bg_bg)}
	LightningGame_68          = weights.NewGames(
		bg_LightningWeightTable68,
		bg_LightningObjectTable68,
	)

	bg_LightningFGWeightTable68 = []int{5, 23, 35, 10, 5, 1}
	bg_LightningFGObjectTable68 = []int{int(bg_fg01), int(bg_fg05), int(bg_fg1), int(bg_fg2), int(bg_fg5), int(bg_fg10)}
	LightningFreeGame_68        = weights.NewGames(
		bg_LightningFGWeightTable68,
		bg_LightningFGObjectTable68,
	)
	//88
	bg_LightningWeightTable88 = []int{800, 150, 50}
	bg_LightningObjectTable88 = []int{int(bg_default), int(bg_fg), int(bg_bg)}
	LightningGame_88          = weights.NewGames(
		bg_LightningWeightTable88,
		bg_LightningObjectTable88,
	)

	bg_LightningFGWeightTable88 = []int{5, 33, 37, 10, 4, 2}
	bg_LightningFGObjectTable88 = []int{int(bg_fg01), int(bg_fg05), int(bg_fg1), int(bg_fg2), int(bg_fg5), int(bg_fg10)}
	LightningFreeGame_88        = weights.NewGames(
		bg_LightningFGWeightTable88,
		bg_LightningFGObjectTable88,
	)
)

var (
	real_reels_pos = []bool{
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
		false, true, true, true, false,
	}

	real_reels_pos_InvertRegularXYAxis = []bool{
		false, false, false, false, false,
		true, true, true, true, true,
		true, true, true, true, true,
		true, true, true, true, true,
		false, false, false, false, false,
	}
)

func BGReel(r games.Reels) []string {
	var reel []string
	for _, v := range r {
		reel = append(reel, bgSymbolName[v])
	}
	return reel
}

func BGName(s games.Symbol) string {
	return bgSymbolName[s]
}
