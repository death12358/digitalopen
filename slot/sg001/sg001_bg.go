package sg001

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
)

const (
	// bg game define
	bg_hifg_def      = "0"
	bg_lightning_def = "1"

	// 多少 SF 才會觸發 Lightning
	// test
	//bg_match_def = 3
	bg_match_def = 6
)

// Lightning
// ------------------------------------------------------------

// IsBigFGWin - check if Big Symbol Free Game is win
func (s *SG001) IsBonusWin(unitbet string, reel *games.ReelStrips) bool {
	// hifg_game := sg001_gameplay[unitbet].hifg_game
	// sc := hifg_game.CalcScatter(*reel, decimal.Zero, slots.SF)

	// if sc.Wins[0].Ways >= bg_match_def {
	// 	return true
	// }

	_, count := reel.CalcSymbolMatches(slots.SF)

	if count >= bg_match_def {
		return true
	}
	return false
}
