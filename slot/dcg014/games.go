package dcg014

import (
	"github.com/death12358/digitalopen/games/slots"
	"github.com/death12358/digitalopen/slot/dcg014/easteregg"
	"github.com/death12358/digitalopen/slot/dcg014/grow"
)

// dcg014_DEF - dcg014 game definition
// 需要增加adventure
// 優化以下內容
type Games struct {
	ng_game *slots.WayGames
	fg_game *grow.Games
	bg_game *easteregg.Games
}

// NewSGames - create new dcg014
func NewGames(ng *slots.WayGames, fg *grow.Games, bg *easteregg.Games) *Games {
	return &Games{
		ng_game: ng,
		fg_game: fg,
		bg_game: bg,
	}
}

var (
	// unit bet
	// Is use the  sync.map  function optimize?
	dcg014_gameplay map[string]*Games = map[string]*Games{
		"100": NewGames(
			dcg014ng,
			dcg014fg,
			dcg014bg,
		),
	}
)
