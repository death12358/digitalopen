package sg001

import (
	"strconv"

	"digitalopen/games"
	"digitalopen/games/random"

	"github.com/shopspring/decimal"
)

// games.BonusFreeGame1

// HiFGFlow - Bonus High pay Symbol Free Game flow
func (s *SG001) HiFGFlow(rtp string, bet decimal.Decimal, round *games.Rounds) (*games.Rounds, error) {
	round.Status = round.Status.Pop(games.Bonus)
	round.Status = round.Status.Push(games.BonusGame1)

	round.Position = games.BonusGame1
	unitbet := round.Result["0"].Pickem[0]

	// default 8 times
	count := 8
	stage := int64(count) // 7 6 5 4 3 2 1 0

	for i := 0; i < count; i++ {
		fg_reel, show := s.SpinHIFGReelStripts(unitbet, rtp)
		round.Stages++
		stage--
		id := strconv.Itoa(int(round.Stages))
		res := s.CalcHIFGResult(unitbet, fg_reel, id, round.Brand, round.Username, show, bet)
		res.Stages = stage
		res.Case = games.NotStartedYet
		round.Result[id] = res
		round.TotalPoint = round.TotalPoint.Add(res.Point)
	}

	// round.Stages = 1
	round.Result["1"].Case = games.Lose
	if !round.Result["1"].Point.IsZero() {
		round.Result["1"].Case = games.Win
	}

	return round, nil
}

// SpinHIFGReelStripts - spin HighPay Free Game reel strips
func (s *SG001) SpinHIFGReelStripts(unitbet, rtp string) (*games.ReelStrips, []string) {
	hifg_game := sg001_gameplay[unitbet].hifg_game
	rtps := games.RTPs(rtp)
	nglen := hifg_game.GetReelsLen(rtps)
	pos := random.Intsn(nglen)

	reelStrips := hifg_game.ContiguousReelStrips(rtps, pos)
	// show_reelStrips := ng_game.SpinNormal(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	show_reelStrips := hifg_game.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()

	return &reelStrips, show_reelStrips
}

// CalcHIFGResult - calc HighPay Free Game  result
func (s *SG001) CalcHIFGResult(unitbet string, reel *games.ReelStrips, id, brand, user string, symbols []string, bet decimal.Decimal) *games.Records {
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.NotStartedYet,
		Pickem:   []string{""},
		Symbols:  symbols,
	}
	high_game := sg001_gameplay[unitbet].hifg_game

	hifg_result := high_game.CalcWayReel(*reel, bet)
	if hifg_result.IsWinGreaterThanZero() {
		record.Case = games.NotStartedYet
		record.Multiplier = hifg_result.TotalWin.Div(bet)
		record.Point = hifg_result.TotalWin
	}

	return &record
}
