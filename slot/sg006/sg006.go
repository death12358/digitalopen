package sg006

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/slot/sg006/vs"
	"github.com/shopspring/decimal"
)

type SG006 struct {
	name string
	info string
	m    *sync.Mutex
}

func New() *SG006 {
	return &SG006{
		name: "SG006",
		info: "SG006",
		m:    &sync.Mutex{},
	}
}

func (s *SG006) Name() string {
	return s.name
}

func (s *SG006) Info() string {
	return s.info
}

func (s *SG006) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
	s.m.Lock()
	defer s.m.Unlock()

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

	// // FreeGame Test
	//pos = []int{1, 0, 0, 0, 1}
	// // VSgame Test
	//pos = []int{26, 11, 8, 8, 29}
	round, reels := s.SpinNormalGame(unitbet, rtp, pos, &round)

	if s.IsFGWin(unitbet, reels) {
		round.Status = round.Status.Push(games.BonusGame2)
		round.Position = round.Status.Push(games.BonusGame2)
		return s.FGFlow(rtp, bet, &round)
	} else {
		VSReel := vs.CreateVSReel(rtp, reels)
		M_VSReel, _ := json.Marshal(VSReel)
		round.Result["0"].Extra = append(round.Result["0"].Extra, "vsreel:")
		round.Result["0"].Extra = append(round.Result["0"].Extra, string(M_VSReel))

		if ok, VSTimes := s.IsVSWin(unitbet, reels, VSReel); ok {
			round.Status = round.Status.Push(games.BonusGame1)
			round.Position = round.Status.Push(games.BonusGame1)
			return s.SpinVSFlow(rtp, bet, &round, VSTimes)
		}
	}

	round.Fisish = time.Now().Unix()

	return &round, nil
}
