package slots

import (
	"github.com/shopspring/decimal"
	"gitlab.com/gaas_module/games"
)

// LineGames - LineGames 結構
type LineGames struct {
	reelStripsTable games.ReelStripList
	reelStripsDef   *ReelStripsDef
	reelLen         games.ReelStripLengthTable
	payTable        *games.PayTable
	payLine         *games.PayLine
	kWinLines       games.KWinLines
	symbolList      []games.Symbol
	scatter         []games.Symbol
	unitbet         decimal.Decimal
}

// NewLineGames - 建立 LineGames
//
//	@param reelStripsTable			轉輪表 < RTP, 轉輪表 >
//	@param reelStripsDef			轉輪個數，陣列大小為幾輪，陣列內容為每輪的數量
//	@param payTable					賠付表
//	@param payLine					賠付線
//	@param kWinLines                幾條賠付線
//	@param symbolList				獎圖列表，可使用預設 slots.SymbolList，也可以自定義
//	@param scatter Scatter			特殊獎圖，可使用預設 slots.Scatter，也可以自定義
//	@param unitbet decimal.Decimal	單位投注
//	@return *LineGames LineGames 物件
func NewLineGames(
	reelStripsTable games.ReelStripList,
	reelStripsDef *ReelStripsDef,
	payTable *games.PayTable,
	payLine *games.PayLine,
	kWinLines games.KWinLines,
	symbolList []games.Symbol,
	scatter []games.Symbol,
	unitbet decimal.Decimal,
) *LineGames {

	// get reel length
	reelLenTable := make(games.ReelStripLengthTable)
	for rtp, reelStrips := range reelStripsTable {
		reelLenTable[rtp] = reelStrips.Lengths()
	}

	return &LineGames{
		reelStripsTable: reelStripsTable,
		reelStripsDef:   reelStripsDef,
		reelLen:         reelLenTable,
		payTable:        payTable,
		payLine:         payLine,
		kWinLines:       kWinLines,
		symbolList:      symbolList,
		scatter:         scatter,
		unitbet:         unitbet,
	}
}

// GetReelsLen - 取得轉輪長度
//
//	@param rtp string	RTP
//	return ReelStripLengths	轉輪長度
func (w *LineGames) GetReelsLen(rtp games.RTPs) []int {
	return w.reelLen[rtp]
}

// ContiguousReelStrips - 連續轉輪
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//
// return ReelStrips				轉輪表
func (w *LineGames) ContiguousReelStrips(rtp games.RTPs, position []int) games.ReelStrips {
	return w.reelStripsTable[rtp].ContiguousReelStrips(*w.reelStripsDef, position)
}

// ShowReelStrips - 顯示表演用轉輪表
//
//	@param rtp 	RTPs				RTP
//	@param position []int			轉輪位置
//	@param top_shift int			轉輪上方空格數
//	@param end_shift int			轉輪下方空格數
//	@return ReelStrips				轉輪表
func (w *LineGames) ShowReelStrips(rtp games.RTPs, position []int, top_shift int, end_shift int) games.ReelStrips {
	return w.reelStripsTable[rtp].ShowReelStrips(*w.reelStripsDef, position, top_shift, end_shift)
}

// CalcLineReel
//
//	@param bet decimal.Decimal	下注金額
//
// return WinDetail		獎金
func (w *LineGames) CalcLineReel(reels games.ReelStrips, bet decimal.Decimal) *WinDetail {
	result := NewWinDetail(bet)
	symb_x := games.Symbol(-1)

	// 1. 從第一條開始檢查到最後一條
	for line_idx := 0; line_idx < w.kWinLines.KWinLinesInt(); line_idx++ {
		match := 0
		matchWW := 0 //為百搭有分數之遊戲設計
		//var symb_x games.Symbol

		//法1 reel改成一列
		/*var reels_list []int
		idx := 0
		for idx_x := 0; idx_x < len(reels[0]); idx_x++ {
			for idx_y := 0; idx_y < len(reels); idx_y++ {
				reels_list[idx] = int(reels[idx_x][idx_y])
				idx++
			}
		}*/

		//法2 把中獎線一維的資料改成行列兩軸
		reel_vert := w.payLine.CalcPayLine(line_idx, 0) / len(reels)
		reel_hori := w.payLine.CalcPayLine(line_idx, 0) % len(reels)

		// 2. 該線第一個位置的獎圖
		symbol := reels[reel_hori][reel_vert]

		// Not calculate scatter
		if symbol == SF {
			continue
		}

		if symbol == WW {
			for reel_idx := 0; reel_idx < len(reels); reel_idx++ {
				reel_vert = w.payLine.CalcPayLine(line_idx, reel_idx) / len(reels)
				reel_hori = w.payLine.CalcPayLine(line_idx, reel_idx) % len(reels)
				if reels[reel_hori][reel_vert] == symbol {
					matchWW++
					match++
					if match == 5 {
						//log.Printf("line_idx: %v  match is 5", line_idx)
					}
				} else {
					symb_x = reels[reel_hori][reel_vert]
					match++
					//log.Printf("symb_x:%v ", symb_x)
					break
				}
			}

			if match < len(reels) {
				for reel_idx := match; reel_idx < len(reels); reel_idx++ {
					reel_vert = w.payLine.CalcPayLine(line_idx, reel_idx) / len(reels)
					reel_hori = w.payLine.CalcPayLine(line_idx, reel_idx) % len(reels)
					if reels[reel_hori][reel_vert] == symb_x || reels[reel_hori][reel_vert] == WW {
						match++
					} else {
						break
					}
				}
			}

			if symb_x == SF {
				continue
			}

			if match > 0 {
				m_countWW := matchWW - 1
				pointWW := w.payTable.CalcPaysTable(WW.Int(), m_countWW, 1).Mul(bet).Div(w.unitbet.Mul(decimal.NewFromInt(int64(w.kWinLines))))
				pointXW := decimal.Zero
				m_count := match - 1
				if symb_x != games.Symbol(-1) {
					pointXW = w.payTable.CalcPaysTable(symb_x.Int(), m_count, 1).Mul(bet).Div(w.unitbet.Mul(decimal.NewFromInt(int64(w.kWinLines))))
				}
				if pointWW.Cmp(pointXW) == 1 { //pointWW比point大    a.Cmp(b) == 1 只a>b
					result.Add(pointWW)
					result.Wins = append(result.Wins, NewLineWins(WW, m_countWW, line_idx, pointWW))
				} else {
					result.Add(pointXW)
					result.Wins = append(result.Wins, NewLineWins(symb_x, m_count, line_idx, pointXW))
				}
			}
		} else { // if first symbol is NOT = WW or SF
			for reel_idx := 0; reel_idx < len(reels); reel_idx++ {
				reel_vert = w.payLine.CalcPayLine(line_idx, reel_idx) / len(reels) //WinLines[line_idx][reel_idx] / len(reels)
				reel_hori = w.payLine.CalcPayLine(line_idx, reel_idx) % len(reels) //(WinLines[line_idx][reel_idx]) % len(reels)
				if reels[reel_hori][reel_vert] == symbol || reels[reel_hori][reel_vert] == WW {
					match++
				} else {
					break
				}
			}
			if match > 0 {
				m_count := match - 1
				point := w.payTable.CalcPaysTable(symbol.Int(), m_count, 1).Mul(bet).Div(w.unitbet.Mul(decimal.NewFromInt(int64(w.kWinLines))))
				result.Add(point)
				result.Wins = append(result.Wins, NewLineWins(symbol, m_count, line_idx, point))
			}
		}
	}

	return result
}

// CalcScatter - 計算 Scatter
//
//	Line 為獎圖個數
func (w *LineGames) CalcScatter(reels games.ReelStrips, bet decimal.Decimal, tar games.Symbol) *WinDetail {
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
