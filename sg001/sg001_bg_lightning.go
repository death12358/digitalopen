package sg001

import (
	"errors"
	"strconv"

	"Golang/slot_digitalopen/games"

	"github.com/shopspring/decimal"
)

// games.BonusFreeGame2

// LightningFlow - FG Respin
//
//	初始次數為三次
func (s *SG001) LightningFlow(rtp string, bet decimal.Decimal, round *games.Rounds) (*games.Rounds, error) {
	round.Status = round.Status.Pop(games.Bonus)
	round.Status = round.Status.Push(games.BonusGame2)
	round.Position = games.BonusGame2

	round.Stages++
	id := strconv.Itoa(int(round.Stages))

	lightning_game := sg001_gameplay[round.Result["0"].Pickem[0]].bgLightning_game
	if round == nil {
		return nil, errors.New("round is nil")
	}

	reels := make(games.Reels, lightning_game.GetReelDef())
	record := lightning_game.HoldNGReel(&reels, round.Result["0"])
	//fmt.Printf("!!!!!!!!!\nround.Result[0]:%+v\n", round.Result["0"])
	//fmt.Printf("!!!!!!!!!\nrecord:%+v\n", record)

	record.Point = record.Multiplier.Mul(bet)
	round.TotalPoint = round.TotalPoint.Add(record.Point)
	round.Result[id] = record

	// 起始從2開始
	// idx := 2
	default_times := int64(3)
	for default_times > 0 {
		round.Stages++
		id := strconv.Itoa(int(round.Stages))
		record := &games.Records{
			Brand:      round.Brand,
			Username:   round.Username,
			Case:       games.NotStartedYet,
			Stages:     default_times, // default 2 1 0
			Pickem:     []string{""},
			Multiplier: decimal.Zero,
			Bet:        decimal.Zero,
			Point:      decimal.Zero,
		}
		/*bi, _ := lightning_game.CheckBlankSpaces(record.reels)
		if len(bi)==0{
			break
		}*/
		/*rs*/
		_, rc, err := lightning_game.Respin(&reels, record, id, bet)
		if err != nil {
			break
		}
		round.Result[rc.Id] = rc
		// reels = *rs
		record = rc
		default_times = rc.Stages
		round.TotalPoint = round.TotalPoint.Add(rc.Point)
		// idx++
	}

	// round.Stages = 1
	round.Result["1"].Case = games.Lose
	if !round.Result["1"].Point.IsZero() {
		round.Result["1"].Case = games.Win
	}
	return round, nil
}
