package dcg014

import (
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/shopspring/decimal"
)

const bg_SFmatch_def = 3

func (s *DCG014) IsFGWin(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.SF)
	return count >= bg_SFmatch_def
}

// NumbersofFGWin - Get the number of times of Big Symbol Free Game trigger.
func (s *DCG014) NumbersofFGWin(unitbet string, reel *games.ReelStrips) int {
	var number int
	_, count := reel.CalcSymbolMatches(slots.SF)
	if count == 3 {
		number = 5
	}
	if count == 4 {
		number = 6
	}
	if count == 5 {
		number = 7
	}
	if count >= 6 {
		number = 10
	}
	return number
}

// FGFlow - Free Game flow
//
//	@param count: initial freespin count.
func (s *DCG014) FreeGameFlow(bet decimal.Decimal, unitbet, rtp string, round *games.Rounds, count int) (games.Rounds, error) {
	level := 1
	//rtps := games.RTPs(rtp)
	fg_game := dcg014_gameplay[unitbet].fg_game
	// for k := range round.Result {
	// 	delete(round.Result, k)
	// }

	round.Status = round.Status.Push(games.FreeGame)
	//fmt.Print(round.Status)
	round.Position = round.Status.Push(games.FreeGame)
	round.TotalBet = bet
	// in branch of NG.
	Count := count

	for i := 1; i <= Count; i++ {

		round.Stages = int64(Count - i)
		fglen := fg_game.GetReelsLen(level, unitbet, rtp)
		//RANDOM
		position := random.Intsn(fglen)
		reelStrips, initialposition, show_reelStrips, _ := fg_game.FirstSpinReelStripts(level, unitbet, rtp, position)
		// round.Stages++
		id := strconv.Itoa(i)
		fg_result, _, _, _, nextLevel, respin := fg_game.CalcFGResult(level, unitbet, rtp, *reelStrips, initialposition, id, round.Brand, round.Username, show_reelStrips, bet)
		if nextLevel > level {
			Count = Count + 1
		}
		Count = Count + respin
		respin = 0
		//fmt.Println("respin: ", Count)

		level = nextLevel
		fg_result.Stages = round.Stages
		round.Result[id] = fg_result
		round.TotalPoint = round.TotalPoint.Add(fg_result.Point)
		fg_result.Case = games.NotStartedYet

	}
	return *round, nil
}
