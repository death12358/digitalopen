package games

import (
	"github.com/shopspring/decimal"
)

func init() {
	decimal.DivisionPrecision = 18
}

// Games - 遊戲定義結構，統一遊戲操作介面
type Games struct {
	src      Source
	fair_src FairnessSource
	jp_src   JackpotSource
	out_src  OutputSource
}

// NewGames - 建立遊戲定義
//  @param src		遊戲來源
//  @return *Games	遊戲定義物件
func NewGames(src Source) *Games {
	if src == nil {
		panic("Source is nil")
	}

	fair_src, _ := src.(FairnessSource)
	jp_src, _ := src.(JackpotSource)
	out_src, _ := src.(OutputSource)

	return &Games{
		src:      src,
		fair_src: fair_src,
		jp_src:   jp_src,
		out_src:  out_src,
	}
}

// Name - 遊戲名稱
//  @return string	遊戲名稱
func (g *Games) Name() string {
	return g.src.Name()
}

// Info - 遊戲資訊
//  @return string	遊戲資訊
func (g *Games) Info() string {
	return g.src.Info()
}

// Spin - 遊戲開始
//  @param rtp				RTP
//  @param bet				投注
//  @param pickem			選擇
//  @param rounds			回合結構
func (g *Games) Spin(rtp string, bet decimal.Decimal, pickem []string, round Rounds) (*Rounds, error) {
	return g.src.Spin(rtp, bet, pickem, round)
}

// Bet - 投注
//  @param seeds			隨機種子
//  @param bet				投注
//  @param pickem			選擇
//  @param rounds			回合結構
func (g *Games) Bet(seeds []byte, bet decimal.Decimal, pickem []string, round Rounds) (*Rounds, error) {
	if g.fair_src == nil {
		return nil, _ErrNotFairnessSource
	}

	return g.fair_src.Bet(seeds, bet, pickem, round)
}

// Output - 輸出
//  @param rtp				RTP
//  @param count			次數
//  @return string			輸出結果
func (g *Games) Output(rtp string, count int) (string, error) {
	if g.out_src == nil {
		return "", _ErrNotOutputSource
	}

	return g.out_src.Output(rtp, count)
}

// CalcRollRate - 計算RollRate
//  @param unitbet			單位投注
//  @param rtp				RTP
//  @param totalbet			總投注
//  @return []decimal.Decimal		RollRate
func (g *Games) CalcRollRate(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal {
	if g.jp_src == nil {
		return nil
	}

	return g.jp_src.CalcRollRate(unitbet, rtp, totalbet)
}

// BaseJackpot - 計算基礎彩金
//  @param unitbet			單位投注
//  @param rtp				RTP
//  @param totalbet			總投注
//  @return []decimal.Decimal		基礎彩金
func (g *Games) BaseJackpot(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal {
	if g.jp_src == nil {
		return nil
	}

	return g.jp_src.BaseJackpot(unitbet, rtp, totalbet)
}

// Source -
// The source of the game.
type Source interface {
	Name() string
	Info() string
	Spin(rtp string, bet decimal.Decimal, pickem []string, round Rounds) (*Rounds, error)
}

// StartRound
// BetRound
// EndRound

// FairnessSource -
// The fairness of the game.
type FairnessSource interface {
	Source
	Bet(seeds []byte, bet decimal.Decimal, pickem []string, round Rounds) (*Rounds, error)
}

// OutputSource -
// Simulate the output of the game.
type OutputSource interface {
	Source
	Output(rtp string, count int) (string, error)
}

// JackpotSource -
// The jackpot of the game.
type JackpotSource interface {
	Source
	CalcRollRate(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal
	BaseJackpot(unitbet, rtp string, totalbet decimal.Decimal) []decimal.Decimal
}
