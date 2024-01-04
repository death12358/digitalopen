// Path: slots/slots.go
// Compare this snippet from slots/slots.go:
// package slots
package slots

import (
	"github.com/death12358/digitalopn/games"

	"github.com/shopspring/decimal"
)

func init() {
	decimal.DivisionPrecision = 18
}

// WinDetail - 中獎結果結構
//
//	@param Wins			中獎結果
//	@param TotalBet		總投注
//	@param TotalWin		總獎金
type WinDetail struct {
	Wins     []Wins
	ToatlBet decimal.Decimal
	TotalWin decimal.Decimal
}

// NewWinDetail - 建立 WinDetail
//
//	@param totalBet		總投注
//	@return *WinDetail	WinDetail 物件
func NewWinDetail(totalBet decimal.Decimal) *WinDetail {
	return &WinDetail{
		Wins:     []Wins{},
		ToatlBet: totalBet,
	}
}

// Add - 加總獎金
//
//	@param win		獎金
//	@return *WinDetail	WinDetail 物件
func (w *WinDetail) Add(win decimal.Decimal) *WinDetail {
	w.TotalWin = w.TotalWin.Add(win)
	return w
}

// IsWinGreaterThanZero - TotoalWin 是否大於 0
//
//	@return bool	是否大於 0
func (w *WinDetail) IsWinGreaterThanZero() bool {
	return w.TotalWin.GreaterThan(decimal.Zero)
}

// Wins - 獎金結構
type Wins struct {
	Symbols games.Symbol
	Match   int
	Ways    int
	Line    int
	Win     decimal.Decimal
}

// NewWins - 建立 Wins
//
//	@param symbols	獎圖
//	@param match	連續匹配數量
//	@param win		獎金
//	@return *Wins	Wins 物件
func NewWins(symbols games.Symbol, match int, win decimal.Decimal) Wins {
	return Wins{
		Symbols: symbols,
		Match:   match,
		Win:     win,
	}
}

// NewWayWins - 建立 WayGame Wins
//
//	@param symbols	獎圖
//	@param match	連續匹配數量
//	@param ways		方式: WayGame 可直接依照每輪計算做判斷
//	@param win		獎金
//	@return *Wins	Wins 物件
func NewWayWins(symbols games.Symbol, match, ways int, win decimal.Decimal) Wins {
	return Wins{
		Symbols: symbols,
		Match:   match,
		Ways:    ways,
		Win:     win,
	}
}

// NewLineWins - 建立 WayGame Wins
//
//	@param symbols	獎圖
//	@param match	連續匹配數量
//	@param line		方式: LineGame 第幾條線
//	@param win		獎金
//	@return *Wins	Wins 物件
func NewLineWins(symbols games.Symbol, match, line int, win decimal.Decimal) Wins {
	return Wins{
		Symbols: symbols,
		Match:   match,
		Line:    line,
		Win:     win,
	}
}
