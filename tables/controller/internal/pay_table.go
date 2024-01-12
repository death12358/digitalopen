package internal

import (
	"log"

	"github.com/death12358/digitalopen/tables/tools"
)

type PayTable struct {
	stand       int
	symbolList  []string
	lineSetting [][]int
	table       map[string][]int
}

// NewPayTable 新創一個pay table
func NewPayTable(calcStand int) *PayTable {
	return &PayTable{
		stand:       calcStand,
		symbolList:  []string{},
		lineSetting: [][]int{},
		table:       make(map[string][]int),
	}
}

// InitPayInfo 初始化
func (p *PayTable) InitPayInfo(inData DataMap) {
	p.setSymbolList(inData)
	// p.setPayTable(inData)
}

// // ConvPBPayData 轉為protobuf 格式
// func (p *PayTable) ConvPBPayData() *pb_game.PayData {
// 	var payData pb_game.PayData
// 	payData.PayMap = make(map[string]*pb_game.Pay)
// 	for symbol, payInfo := range p.table {
// 		var pay pb_game.Pay
// 		pay.PayValue = make([]int32, len(payInfo))
// 		for i, v := range payInfo {
// 			pay.PayValue[i] = int32(v)
// 		}
// 		payData.PayMap[symbol] = &pay
// 	}
// 	return &payData
// }

// setSymbolList 設定賠付圖案
func (p *PayTable) setSymbolList(inData DataMap) bool {
	var ok bool
	p.symbolList, ok = inData.GetDataByKey("SYMBOL")
	return ok
}

// // setPayTable 設定Pay Table
// func (p *PayTable) setPayTable(inData DataMap) {
// 	for key := range inData {
// 		if valSlice, ok := inData.ParseRowDataToInt(key); ok {
// 			p.table[key] = valSlice
// 		}
// 	}
// }

// func (pay *PayTable) setPayLine(inData DataMap) error {

// 	for i := 1; i < len(inData); i++ {
// 		lineKey := strconv.Itoa(i)
// 		lineSetting, ok := inData.ParseRowDataToInt(lineKey)
// 		if ok {
// 			pay.lineSetting = append(pay.lineSetting, lineSetting)
// 		} else {
// 			return fmt.Errorf("line no:%v, 不存在", i)
// 		}
// 	}
// 	return nil
// }

// Calculate 計算出賠付結果
func (p *PayTable) Calculate(inReels [][]string, base int32) *PayResult {

	payRslt := NewPayResult(inReels)

	switch GameStand(p.stand) {
	case Line:
		p.calcLineGame(payRslt, base)
	case Way:
		p.calcWayGame(payRslt)
	default:
		return nil
	}
	return payRslt
}

// 計算line賠付
func (p *PayTable) calcLineGame(inPayResult *PayResult, base int32) {
	//檢查線圖
	for num, setting := range p.lineSetting[:base] {
		var (
			symbolLines int
			wildLines   int
			checkSymbol string
		)
		checkSymbol = Wild //先從是不是wild開始
		for col, row := range setting {
			inReelSymbol := inPayResult.Reels[col][row]
			if inReelSymbol == Wild && inReelSymbol == checkSymbol {
				symbolLines++
				wildLines++
				continue
			} else if inReelSymbol != Wild && checkSymbol == Wild {
				checkSymbol = inReelSymbol //wild要代表的圖
			}

			if inReelSymbol == Wild || inReelSymbol == checkSymbol {
				symbolLines++
				continue
			}
			break
		}

		//一般圖騰的連線與wild顆數的連線倍付比較(取高) e.g wild達成4線連 vs 圖騰的5連線
		symbolP := p.CreatePayDetail(checkSymbol, int32(num+1), int32(symbolLines))
		wildP := p.CreatePayDetail(Wild, int32(num+1), int32(wildLines))
		detail := &PayDetail{}
		if symbolP.Rate != 0 || wildP.Rate != 0 { //賠付
			if symbolP.Rate > wildP.Rate {
				detail = symbolP
			} else {
				detail = wildP
			}
			inPayResult.RecordByLineGame(setting, detail)
		}

	}

}

// getPayRate 取出賠付
func (p *PayTable) getPayRate(inSymbol string, inLine int32) int32 {
	if inSymbol == "" {
		log.Println("GetPayRate input symbol issue")
		return 0
	}
	payVal := p.table[inSymbol]
	if inLine-1 < 0 || inLine-1 > int32(len(payVal)) {
		return 0
	}
	return int32(p.table[inSymbol][inLine-1])
}

// CreatePayDetail 建立賠付詳情
func (p *PayTable) CreatePayDetail(symbol string, numOrWay, lines int32) *PayDetail {
	rate := p.getPayRate(symbol, lines)
	if p.stand == int(Way) && rate != 0 {
		rate *= numOrWay
	}
	return NewPayDetail(symbol, numOrWay, lines, rate)
}

// 計算way賠付
func (p *PayTable) calcWayGame(inPayResult *PayResult) {
	if len(inPayResult.Reels) <= 0 {
		log.Println("calcWayGame in reels data issue")
		return
	}
	//第一輪去重複
	tempSlice := make([]string, len(inPayResult.Reels[0]))
	copy(tempSlice, inPayResult.Reels[0])
	uniqueSlice := tools.RemoveDuplicates(tempSlice).([]string)

	for _, symbol := range uniqueSlice {
		var (
			lines int
			ways  int = 1
		)
		if symbol == Scatter { //scatter跳過不判斷
			continue
		}
		for col := 0; col < len(inPayResult.Reels); col++ {
			symbolCount := tools.CalcSliceElementCount(inPayResult.Reels[col], symbol)
			wildCount := tools.CalcSliceElementCount(inPayResult.Reels[col], Wild)
			sum := symbolCount + wildCount
			if sum == 0 {
				break
			}
			lines++
			ways *= sum
		}

		detail := p.CreatePayDetail(symbol, int32(ways), int32(lines))

		inPayResult.RecordByWayGame(detail)

	}
}

// CalcWinScore 計算贏分
func CalcWinScore(bet, base, pay int32) int32 {
	if base == 0 {
		return 0
	}
	return bet / base * pay
}
