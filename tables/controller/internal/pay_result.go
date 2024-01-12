package internal

// PayDetail 賠付詳情
type PayDetail struct {
	Symbol   string //圖騰
	NumOrWay int32  //第幾線或Way
	Lines    int32  //幾連線
	Rate     int32  //賠付
}

// NewPayDetail 建立賠付結項資料
func NewPayDetail(symbol string, numOrWay, lines, rate int32) *PayDetail {
	return &PayDetail{
		Symbol:   symbol,
		NumOrWay: numOrWay,
		Lines:    lines,
		Rate:     rate,
	}
}

// PayResult 賠付結果
type PayResult struct {
	Reels       [][]string
	positionMap map[Position]bool
	totalRate   int32
	details     []*PayDetail
}

// Position 位置結構
type Position struct {
	x int
	y int
}

// NewPosition 建立position資訊
func NewPosition(x, y int) Position {
	return Position{
		x: x,
		y: y,
	}
}

func (p *PayResult) SetPosition(x, y int) {
	pos := NewPosition(x, y)
	if _, ok := p.positionMap[pos]; ok {
		p.positionMap[pos] = true
	}
}

// NewPayResult new一個賠付結果的結構
func NewPayResult(inReels [][]string) *PayResult {

	rslt := PayResult{
		Reels:       inReels,
		positionMap: make(map[Position]bool),
		totalRate:   0,
		details:     []*PayDetail{},
	}

	for x := 0; x < len(inReels); x++ {
		for y := 0; y < len(inReels[len(inReels)-1]); y++ {
			pos := NewPosition(x, y)
			rslt.positionMap[pos] = false
		}
	}

	return &rslt
}

// RecordByWayGame 記錄way game位置
func (p *PayResult) RecordByWayGame(inDetail *PayDetail) {
	if inDetail.Rate != 0 {
		p.details = append(p.details, inDetail)
		// p.totalRate += (inDetail.Rate * inDetail.NumOrWay)
		p.totalRate += inDetail.Rate
		for x := 0; x < int(inDetail.Lines); x++ {
			for y := 0; y < len(p.Reels[x]); y++ {
				if p.Reels[x][y] == inDetail.Symbol || p.Reels[x][y] == Wild {
					p.SetPosition(x, y)
				}
			}
		}
	}
}

// RecordByLineGame 記錄line game位置
func (p *PayResult) RecordByLineGame(inLineSetting []int, inDetail *PayDetail) {
	if inDetail.Rate != 0 {
		p.details = append(p.details, inDetail)
		p.totalRate += inDetail.Rate
		for x, y := range inLineSetting {
			if x < int(inDetail.Lines) {
				p.SetPosition(x, y) //記錄下來有達成連線的位置
			}
		}
	}
}

// GetPayDetail 取得賠付詳情
func (p *PayResult) GetPayDetail() []*PayDetail {
	return p.details
}

// CalcWin 計算贏分
func (p *PayResult) CalcWin(inBet, inBase int32) int32 {
	return CalcWinScore(inBet, inBase, p.totalRate)
}
