package games

// ReelStripList - 不同 RTP 轉輪列表
type ReelStripList map[RTPs]*ReelStrips

// ReelStripLengthTable - 不同 RTP 轉輪列表
type ReelStripLengthTable map[RTPs][]int

// NewReelStripList - 建立轉輪表列表
func NewReelStripList() ReelStripList {
	return make(ReelStripList)
}

// ReelStrips- 轉輪表結構
type ReelStrips []Reels

// Reels 轉輪結構
type Reels []Symbol

// ContiguousReelStrips - 取得轉輪表
//  @param count 獎圖數量
//  @param position 轉輪位置
func (r ReelStrips) ContiguousReelStrips(count, position []int) ReelStrips {
	rCount := len(r)
	reelStrips := make(ReelStrips, len(r))
	for i := 0; i < rCount; i++ {
		reelStrips[i] = r[i].ContiguousSymbols(count[i], position[i])
	}
	return reelStrips
}

// RepeatedReelStrips - 取得轉輪表
//  @param count 獎圖數量
//  @param position 轉輪位置
func (r ReelStrips) RepeatedReelStrips(count, position []int, repeated []bool) ReelStrips {
	rCount := len(r)
	reelStrips := make(ReelStrips, len(r))

	for i := 0; i < rCount; i++ {
		if repeated[i] {
			reelStrips[i] = r[i].RepeatedSymbols(count[i], position[i])
		} else {
			reelStrips[i] = r[i].ContiguousSymbols(count[i], position[i])
		}
	}

	return reelStrips
}

// ShowReelStrips - 取得表演用轉輪表
//  @param count 獎圖數量
//  @param position 轉輪位置
//  @param top_shift 上方轉輪偏移
//  @param end_shift 下方轉輪偏移
func (r ReelStrips) ShowReelStrips(count, position []int, top_shift, end_shift int) ReelStrips {
	rCount := len(r)
	reelStrips := make(ReelStrips, len(r))

	for i := 0; i < rCount; i++ {
		reelStrips[i] = r[i].ShowReels(count[i], position[i], top_shift, end_shift)
	}

	return reelStrips
}

// ShowRepeatedReelStrips - 取得轉輪表
//  @param count 獎圖數量
//  @param position 轉輪位置
func (r ReelStrips) ShowRepeatedReelStrips(count, position []int, repeated []bool, top_shift, end_shift int) ReelStrips {
	rCount := len(r)
	reelStrips := make(ReelStrips, len(r))

	for i := 0; i < rCount; i++ {
		if repeated[i] {
			reelStrips[i] = r[i].ShowRepeatedSymbols(count[i], position[i], top_shift, end_shift)
		} else {
			reelStrips[i] = r[i].ShowReels(count[i], position[i], top_shift, end_shift)
		}
	}

	return reelStrips
}

// Lengths - 轉輪表長度
func (r ReelStrips) Lengths() []int {
	lengths := make([]int, len(r))

	for i, reel := range r {
		lengths[i] = len(reel)
	}

	return lengths
}

// CalcSymbolMatches - 計算獎圖數量
//  @param targetSymbol 目標獎圖
//  @return int 返回幾輪中獎
//  @return int 返回獎圖數量
func (r ReelStrips) CalcSymbolMatches(targetSymbol Symbol) (int, int) {
	match := 0
	reelMatchCount := 0
	count := 0

	for _, reel := range r {
		reelMatchCount = reel.CountSymbolsInReel(targetSymbol)
		if reelMatchCount > 0 {
			count += reelMatchCount
			match++
		}
	}
	return match, count
}

// CalcSymbolsMatchFromLeft - 計算轉輪表左邊連線數量
//  @param targetSymbol 多個目標獎圖
//  @return []int 返回每輪共有多少個目標獎圖，陣列個數為中獎數量
//  @return int 返回 Way
func (r ReelStrips) CalcSymbolsMatchFromLeft(targetSymbol ...Symbol) ([]int, int) {
	match := []int{}
	multi := 1
	count := 0

	for _, reel := range r {
		count = reel.CountSymbols(targetSymbol)
		if count <= 0 {
			return match, multi
		}
		multi *= count
		match = append(match, count)
	}
	return match, multi
}

// InvertRegularXYAxis - inverts the XY axis of the ReelStrips array
func (r ReelStrips) InvertRegularXYAxis() ReelStrips {
	// Create a new ReelStrips with the inverted dimensions of the original
	reel_len := len(r)
	reel_count := len(r[0])
	reelStrips := make(ReelStrips, reel_count)
	for i := 0; i < reel_count; i++ {
		reelStrips[i] = make(Reels, reel_len)
	}

	// Invert the values from the original ReelStrips to the new one
	for i := 0; i < reel_len; i++ {
		for j := 0; j < reel_count; j++ {
			reelStrips[j][i] = r[i][j]
		}
	}
	return reelStrips
}

// // ReverseStrings - reverses the order of the strings in the ReelStrips
// func (r ReelStrips) ReverseStrings() []string {
//     // Create a new slice with the same length as the original ReelStrips
//     reversed := make([]string, len(r))

//     // Reverse the strings from the original ReelStrips to the new slice
//     for i, s := range r {
//         reversed[len(r) - i - 1] = s
//     }
//     return reversed
// }

// String - 轉輪表字串
func (r ReelStrips) Strings() []string {
	strs := make([]string, 0)
	for _, reel := range r {
		strs = append(strs, reel.Strings()...)
	}
	return strs
}

// Length - 返回轉輪長度
//  @return int 轉輪長度
func (r Reels) Length() int {
	return len(r)
}

// Contiguous - 返回連續轉輪，待優化
//  @param count 獎圖數量
//  @param position 轉輪位置
//  @param top_shift 上方轉輪顯示
//  @param end_shift 下方轉輪顯示
//  @return Reels 返回連續轉輪
func (r Reels) ContiguousSymbols(count, position int) Reels {
	contiguousReels := make(Reels, count)
	rLen := r.Length()
	position += rLen
	currentPosition := position
	for i := 0; i < (count); i++ {
		currentPosition = position + i
		contiguousReels[i] = r[currentPosition%rLen]
	}

	return contiguousReels
}

// ShowReels 顯示演出用轉輪
//  @param position 位置
//  @param topShift 上方轉輪顯示
//  @param endShift 下方轉輪顯示
//  @return Reels 返回演出用轉輪
func (r Reels) ShowReels(count, position, topShift, endShift int) Reels {

	// 計算連續轉輪的數量，及初始化轉輪
	count += topShift + endShift
	contiguousReels := make(Reels, 0)

	// 取得中間連續值，並多取最後連續值
	contiguousSymbol := r.ContiguousSymbols(count-topShift, position)

	// 	// 重新定位最上層
	rLen := r.Length()
	position += rLen
	position -= topShift

	contiguousReels = append(contiguousReels, r[position%rLen])
	contiguousReels = append(contiguousReels, contiguousSymbol...)
	// contiguousReels = append(contiguousReels, r[(position+count)%rLen])

	// 取得連續轉輪
	return contiguousReels
}

// RepeatedSymbols 返回重複獎圖轉輪
//  @param count 獎圖數量
//  @return Reels 返回重複轉輪
func (r Reels) RepeatedSymbols(count, position int) Reels {
	repeatedSymbol := r[position]

	// 創建重複轉輪
	repeatedReels := make(Reels, count)
	for i := 0; i < count; i++ {
		repeatedReels[i] = repeatedSymbol
	}

	return repeatedReels
}

// ShowRepeatedSymbols 返回重複獎圖轉輪
//  @param count 獎圖數量
//  @return Reels 返回重複轉輪
func (r Reels) ShowRepeatedSymbols(count, position, topShift, endShift int) Reels {
	repeatedSymbol := r.RepeatedSymbols(count, position)

	// 計算轉輪的數量，及初始化轉輪
	count += topShift + endShift - 1
	repeatedReels := make(Reels, 0)

	// 重新定位最上層
	rLen := r.Length()
	position += rLen
	position -= topShift

	repeatedReels = append(repeatedReels, r[position%rLen])
	repeatedReels = append(repeatedReels, repeatedSymbol...)
	repeatedReels = append(repeatedReels, r[(position+count)%rLen])

	return repeatedReels
}

// RemoveDuplicates - removes duplicate symbols from the Reels
//  @return Reels - the Reels with duplicate symbols removed
func (r Reels) RemoveDuplicates() Reels {
	// Create a map to keep track of encountered symbols
	encountered := map[Symbol]bool{}
	result := Reels{}

	// Loop through the symbols in the Reels and add them to the result slice if they haven't been encountered before
	for _, symbol := range r {
		if !encountered[symbol] {
			encountered[symbol] = true
			result = append(result, symbol)
		}
	}
	return result
}

// CountSymbolsInReel - calculates the number of times the target symbol appears in the Reels
//  @param targetSymbol the target symbol
//  @return int the number of times the target symbol appears in the Reels
func (r Reels) CountSymbolsInReel(targetSymbol Symbol) int {
	count := 0
	for _, reel := range r {
		if reel == targetSymbol {
			count++
		}
	}

	return count
}

// CountSymbols - 計算轉輪中有多少個指定的獎圖
//  @param targetSymbol 多個目標獎圖
//  @return int 返回該輪共有多少個目標獎圖
func (r Reels) CountSymbols(targetSymbol []Symbol) int {
	var count int
	for _, reel := range r {
		for _, symbol := range targetSymbol {
			if reel == symbol {
				count++
			}
		}
	}

	return count
}

// String - 返回轉輪結構字串
//  @return []string 轉輪結構字串
func (r Reels) Strings() []string {
	var reels []string
	for _, reel := range r {
		reels = append(reels, reel.String())
	}
	return reels
}

// DisplayreelStrips - 根据位置陈列奖图
func (r ReelStrips) Display(p [][]int) ReelStrips {
	rCount := len(r)
	reelStrips := make(ReelStrips, rCount)
	for i := 0; i < rCount; i++ {
		reelStrips[i] = make(Reels, len(p[i]))
		for j := 0; j < len(p[i]); j++ {
			if p[i][j] < 0 {
				reelStrips[i][j] = -1
			} else {
				reelStrips[i][j] = r[i][p[i][j]%len(r[i])]
			}
		}
	}
	return reelStrips
}

//   深拷贝
func (r ReelStrips) Deepcopy() ReelStrips {
	rr := make(ReelStrips, 0)
	for key := range r {
		C := make(Reels, 0)
		for index := range r[key] {
			C = append(C, r[key][index])
		}
		rr = append(rr, C)
	}
	return rr
}

// RemoveMark - 消除在ways上赢得的symbol
//  @param targetSymbol 多個目標獎圖
//  @param match 目标奖图对应的列（match）
//  @return ReelStrips 返回旧奖图被消除而新奖图未落下的盘面.
func (r ReelStrips) RemoveMark(targetSymbol []Symbol, match int) ReelStrips {
	//	removemark := r.Deepcopy()
	for i := 0; i < match; i++ {
		for j := 0; j < len(r[i]); j++ {
			for _, v := range targetSymbol {
				if r[i][j] == v {
					r[i][j] = -1
				}
			}
		}
	}
	return r
}
