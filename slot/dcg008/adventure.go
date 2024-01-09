package dcg008

import (
	"encoding/json"
	"strconv"

	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

func (s *DCG008) AdventureFlow(bet decimal.Decimal, unitbet string, round *games.Rounds) (games.Rounds, error) {
	//fmt.Printf("Go To AdventureFlow!!!!!!!!!!!!!!!!!!!!!")
	type ExtraAdventure struct {
		Floor string `json:"floor"`
		HP    string `json:"hp"`
	}
	adventure_game := dcg008_gameplay[unitbet].bg_game
	round.Status = round.Status.Push(games.Bonus)
	round.Position = round.Position.Push(games.Bonus)
	hp := int64(2)

	adventure_process := adventure_game.PickProcess()
	round.Stages++
	// fmt.Printf("adventure_process: %v\n ", adventure_process)
	//	round.TotalPoint = round.TotalPoint.Add(record.Point)
	//剩余数？
	for i := 0; i < len(adventure_process); i++ {
		score := adventure_game.CalcAdventureOneRound(adventure_process[i], bet)
		if score.IsZero() {
			hp -= 1
		}

		id := strconv.Itoa(int(round.Stages))
		record := &games.Records{
			Id:         id,
			Brand:      round.Brand,
			Username:   round.Username,
			Case:       games.NotStartedYet,
			Stages:     round.Stages,
			Extra:      make([]string, 0),
			Pickem:     []string{unitbet},
			Symbols:    []string{adventure_process[i].String()},
			Multiplier: score.Div(bet),
			Bet:        bet,
			Point:      score,
		}
		round.Result[id] = record
		round.TotalPoint = round.TotalPoint.Add(record.Point)
		round.Stages++
		//[]string{strconv.FormatInt(int64(i+1), 10), strconv.FormatInt(hp, 10)},
		c, _ := json.Marshal(ExtraAdventure{Floor: strconv.FormatInt(int64(i+1), 10),
			HP: strconv.FormatInt(hp, 10)})
		record.Extra = []string{string(c)}

	}

	return *round, nil
}
