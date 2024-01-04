package sg001

import (
	"digitalopen/games/slots"

	"digitalopen/sg001/jackpot"
	"digitalopen/sg001/lightning"
)

// SG001_DEF - sg001 game definition

// 優化以下內容
type Games struct {
	ng_game          *slots.WayGames
	hifg_game        *slots.WayGames
	bigfg_game       *slots.WayGames
	bgLightning_game *lightning.Games
	jp_game          *jackpot.Games
}

// NewSGames - create new SG001_DEF
func NewGames(ng, hifg, bigfg *slots.WayGames, lightning *lightning.Games, jackpot *jackpot.Games) *Games {
	return &Games{
		ng_game:          ng,
		hifg_game:        hifg,
		bigfg_game:       bigfg,
		bgLightning_game: lightning,
		jp_game:          jackpot,
	}
}

var (

	// unit bet
	// 是否使用 sync.map 優化
	sg001_gameplay map[string]*Games = map[string]*Games{
		"8": NewGames(
			sg001ng_8,
			sg001hifg_8,
			sg001bigfg_8,
			sg001lightning_8,
			sg001jackpot,
		),
		"18": NewGames(
			sg001ng_18,
			sg001hifg_18,
			sg001bigfg_18,
			sg001lightning_18,
			sg001jackpot,
		),
		"38": NewGames(
			sg001ng_38,
			sg001hifg_38,
			sg001bigfg_38,
			sg001lightning_38,
			sg001jackpot,
		),
		"68": NewGames(
			sg001ng_68,
			sg001hifg_68,
			sg001bigfg_68,
			sg001lightning_68,
			sg001jackpot,
		),
		"88": NewGames(
			sg001ng_88,
			sg001hifg_88,
			sg001bigfg_88,
			sg001lightning_88,
			sg001jackpot,
		),
	}
)
