package dcg008

import (
	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/slot/dcg008/adventure"
)

// dcg008_DEF - dcg008 game definition
// 需要增加adventure
// 優化以下內容
type Games struct {
	ng_game *slots.WayGames
	fg_game *slots.WayGames
	bg_game *adventure.Games
}

// NewSGames - create new dcg008
func NewGames(ng *slots.WayGames, fg *slots.WayGames, bg *adventure.Games) *Games {
	return &Games{
		ng_game: ng,
		fg_game: fg,
		bg_game: bg,
	}
}

var (

	// unit bet
	// 是否使用 sync.map 優化
	dcg008_gameplay map[string]*Games = map[string]*Games{
		"100": NewGames(
			dcg008ng,
			dcg008fg,
			dcg008bg,
		),
	}
)
