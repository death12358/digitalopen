package dcg014

import (
	"errors"
	"sync"
	"time"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/shopspring/decimal"
)

type DCG014 struct {
	name string
	info string
	m    *sync.Mutex
}

func New() *DCG014 {
	return &DCG014{
		name: "DCG014",
		info: "DCG014",
		m:    &sync.Mutex{},
	}
}
func (s *DCG014) Name() string {
	return s.name
}
func (s *DCG014) Info() string {
	return s.info
}

func (s *DCG014) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
	s.m.Lock()
	defer s.m.Unlock()
	if len(pickem) == 0 {
		return nil, errors.New("not pick unitbet.")
	}

	round.Start = time.Now().Unix()
	round.TotalBet = bet
	unitbet := pickem[0]

	// Get the length of reels.
	ngReelsLen, err := s.GetNGReelsLen(unitbet, rtp)
	//fmt.Println(ngReelsLen)
	if err != nil {
		return nil, err
	}
	// get the random position.
	pos := random.Intsn(ngReelsLen)
	//fmt.Println("laji: v\n", pos)
	//pos = []int{150, 151, 18, 84, 9, 119}
	//pos = []int{15, 131, 164, 86, 11, 44}
	// BG Test f
	//pos = []int{66, 86, 29, 13, 14}
	// // FG&BG Test y
	//pos = []int{27, 21, 54, 32, 52}

	round, reels := s.NormalGameFlow(bet, unitbet, rtp, pos, &round)
	//Tes/test////////////////////////////////
	// byteRound, _ := json.Marshal(round)
	// fmt.Println(string(byteRound))
	//////////////////////////////////////////
	if s.IsFGWin(unitbet, reels) {
		count := s.NumbersofFGWin(unitbet, reels)
		round.Status = round.Status.Push(games.FreeGame)
		round.Position = round.Status.Push(games.FreeGame)
		round, _ = s.FreeGameFlow(bet, unitbet, rtp, &round, count)

	}

	if s.IsBonusWin(unitbet, reels) {
		round.Status = round.Status.Push(games.BonusGame1)
		round.Position = round.Status.Push(games.BonusGame1)
		round, _ = s.EggFlow(bet, unitbet, &round)

	}

	round.Fisish = time.Now().Unix()

	return &round, nil
}
