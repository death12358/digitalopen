package sg001

import (
	"strconv"

	"github.com/death12358/digitalopn/games"
	"github.com/death12358/digitalopn/games/random"
	"github.com/death12358/digitalopn/games/slots"

	"github.com/shopspring/decimal"
)

// IsBigFGWin - check if Big Symbol Free Game is win
func (s *SG001) IsBigFGWin(unitbet string, reel *games.ReelStrips) bool {
	bigfg_game := sg001_gameplay[unitbet].bigfg_game

	sc := bigfg_game.CalcScatter(*reel, decimal.Zero, slots.WW)

	if sc.Wins[0].Match >= 3 {
		return true
	}
	return false
}

// BigFGFlow - Big Symbol Free Game flow
func (s *SG001) BigFGFlow(rtp string, bet decimal.Decimal, round *games.Rounds) *games.Rounds {
	count := 6
	stage := int64(count) // 5 4 3 2 1 0
	round.Status = round.Status.Push(games.FreeGame)
	round.Position = round.Status.Push(games.FreeGame)
	unitbet := round.Result["0"].Pickem[0]
	for i := 0; i < count; i++ {
		fg_reel, show := s.SpinBigFGReelStripts(unitbet, rtp)
		round.Stages++
		stage--
		id := strconv.Itoa(int(round.Stages))
		res := s.CalcBigFGResult(unitbet, fg_reel, id, round.Brand, round.Username, show, bet)
		res.Stages = stage
		res.Case = games.NotStartedYet
		round.Result[id] = res
		round.TotalPoint = round.TotalPoint.Add(res.Point)
		/*if s.IsBigFGWin(unitbet, fg_reel) {
			count += 6
			stage += 6
		}*/
	}

	// round.Stages = 1
	// round.Result["1"].Case = games.Lose
	// if !round.Result["1"].Point.IsZero() {
	// 	round.Result["1"].Case = games.Win
	// }
	return round
}

// SpinBigFGReelStripts - spin Big Symbol Free Game reel strips
func (s *SG001) SpinBigFGReelStripts(unitbet, rtp string) (*games.ReelStrips, []string) {
	bigfg_game := sg001_gameplay[unitbet].bigfg_game
	rtps := games.RTPs(rtp)
	bigfglen := bigfg_game.GetReelsLen(rtps)
	pos := random.Intsn(bigfglen)
	pos[2] = pos[1]
	pos[3] = pos[1]

	bigReel := []bool{false, true, true, true, false}

	reelStrips := bigfg_game.BigSymbolsReelStrips(rtps, pos, []bool{false, true, true, true, false})
	// show_reelStrips := ng_game.SpinNormal(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	show_reelStrips := bigfg_game.ShowBigSymbolReelStrips(rtps, pos, bigReel, 1, 1).InvertRegularXYAxis().Strings()
	return &reelStrips, show_reelStrips
}

// CalcBigFGResult - calc Big Symbol Free Game  result
func (s *SG001) CalcBigFGResult(unitbet string, reel *games.ReelStrips, id, brand, user string, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.NotStartedYet,
		Pickem:   []string{""},
		Symbols:  symbols,
	}
	bigfg_game := sg001_gameplay[unitbet].bigfg_game

	hifg_result := bigfg_game.CalcWayReel(*reel, bet)
	if hifg_result.IsWinGreaterThanZero() {
		record.Case = games.NotStartedYet
		record.Multiplier = hifg_result.TotalWin.Div(bet)
		record.Point = hifg_result.TotalWin
	}

	return &record
}
