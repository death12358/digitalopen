package dcg008

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

// update existing reelstrips according to "EXPANDINGWILD" rules
func ExpandingWILD(r games.ReelStrips) games.ReelStrips {
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] == WW {
				for k := 0; k < len(r[i]); k++ {
					if r[i][k] != SF && r[i][k] != SB {
						r[i][k] = WW
					}
				}
			}
		}

	}
	return r
}

func (s *DCG008) NormalGameFlow(bet decimal.Decimal, unitbet, rtp string, position []int, round *games.Rounds) (games.Rounds, *games.ReelStrips) {
	// Check the length of symbols.
	// get the unitbet playsource.
	//  game_math, ok := dcg008_gameplay[unitbet]
	//  if !ok {
	//  	return nil, nil, errors.New("unitbet not found: " + unitbet)
	//  }

	//stage := 0 // 0,1,2,3...
	rtps := games.RTPs(rtp)
	round.TotalBet = bet
	//DELET DATA
	for k := range round.Result {
		delete(round.Result, k)
	}
	round.TotalPoint = decimal.Zero
	// ***********************[position undetermined]***************************
	//round.Position = round.Status.Push(games.FreeGame)

	// get the initial reelstips and other initial information.
	// 初始化 games.Records 結構
	// ng_result := &games.Records{
	// 	Id:       "0",
	// 	Brand:    round.Brand,
	// 	Username: round.Username,
	// 	Case:     games.Lose,
	// 	Pickem:   []string{""},
	// 	Symbols:  []string{""},
	// 	Bet:      round.TotalBet,
	// 	Point:    decimal.Zero,
	// }
	reelStrips, initialposition, show_reelStrips, _ := s.FirstSpinNGReelStripts(unitbet, rtps, position)

	// Change to NG Game Lightning symbols
	//round.Result["0"] = game_math.bgLightning_game.ReplaceNGSymbol(record)

	// Divided by unitbet
	// unit_bet, _ := decimal.NewFromString(unitbet)
	// bet := record.Bet.Mul(unit_bet)

	// Write result[0] and updat to status and totalpoint in the round.
	id := strconv.Itoa(0)
	ng_result, _, nextreelstrips := s.CalcNGResult(unitbet, rtp, *reelStrips, initialposition, id, round.Brand, round.Username, show_reelStrips, bet)
	//test
	//fmt.Printf("ng_result: %+v\n, nextposition: %+v\n ,nextreelstrips: %+v\n", ng_result, nextposition, nextreelstrips)

	round.Result[id] = ng_result
	round.Status = ng_result.Case
	round.TotalPoint = round.TotalPoint.Add(ng_result.Point)

	//Write result[1],result[2],result[3]... And update the status and totalpoint in the round.
	// for ng_result.Case == games.Win {
	// 	//round.Stages++
	// 	stage++
	// 	id = strconv.Itoa(int(stage))
	// 	show_reelStrips = nextreelstrips.InvertRegularXYAxis().Strings()
	// 	ng_result, nextposition, nextreelstrips = s.CalcNGResult(unitbet, rtp, *nextreelstrips, nextposition, id, round.Brand, round.Username, show_reelStrips, bet)
	// 	//fmt.Printf("ng_result: %+v\n, nextposition: %+v\n ,nextreelstrips: %+v\n", ng_result, nextposition, nextreelstrips)

	// 	//ng_result.Stages = 0
	// 	round.Result[id] = ng_result
	// 	//fmt.Printf("round.Result[%+v]: %+v\n", id, round.Result[id])
	// 	// ***********************[round.Status undetermined]***************************
	// 	//round.Status = round.Status.Push(ng_result.Case)
	// 	round.TotalPoint = round.TotalPoint.Add(ng_result.Point)
	// }

	return *round, nextreelstrips
}

// GetNGReelsLen - get the NG reels length
func (s *DCG008) GetNGReelsLen(unitbet, rtp string) ([]int, error) {
	// get the unitbet playsource.
	game_math, ok := dcg008_gameplay[unitbet]
	if !ok {
		return nil, errors.New("unitbet not found: " + unitbet)
	}
	rtps := games.RTPs(rtp)
	return game_math.ng_game.GetReelsLen(rtps), nil
}

// SpinNGReelStripts - spin NG reel strips
func (s *DCG008) FirstSpinNGReelStripts(unitbet string, rtp games.RTPs, pos []int) (*games.ReelStrips, [][]int, []string, error) {
	game_math, ok := dcg008_gameplay[unitbet]
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

func (s *DCG008) CalcNGResult(unitbet, rtp string, reel games.ReelStrips, pos [][]int, id, brand, user string, symbols []string, bet decimal.Decimal) (*games.Records, [][]int, *games.ReelStrips) {
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
	game_math := dcg008_gameplay[unitbet]
	reel = ExpandingWILD(reel)
	ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result := game_math.ng_game.CalcCascadingWayReel(rtps, pos, reel, bet)
	if !ng_result.IsWinGreaterThanZero() {
		record.Extra = []string{"0"}
	}
	//fmt.Printf("ng_result_removedstrips: %+v\n, ng_result_newpos: %+v\n ,ng_result_newstrips: %+v\n,ng_result: %+v\n", ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result)
	for ng_result.IsWinGreaterThanZero() {
		record.Case = games.Win
		record.Multiplier = record.Multiplier.Add(ng_result.TotalWin.Div(bet))
		record.Extra = append(record.Extra, ng_result.TotalWin.String())
		record.Point = record.Point.Add(ng_result.TotalWin)
		record.Symbols = append(append(record.Symbols, ng_result_removedstrips.InvertRegularXYAxis().Strings()...), ng_result_newstrips.InvertRegularXYAxis().Strings()...)
		reel = ExpandingWILD(ng_result_newstrips)
		ng_result_removedstrips, ng_result_newpos, ng_result_newstrips, ng_result = game_math.ng_game.CalcCascadingWayReel(rtps, ng_result_newpos, reel, bet)
	}
	c, _ := json.Marshal(ExtraCascading{Cascadingpoint: record.Extra})
	record.Extra = []string{string(c)}
	//else {
	// 	record.Case = games.Lose
	// }

	return &record, ng_result_newpos, &ng_result_newstrips
}
