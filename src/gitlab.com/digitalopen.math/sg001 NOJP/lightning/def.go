package lightning

import (
	"github.com/shopspring/decimal"
	"gitlab.com/gaas_module/games"

	weights "gitlab.com/gaas_module/games/weight"
)

const (

	// Lightning - 獎圖定義
	bg_default = games.Symbol(0) + iota
	bg_fg8
	bg_fg18
	bg_fg38
	bg_fg68
	bg_fg88
	bg_fg888
	bg_bg
	bg_fg
)

var (
	bgLightningSymbolList = []games.Symbol{
		bg_default,
		bg_fg8,
		bg_fg18,
		bg_fg38,
		bg_fg68,
		bg_fg88,
		bg_fg888,
		bg_bg}

	bgSymbolName = []string{
		"0",
		"fg8",
		"fg18",
		"fg38",
		"fg68",
		"fg88",
		"fg888",
		"bg",
		"fg",
	}

	bgSymbolMapping = map[string]games.Symbol{
		"fg8":   bg_fg8,
		"fg18":  bg_fg18,
		"fg38":  bg_fg38,
		"fg68":  bg_fg68,
		"fg88":  bg_fg88,
		"fg888": bg_fg888,
	}

	// RTP 修正數
	rtp_fix = decimal.NewFromFloat(0.01163)

	bgLightningPayTable = games.Pays{
		decimal.NewFromInt(0),
		decimal.NewFromInt(8).Mul(rtp_fix),
		decimal.NewFromInt(18).Mul(rtp_fix),
		decimal.NewFromInt(38).Mul(rtp_fix),
		decimal.NewFromInt(68).Mul(rtp_fix),
		decimal.NewFromInt(88).Mul(rtp_fix),
		decimal.NewFromInt(888).Mul(rtp_fix),
		decimal.NewFromInt(0),
		decimal.NewFromInt(0),
	}

	bg_LightningWeightTable = []int{90, 8, 2}
	bg_LightningObjectTable = []int{int(bg_default), int(bg_fg), int(bg_bg)}

	LightningGame = weights.NewGames(
		bg_LightningWeightTable,
		bg_LightningObjectTable,
	)

	bg_LightningFGWeightTable = []int{100, 200, 1200, 600, 200, 10}
	bg_LightningFGObjectTable = []int{int(bg_fg8), int(bg_fg18), int(bg_fg38), int(bg_fg68), int(bg_fg88), int(bg_fg888)}

	LightningFreeGame = weights.NewGames(
		bg_LightningFGWeightTable,
		bg_LightningFGObjectTable,
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
