package dcg008

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/games/slots"
	"github.com/shopspring/decimal"
)

// IsBigFGWin - check Big Symbol Free Game can be trigger.
func (s *DCG008) IsBigFGWin(unitbet string, reel *games.ReelStrips) bool {
	_, count := reel.CalcSymbolMatches(slots.SF)
	return count >= bg_SFmatch_def
}

// NumbersofFGWin - obtain the number of times of Big Symbol Free Game trigger.
func (s *DCG008) NumbersofFGWin(unitbet string, reel *games.ReelStrips) int {
	var number int
	_, count := reel.CalcSymbolMatches(slots.SF)
	if count == 3 {
		number = 7
	}
	if count == 4 {
		number = 10
	}
	if count >= 5 {
		number = 15
	}
	return number
}

// BigFGFlow - Big Symbol Free Game flow
func (s *DCG008) BigFreeGameFlow(bet decimal.Decimal, unitbet, rtp string, round *games.Rounds, count int) (games.Rounds, error) {

	//Stage := int64(count) // 6,5,4,3,2,1,0

	rtps := games.RTPs(rtp)
	bigfg_game := dcg008_gameplay[unitbet].fg_game
	// if round == nil {
	// 	return nil, errors.New("round is nil")
	// }
	// for k := range round.Result {
	// 	delete(round.Result, k)
	// }
	bigfglen := bigfg_game.GetReelsLen(rtps)
	round.Status = round.Status.Push(games.FreeGame)
	//fmt.Print(round.Status)
	round.Position = round.Status.Push(games.FreeGame)
	round.TotalBet = bet
	// in branch of NG.

	for i := 1; i <= count; i++ {
		round.Stages = int64(count - i)
		position := random.Intsn(bigfglen)
		reelStrips, initialposition, show_reelStrips, _ := s.FirstSpinFGReelStripts(unitbet, rtps, position)
		// round.Stages++

		id := strconv.Itoa(i)
		fg_result, _, _ := s.CalcBigFGResult(unitbet, rtp, *reelStrips, initialposition, id, round.Brand, round.Username, show_reelStrips, bet)
		fg_result.Stages = round.Stages
		//fg_result.Case = games.NotStartedYet

		round.Result[id] = fg_result
		//round.Result.stages
		//待定
		//round.Status = fg_result.Case
		round.TotalPoint = round.TotalPoint.Add(fg_result.Point)
		// j := 0
		// for fg_result.Case == games.Win {
		// 	j++
		// 	id = "f" + strconv.Itoa(i) + "-" + strconv.Itoa(j)
		// 	show_reelStrips = nextreelstrips.InvertRegularXYAxis().Strings()
		// 	fg_result, nextposition, nextreelstrips = s.CalcBigFGResult(unitbet, rtp, *nextreelstrips, nextposition, id, round.Brand, round.Username, show_reelStrips, bet)
		// 	fg_result.Stages = round.Stages
		// 	round.Result[id] = fg_result
		// 	//fmt.Printf("round.Result[%+v]: %+v\n", id, round.Result[id])
		// 	// ***********************[round.Status undetermined]***************************
		// 	//round.Status = round.Status.Push(ng_result.Case)
		// 	round.TotalPoint = round.TotalPoint.Add(fg_result.Point)
		// }
		fg_result.Case = games.NotStartedYet
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
	return *round, nil
}

// SpinNGReelStripts - spin NG reel strips
func (s *DCG008) FirstSpinFGReelStripts(unitbet string, rtp games.RTPs, pos []int) (*games.ReelStrips, [][]int, []string, error) {
	game_math, ok := dcg008_gameplay[unitbet]
	if !ok {
		return nil, nil, nil, errors.New("unitbet not found: " + unitbet)
	}
	pos[2] = pos[1]
	pos[3] = pos[1]
	//fmt.Printf("pos: %+v\n", pos)
	rtps := games.RTPs(rtp)
	fglen := game_math.fg_game.GetReelsLen(rtps)
	//fmt.Printf("fglen: %+v\n ", fglen)
	reelStrips := game_math.fg_game.BigSymbolsReelStrips(rtps, pos, bigReel)
	//fmt.Printf("reelStrips: %+v\n  ", reelStrips)
	Pos := make([][]int, len(fglen))

	for i := 0; i < len(reelStrips); i++ {

		Pos[i] = make([]int, len(reelStrips[i]))
		Pos[i][0] = pos[i]
		for j := 1; j < len(reelStrips[i]); j++ {
			if !bigReel[i] {
				Pos[i][j] = Pos[i][j-1] + 1
			} else {
				Pos[i][j] = Pos[i][0]
			}

		}
	}
	//fmt.Printf("🔵contiguousPos: %+v\n", Pos)

	//fmt.Printf("contiguousPos: %+v\n", Pos)
	show_reelStrips := game_math.fg_game.ShowBigSymbolReelStrips(rtps, pos, bigReel, 1, 1).InvertRegularXYAxis().Strings()

	//fmt.Printf("🔴show_reelStrips: %+v\n", show_reelStrips)
	return &reelStrips, Pos, show_reelStrips, nil
}

// CalcBigFGResult - calc Big Symbol Free Game  result
func (s *DCG008) CalcBigFGResult(unitbet, rtp string, reel games.ReelStrips, pos [][]int, id, brand, user string, symbols []string, bet decimal.Decimal) (*games.Records, [][]int, *games.ReelStrips) {
	type ExtraCascading struct {
		Cascadingpoint []string `json:"cascadingpoints"`
	}
	record := games.Records{
		Id:       id,
		Brand:    brand,
		Username: user,
		Case:     games.Lose,
		Extra:    make([]string, 0),
		Pickem:   []string{unitbet},
		Bet:      bet,
		Symbols:  symbols,
		Point:    decimal.NewFromFloat(0.0),
	}
	rtps := games.RTPs(rtp)
	fg_game := dcg008_gameplay[unitbet].fg_game
	fg_result_removedstrips, fg_result_newpos, fg_result_newstrips, fg_result := fg_game.CalcCascadingWayReel(rtps, pos, reel, bet)
	for i := 1; i < len(reel)-1; i++ {
		//if bigReel[i] {
		for j := 0; j < len(reel[i]); j++ {
			fg_result_newpos[i][j] = fg_result_newpos[i][len(reel[i])-1]
		}
		//}
		fg_result_newstrips = fg_game.Display(rtps, fg_result_newpos)
	}
	if !fg_result.IsWinGreaterThanZero() {
		record.Extra = []string{"0"}
	}
	i := 0
	for fg_result.IsWinGreaterThanZero() {
		i++
		//record.Case = games.Win
		//	println("111")
		record.Multiplier = record.Multiplier.Add(fg_result.TotalWin.Div(bet))
		//	println("222")
		record.Extra = append(record.Extra, fg_result.TotalWin.String())
		//	println("333")
		record.Point = record.Point.Add(fg_result.TotalWin)
		record.Symbols = append(append(record.Symbols, fg_result_removedstrips.InvertRegularXYAxis().Strings()...), fg_result_newstrips.InvertRegularXYAxis().Strings()...)

		fg_result_removedstrips, fg_result_newpos, fg_result_newstrips, fg_result = fg_game.CalcCascadingWayReel(rtps, fg_result_newpos, fg_result_newstrips, bet)
		for i := 1; i < len(reel)-1; i++ {
			//if bigReel[i] {
			for j := 0; j < len(reel[i]); j++ {
				fg_result_newpos[i][j] = fg_result_newpos[i][len(reel[i])-1]
			}
			//}
			fg_result_newstrips = fg_game.Display(rtps, fg_result_newpos)

		}
	}
	c, _ := json.Marshal(ExtraCascading{Cascadingpoint: record.Extra})
	record.Extra = []string{string(c)}
	return &record, fg_result_newpos, &fg_result_newstrips
}
