package dcg014

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

func (s *DCG014) NormalGameFlow(bet decimal.Decimal, unitbet, rtp string, position []int, round *games.Rounds) (games.Rounds, *games.ReelStrips) {
	// Check the length of symbols.
	// get the unitbet playsource.
	//  game_math, ok := dcg014_gameplay[unitbet]
	//  if !ok {
	//  	return nil, nil, errors.New("unitbet not found: " + unitbet)
	//  }
	//grow_game := dcg014_gameplay[unitbet].fg_game
	//stage := 0 // 0,1,2,3...
	//rtps := games.RTPs(rtp)
	rtps := games.RTPs(rtp)
	round.TotalBet = bet
	//DELET DATA
	for k := range round.Result {
		delete(round.Result, k)
	}
	reelStrips, initialposition, show_reelStrips, _ := s.FirstSpinNGReelStripts(unitbet, rtps, position)
	//fmt.Printf(" reelStrips: %+v\n initialposition: %+v\n show_reelStrips: %+v\n", reelStrips, initialposition, show_reelStrips)
	id := strconv.Itoa(0)
	ng_result, _, nextreelstrips := s.CalcNGResult(unitbet, rtp, *reelStrips, initialposition, id, round.Brand, round.Username, show_reelStrips, bet)

	//test
	//fmt.Printf("ng_result: %+v\n", ng_result)
	//fmt.Printf("nextreelstrips: %+v\n", nextreelstrips)
	round.Result[id] = ng_result
	round.Status = ng_result.Case
	round.TotalPoint = round.TotalPoint.Add(ng_result.Point)

	return *round, nextreelstrips
}

// GetNGReelsLen - get the NG reels length
func (s *DCG014) GetNGReelsLen(unitbet, rtp string) ([]int, error) {
	// get the unitbet playsource.
	game_math, ok := dcg014_gameplay[unitbet]
	if !ok {
		return nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	return game_math.ng_game.GetReelsLen(rtps), nil
}

// SpinNGReelStripts - spin NG reel strips
func (s *DCG014) FirstSpinNGReelStripts(unitbet string, rtp games.RTPs, pos []int) (*games.ReelStrips, [][]int, []string, error) {
	game_math, ok := dcg014_gameplay[unitbet]
	if !ok {
		return nil, nil, nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	nglen := game_math.ng_game.GetReelsLen(rtps)
	reelStrips := game_math.ng_game.ContiguousReelStrips(rtps, pos)
	//fmt.Printf("reelStrips: %+v\n", reelStrips)
	contiguousPos := make([][]int, len(nglen))
	for i := 0; i < len(reelStrips); i++ {
		contiguousPos[i] = make([]int, len(reelStrips[i]))
		contiguousPos[i][0] = pos[i]
		for j := 1; j < len(reelStrips[i]); j++ {
			contiguousPos[i][j] = (contiguousPos[i][j-1] + 1) % nglen[i]
		}
	}
	//fmt.Printf("🔵contiguousPos: %+v\n", contiguousPos)
	show_reelStrips := game_math.ng_game.ShowReelStrips(rtps, pos, 1, 1).InvertRegularXYAxis().Strings()
	//fmt.Printf("🔴show_reelStrips: %+v\n", show_reelStrips)
	return &reelStrips, contiguousPos, show_reelStrips, nil
}

// CalcNGResult - calc NG result
func (s *DCG014) CalcNGResult(unitbet, rtp string, reel games.ReelStrips, pos [][]int, id, brand, user string, symbols []string, bet decimal.Decimal) (*games.Records, [][]int, *games.ReelStrips) {
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
		Symbols:  symbols,
		Bet:      bet,
		Point:    decimal.NewFromFloat(0.0),
	}
	rtps := games.RTPs(rtp)
	game_math := dcg014_gameplay[unitbet]

	ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result := game_math.ng_game.CalcCascadingWayReel(rtps, pos, reel, bet)
	if !ng_result.IsWinGreaterThanZero() {
		record.Extra = []string{"0"}
	}
	//fmt.Printf(" ng_result_removedstrips: %+v\n ng_result_newpos: %+v\n ng_result_newstrips: %+v\n ng_result: %+v\n", ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result)
	for ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = record.Multiplier.Add(ng_result.TotalWin.Div(bet))

		record.Extra = append(record.Extra, ng_result.TotalWin.String())
		record.Point = record.Point.Add(ng_result.TotalWin)

		record.Symbols = append(append(record.Symbols, ng_result_removedstrips.InvertRegularXYAxis().Strings()...), ng_result_newstrips.InvertRegularXYAxis().Strings()...)
		//fmt.Printf(" ng_result_removedstrips: %+v\n ng_result_newpos: %+v\n ng_result_newstrips: %+v\n ng_result: %+v\n", ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result)
		ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result = game_math.ng_game.CalcCascadingWayReel(rtps, ng_result_newpos, ng_result_newstrips, bet)

	}
	c, _ := json.Marshal(ExtraCascading{Cascadingpoint: record.Extra})
	record.Extra = []string{string(c)}
	//else {
	// 	record.Case = games.Lose
	// }

	return &record, ng_result_newpos, &ng_result_newstrips
}
