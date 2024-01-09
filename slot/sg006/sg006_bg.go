package sg006

import (
	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/slots"
)

const (
	// bg game define

	// 多少 SF 才會觸發 Lightning
	//bg_SFmatch_def = 3 //Test
	bg_SFmatch_def = 4
	VS_H1match_def = 1
)

// IsFGWin - check if Free Game is win
func (s *SG006) IsFGWin(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.SF)

	if count >= bg_SFmatch_def {
		return true
	}
	return false
}

func (s *SG006) IsVSWin(unitbet string, reel *games.ReelStrips, VSReel [5]int) (bool, int) {
	_, count := reel.CalcSymbolMatches(slots.H1)
	VSNumber := 0
	for _, v := range VSReel {
		VSNumber += v
	}
	if count >= VS_H1match_def && VSNumber > 0 {
		return true, VSNumber
	}
	return false, VSNumber
}
