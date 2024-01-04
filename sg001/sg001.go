package sg001

import (
	"digitalopen/games"
	"digitalopen/games/random"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type SG001 struct {
	name string
	info string
	m    *sync.Mutex
}

func New() *SG001 {
	decimal.DivisionPrecision = 18

	return &SG001{
		name: "SG001",
		info: "SG001",
		m:    &sync.Mutex{},
	}
}

func (s *SG001) Name() string {
	return s.name
}

func (s *SG001) Info() string {
	return s.info
}

func (s *SG001) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if round.Position != 0 { /////////////????????????
		return s.Bonus(rtp, bet, pickem, &round)
	}
	if len(pickem) == 0 {
		return nil, errors.New("not pick unitbet.")
	}
	round.Start = time.Now().Unix()
	round.TotalBet = bet
	unitbet := pickem[0]

	// 取得轉輪表每輪長度
	ngReelsLen, err := s.GetNGReelsLen(unitbet, rtp)
	if err != nil {
		return nil, err
	}
	// 亂數位置
	pos := random.Intsn(ngReelsLen)
	/*
		TestWeightTable := []int{1, 1, 1}
		TestObjectTable := []int{1, 2, 3}

		TestGame := weights.NewGames(
			TestWeightTable,
			TestObjectTable,
		)
		dice := random.Intn(TestGame.Sum())
		pick, _ := TestGame.Pick(dice)
		if pick == 1 {
			pos = []int{0, 1, 1, 0, 0}
		}
		if pick == 2 {
			pos = []int{0, 0, 0, 0, 0}
		}
		if pick == 3 {
			pos = []int{46, 46, 3, 3, 0}
		}
	*/
	// // FG Test (BIG)
	//pos = []int{34, 5, 60, 23, 0}
	// // BG Test (二選一)
	//pos = []int{2, 24, 0, 3, 0}
	// // JP Test
	//pos = []int{34, 4, 2, 22, 0}
	//沒中獎盤面
	//pos = []int{0, 0, 0, 0, 0}
	//round, _ = s.SpinNormalGame(unitbet, rtp, pos, &round)
	round, reels := s.SpinNormalGame(unitbet, rtp, pos, &round)

	if s.IsBigFGWin(unitbet, reels) {
		round.Status = round.Status.Push(games.FreeGame)
		round.Position = round.Status.Push(games.FreeGame)
		round = *s.BigFGFlow(rtp, bet, &round)
	} else if s.IsJackpotTrigger(unitbet, reels) {
		round.Status = round.Status.Push(games.Jackpot)
		round.Position = round.Status.Push(games.Jackpot)
		round = *s.Jackpot(unitbet, bet, &round)
	}
	if s.IsBonusWin(unitbet, reels) {
		round.Status = round.Status.Push(games.Bonus)
		// 暫時更改Rest，待玩家選擇
		round.Position = games.Rest
	}

	round.Fisish = time.Now().Unix()

	return &round, nil
}

func (s *SG001) CalcRollRate(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal {
	s.m.Lock()
	defer s.m.Unlock()
	jp_game := sg001_gameplay[unitbet].jp_game
	return jp_game.CalcRollRate(totalbet)
}

func (s *SG001) BaseJackpot(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal {
	s.m.Lock()
	defer s.m.Unlock()
	jp_game := sg001_gameplay[unitbet].jp_game
	return jp_game.GetBasePool(totalbet)
}

// func (s *SG001) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
// 	if round.Position != 0 {
// 		return s.Bonus(rtp, bet, pickem, &round)
// 	}
// 	if len(pickem) == 0 {
// 		return nil, errors.New("not pick unitbet.")
// 	}
// 	round.Start = time.Now().Unix()
// 	unitbet := pickem[0]
// 	reels, display, err := s.SpinNGReelStripts(unitbet, rtp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	game_math := sg001_gameplay[unitbet]

// 	record := s.CalcNGResult(unitbet, reels, round.Id, round.Brand, round.Username, pickem, display, bet)

// 	round.Result["0"] = game_math.bgLightning_game.ReplaceNGSymbol(record)
// 	round.Status = round.Status.Push(record.Case)
// 	round.TotalBet = round.TotalBet.Add(bet)
// 	round.TotalPoint = round.TotalPoint.Add(record.Point)

// 	if s.IsBigFGWin(unitbet, reels) {
// 		round.Status = round.Status.Push(games.FreeGame)
// 		round.Position = round.Status.Push(games.FreeGame)
// 		round = *s.BigFGFlow(rtp, bet, &round)
// 	}
// 	if s.IsBonusWin(unitbet, reels) {
// 		round.Status = round.Status.Push(games.Bonus)
// 		// 暫時更改Rest，待玩家選擇
// 		round.Position = games.Rest
// 	}
// 	round.Fisish = time.Now().Unix()

// 	return &round, nil
// }

// Bonus - bonus game
func (s *SG001) Bonus(rtp string, bet decimal.Decimal, pickem []string, round *games.Rounds) (*games.Rounds, error) {
	chosic := pickem[0]
	round.Result["0"].Pickem = append(round.Result["0"].Pickem, chosic)

	switch chosic {
	case bg_hifg_def:
		return s.HiFGFlow(rtp, bet, round)
	case bg_lightning_def:
		// unit_bet, _ := decimal.NewFromString(round.Result["0"].Pickem[0])
		// bet = round.Result["0"].Bet.Div(unit_bet)
		return s.LightningFlow(rtp, bet, round)
	default:
		return nil, errors.New("not pick chosic.")
	}

}

// Jackpot - jackpot game
func (s *SG001) Jackpot(unitbet string, bet decimal.Decimal, round *games.Rounds) *games.Rounds {
	jp_game := sg001_gameplay[unitbet].jp_game
	jp_sym := jp_game.PickJackpot()

	bet = round.Result["0"].Bet
	round.Stages++
	id := strconv.Itoa(int(round.Stages))
	point := jp_game.GetJackpotPay(jp_sym)

	record := &games.Records{
		Id:         id,
		Brand:      round.Brand,
		Username:   round.Username,
		Case:       games.NotStartedYet,
		Pickem:     []string{""},
		Symbols:    []string{jp_sym.String()},
		Multiplier: point,
		Point:      point.Mul(bet),
	}
	round.Result[id] = record
	round.TotalPoint = round.TotalPoint.Add(record.Point)

	return round
}
