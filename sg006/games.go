package sg006

import (
	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/sg006/vs"
)

// SG006_DEF - sg006 game definition

// 優化以下內容
type Games struct {
	ng_game *slots.WayGames
	fg_game *slots.WayGames
	vs_game map[string]*vs.Games
}

// NewSGames - create new SG006_DEF
func NewGames(ng, fg *slots.WayGames, vs map[string]*vs.Games) *Games {
	return &Games{
		ng_game: ng,
		fg_game: fg,
		vs_game: vs,
	}
}

var (

	// unit bet
	// 是否使用 sync.map 優化
	sg006_gameplay map[string]*Games = map[string]*Games{
		"100": NewGames(
			sg006ng_100,
			sg006fg_100,
			sg006vs,
		),
	}
)
