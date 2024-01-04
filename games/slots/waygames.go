package slots

import (
	"digitalopen/games"

	"github.com/shopspring/decimal"
)

// WayGames - WayGames 結構
type WayGames struct {
	reelStripsTable games.ReelStripList
	reelStripsDef   *ReelStripsDef
	reelLen         games.ReelStripLengthTable
	payTable        *games.PayTable
	symbolList      []games.Symbol
	scatter         []games.Symbol
	unitbet         decimal.Decimal
}

// NewWayGames - 建立 WayGames
//
//	@param reelStripsTable			轉輪表 < RTP, 轉輪表 >
//	@param reelStripsDef			轉輪個數，陣列大小為幾輪，陣列內容為每輪的數量
//	@param payTable					賠付表
//	@param symbolList				獎圖列表，可使用預設 slots.SymbolList，也可以自定義
//	@param scatter Scatter			特殊獎圖，可使用預設 slots.Scatter，也可以自定義
//	@param unitbet decimal.Decimal	單位投注
//	@return *WayGames WayGames 物件
func NewWayGames(
	reelStripsTable games.ReelStripList,
	reelStripsDef *ReelStripsDef,
	payTable *games.PayTable,
	symbolList []games.Symbol,
	scatter []games.Symbol,
	unitbet decimal.Decimal,
) *WayGames {

	// get reel length
	reelLenTable := make(games.ReelStripLengthTable)
	for rtp, reelStrips := range reelStripsTable {
		reelLenTable[rtp] = reelStrips.Lengths()
	}

	return &WayGames{
		reelStripsTable: reelStripsTable,
		reelStripsDef:   reelStripsDef,
		reelLen:         reelLenTable,
		payTable:        payTable,
		symbolList:      symbolList,
		scatter:         scatter,
		unitbet:         unitbet,
	}
}

// GetReelsLen - 取得轉輪長度
//
//	@param rtp string	RTP
//	return ReelStripLengths	轉輪長度
func (w *WayGames) GetReelsLen(rtp games.RTPs) []int {
	return w.reelLen[rtp]
}

// ContiguousReelStrips - 連續轉輪
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//
// return ReelStrips				轉輪表
func (w *WayGames) ContiguousReelStrips(rtp games.RTPs, position []int) games.ReelStrips {
	return w.reelStripsTable[rtp].ContiguousReelStrips(*w.reelStripsDef, position)
}

// ShowReelStrips - 顯示表演用轉輪表
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//	@param top_shift int			轉輪上方空格數
//	@param end_shift int			轉輪下方空格數
//	@return ReelStrips				轉輪表
func (w *WayGames) ShowReelStrips(rtp games.RTPs, position []int, top_shift int, end_shift int) games.ReelStrips {
	return w.reelStripsTable[rtp].ShowReelStrips(*w.reelStripsDef, position, top_shift, end_shift)
}

// BigSymbolsReelStrips - 顯示連續轉輪((大獎圖)
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//	@param bigSym []bool			大獎圖
//
// return ReelStrips				轉輪表
func (w *WayGames) BigSymbolsReelStrips(rtp games.RTPs, position []int, bigSym []bool) games.ReelStrips {
	return w.reelStripsTable[rtp].RepeatedReelStrips(*w.reelStripsDef, position, bigSym)
}

// ShowBigSymbolReelStrips - 顯示表演用轉輪表(大獎圖)
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//	@param bigSym []bool			大獎圖
//	@param top_shift int			轉輪上方空格數
//	@param end_shift int			轉輪下方空格數
//	@return ReelStrips				轉輪表
func (w *WayGames) ShowBigSymbolReelStrips(rtp games.RTPs, position []int, bigSym []bool, top_shift int, end_shift int) games.ReelStrips {
	return w.reelStripsTable[rtp].ShowRepeatedReelStrips(*w.reelStripsDef, position, bigSym, top_shift, end_shift)
}

// CalcWayReel
//
//	@param bet decimal.Decimal	下注金額
//
// return WinDetail		獎金
func (w *WayGames) CalcWayReel(reels games.ReelStrips, bet decimal.Decimal) *WinDetail {
	result := NewWinDetail(bet)

	// 1. 排除第一輪重複 Symbol
	tarSymbol := reels[0].RemoveDuplicates()
	// fmt.Printf("tarSymbol: %v\n", tarSymbol)
	bet = bet.Div(w.unitbet)

	// 2. 計算連續數量(match), 並計算中獎數量(multi)
	for _, symbol := range tarSymbol {
		match, multi := reels.CalcSymbolsMatchFromLeft(symbol, WW)
		// fmt.Printf("symbol: %v, match: %v, multi: %v\n", symbol, match, multi)
		m_count := len(match) - 1
		point := w.payTable.CalcPaysTable(symbol.Int(), m_count, multi).Mul(bet)
		result.Add(point)
		result.Wins = append(result.Wins, NewWayWins(symbol, m_count, multi, point))
	}

	return result
}

// CalcScatter - 計算 Scatter
//
//	Way 為獎圖個數
func (w *WayGames) CalcScatter(reels games.ReelStrips, bet decimal.Decimal, tar games.Symbol) *WinDetail {
	result := NewWinDetail(bet)
	bet = bet.Div(w.unitbet)

	match, count := reels.CalcSymbolMatches(tar)
	point := decimal.Zero
	if match > 0 {
		point = w.payTable.CalcPaysTable(tar.Int(), match-1, 1).Mul(bet)
	}
	result.Add(point)
	result.Wins = append(result.Wins, NewWayWins(tar, match, count, point))

	return result
}

// Display-根据位置陈列盘面
// p 位置
func (w *WayGames) Display(rtp games.RTPs, p [][]int) games.ReelStrips {
	return w.reelStripsTable[rtp].Display(p)
}

// Abs 绝对值
func Abs(a int) int {

	if a < 0 {
		a *= (-1)
	}
	return a
}

// CalcCascadingWayReel 输出消除一次的结果以及消除阵列/新阵列/得分信息
// @param rtps标识
// @param pos 位置
// @param reels 盘面
// @param bet decimal.Decimal	下注金額
//
//	return ReelStrips    被消除盘面（-1表消除）
//	return AllReelPos    新位置
//	return ReelStrips    新盘面
//	return WinDetail     本次消除獎金信息
func (w *WayGames) CalcCascadingWayReel(rtp games.RTPs, pos [][]int, reels games.ReelStrips, bet decimal.Decimal) (games.ReelStrips, [][]int, games.ReelStrips, *WinDetail) {
	//reels := w.reelStripsTable[rtp].Display(pos)
	R := reels.Deepcopy()

	result := NewWinDetail(bet)
	lenth := w.GetReelsLen(rtp)
	// 1. 排除第一輪重複 Symbol
	tarSymbol := reels[0].RemoveDuplicates()
	//fmt.Printf("tarSymbol: %v\n", tarSymbol)
	bet = bet.Div(w.unitbet)
	// 2. 計算連續數量(match), 並計算中獎數量(multi)
	for _, symbol := range tarSymbol {
		b := true
		//fmt.Printf("mATRIX: %v", R)
		match, multi := reels.CalcSymbolsMatchFromLeft(symbol, WW)
		//fmt.Printf("symbol: %v, match: %v, multi: %v\n", symbol, match, multi)
		m_count := len(match) - 1
		point := w.payTable.CalcPaysTable(symbol.Int(), m_count, multi).Mul(bet)
		//3.排除不可消除的symbol
		for _, eachscattersymbol := range w.scatter {

			if symbol == eachscattersymbol {
				b = false
				//fmt.Printf("b**************************: %v\n", b)
				break
			}
		}
		//4.标记消除了的symbol
		if !point.IsZero() && b {
			R = R.RemoveMark([]games.Symbol{symbol, WW}, len(match))
		}
		result.Add(point)
		result.Wins = append(result.Wins, NewWayWins(symbol, m_count, multi, point))

	}
	//5.计算新盘面的位置
	for i := 0; i < len(pos); i++ {
		for j := 0; j < len(pos[i]); j++ {
			//排除0的存在
			if pos[i][j] == 0 {
				pos[i][j] = lenth[i]
			}
			if R[i][j] == -1 {
				pos[i][j] = -pos[i][j]
			}
		}
	}

	var fallnumbers int
	var temp int
	position1 := make([][]int, len(pos))
	position := make([][]int, len(pos))
	for i := 0; i < len(pos); i++ {
		fallnumbers = 0
		temp = Abs(pos[i][0])
		for j := 0; j < len(pos[i]); j++ {
			if pos[i][j] < 0 {
				fallnumbers += 1
			} else {
				position1[i] = append(position1[i], pos[i][j]%lenth[i])
			}
		}
		for k := fallnumbers; k > 0; k-- {
			position[i] = append(position[i], (temp-k+lenth[i])%lenth[i])
		}
		position[i] = append(position[i], position1[i]...)
	}

	return R, position, w.Display(rtp, position), result
}

func (w *WayGames) CalcWayReel_WWmulti(reels games.ReelStrips, bet decimal.Decimal, WWmulti decimal.Decimal) (*WinDetail, decimal.Decimal) {
	result := NewWinDetail(bet)

	// 1. 排除第一輪重複 Symbol
	tarSymbol := reels[0].RemoveDuplicates()
	// fmt.Printf("tarSymbol: %v\n", tarSymbol)
	bet = bet.Div(w.unitbet)
	WWLines := 0
	for _, symbol := range tarSymbol {
		if symbol != 13 {
			match, multi := reels.CalcSymbolsMatchFromLeft(symbol, WW)
			matchNoWW, multiNoWW := reels.CalcSymbolsMatchFromLeft(symbol)
			// fmt.Printf("symbol: %v, match: %v, multi: %v\n", symbol, match, multi)
			m_count := len(match) - 1
			m_countNoWW := len(matchNoWW) - 1
			if (m_count > m_countNoWW || multi > multiNoWW) && m_count >= 2 {
				WWLines++
			}
		}
	}
	if WWLines == 0 {
		WWmulti = decimal.NewFromInt(int64(1))
	}
	// 2. 計算連續數量(match), 並計算中獎數量(multi)
	for _, symbol := range tarSymbol {
		if symbol != 13 {
			match, multi := reels.CalcSymbolsMatchFromLeft(symbol, WW)
			matchNoWW, multiNoWW := reels.CalcSymbolsMatchFromLeft(symbol)
			// fmt.Printf("symbol: %v, match: %v, multi: %v\n", symbol, match, multi)
			m_count := len(match) - 1
			m_countNoWW := len(matchNoWW) - 1
			if WWmulti.GreaterThan(decimal.NewFromInt(int64(1))) {
				point := decimal.Zero
				if m_count == m_countNoWW {
					pointNoWW := w.payTable.CalcPaysTable(symbol.Int(), m_countNoWW, multiNoWW).Mul(bet)
					pointWithWW := w.payTable.CalcPaysTable(symbol.Int(), m_count, multi-multiNoWW).Mul(bet).Mul(WWmulti)
					point = pointNoWW.Add(pointWithWW)
				} else {
					point = w.payTable.CalcPaysTable(symbol.Int(), m_count, multi).Mul(bet).Mul(WWmulti)
				}
				result.Add(point)
				result.Wins = append(result.Wins, NewWayWins(symbol, m_count, multi, point))
			} else {
				point := w.payTable.CalcPaysTable(symbol.Int(), m_count, multi).Mul(bet)
				result.Add(point)
				result.Wins = append(result.Wins, NewWayWins(symbol, m_count, multi, point))
			}
		}
		if symbol == 13 {
			match, multi := reels.CalcSymbolsMatchFromLeft(symbol)
			// fmt.Printf("symbol: %v, match: %v, multi: %v\n", symbol, match, multi)
			m_count := len(match) - 1
			point := w.payTable.CalcPaysTable(symbol.Int(), m_count, multi).Mul(bet)
			result.Add(point)
			result.Wins = append(result.Wins, NewWayWins(symbol, m_count, multi, point))
		}
	}

	return result, WWmulti
}
