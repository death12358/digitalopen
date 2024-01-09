package dcg008

import (
	"errors"
	"sync"
	"time"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/shopspring/decimal"
)

type DCG008 struct {
	name string
	info string
	m    *sync.Mutex
}

func New() *DCG008 {
	decimal.DivisionPrecision = 18
	return &DCG008{
		name: "dcg008",
		info: "dcg008",
		m:    &sync.Mutex{},
	}
}
func (s *DCG008) Name() string {
	return s.name
}
func (s *DCG008) Info() string {
	return s.info
}

func (s *DCG008) Spin(rtp string, bet decimal.Decimal, pickem []string, round games.Rounds) (*games.Rounds, error) {
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
	//fmt.Println(ngReelsLen)
	if err != nil {
		return nil, err
	}
	// 亂數位置
	pos := random.Intsn(ngReelsLen)
	//fmt.Println("laji: v\n", pos)
	//pos = []int{93, 7, 123, 69, 100}
	// BG Test
	//pos = []int{66, 86, 29, 13, 14}
	// // FG Test
	//pos = []int{135, 82, 135, 154, 30}
	//expanding wild test
	//pos = []int{0, 0, 0, 0, 15}
	round, reels := s.NormalGameFlow(bet, unitbet, rtp, pos, &round)
	//Tes/test////////////////////////////////
	// byteRound, _ := json.Marshal(round)
	// fmt.Println(string(byteRound))
	//////////////////////////////////////////
	if s.IsBigFGWin(unitbet, reels) {
		count := s.NumbersofFGWin(unitbet, reels)
		//fmt.Printf("Total: %v\n", count)
		round.Status = round.Status.Push(games.FreeGame)
		round.Position = round.Status.Push(games.FreeGame)
		round, _ = s.BigFreeGameFlow(bet, unitbet, rtp, &round, count)
	}

	if s.IsBonusWin(unitbet, reels) {
		//fmt.Printf("GO TO BonusGame!!!!!!!!!!!!!!!!!!!!!!!!")
		round.Status = round.Status.Push(games.BonusGame1)
		round.Position = round.Status.Push(games.BonusGame1)
		round, _ = s.AdventureFlow(bet, unitbet, &round)
	}

	round.Fisish = time.Now().Unix()

	return &round, nil
}
