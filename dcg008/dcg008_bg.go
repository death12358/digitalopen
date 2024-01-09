package dcg008

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
)

const (
	// bg game define
	//
	bg_SBmatch_def = 3
	bg_SFmatch_def = 3
)

var bigReel = []bool{false, true, true, true, false}

// Lightning
// ------------------------------------------------------------

// IsBonusWin - check Maya adventure bonus game can be trigger.
func (s *DCG008) IsBonusWin(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.SB)
	return count >= bg_SBmatch_def
}
