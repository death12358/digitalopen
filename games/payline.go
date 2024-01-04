package games

// kWinLines - 幾條線
type KWinLines int

// PayLine - 賠付線
type PayLine [][]int

// CalcPayLine - 計算獎金
func (p PayLine) CalcPayLine(line_idx int, reel_idx int) int {
	return p[line_idx][reel_idx]
}

// KWinLinesInt - 計算獎金
func (k KWinLines) KWinLinesInt() int {
	return int(k)
}
