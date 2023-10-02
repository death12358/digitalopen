package sg001_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitlab.com/gaas_module/games/random"
	weights "gitlab.com/gaas_module/games/weight"

	"github.com/shopspring/decimal"
	"gitlab.com/gaas_math/slotmachine/sg001"
	"gitlab.com/gaas_math/slotmachine/sg001/lightning"
	"gitlab.com/gaas_module/games"
)

var (
	r_default = games.Rounds{
		Id:         "1234567890",
		GameCode:   "SG001",
		Brand:      "brand_test",
		Username:   "user_test",
		Status:     games.State(0),
		Position:   games.State(0),
		Stages:     0,
		Result:     games.NewResults(),
		Currency:   "TestCoin",
		Start:      1669596839071688000,
		Fisish:     1669596839071688000,
		TotalBet:   decimal.Zero,
		TotalPoint: decimal.Zero,
	}
)

// ONE SPIN
func Test_SpinSlotSG001(t *testing.T) {
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("CalcRollRate: %s", mock_game.CalcRollRate("8", "98", decimal.NewFromFloat(4)))
	// spin
	round, err := mock_game.Spin("98", decimal.NewFromFloat(2.46), []string{"38"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %s", err.Error())
	}

	// respin
	if round.Position == 16 {
		round, err = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"1"}, *round)
		if err != nil {
			t.Errorf("Respin error: %s", err.Error())
		}
	}
	// clear round
	// round.RemoveNotStarted()
	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf(string(j_round))

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////

}

func Test_SpinRoundsSG001(t *testing.T) {
	Rounds := 10000000
	unitbet := 8
	var Totalbet float64 = 44.0
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("NGtest Spin %d times", Rounds)
	//紀錄BOUNS中獎率
	BIGFGTIMES := 0
	BGTIMES := 0
	JPTIMES := 0
	Total := decimal.Zero
	for times := 0; times < Rounds; times++ {
		// spin
		round, err := mock_game.Spin("98", decimal.NewFromFloat(Totalbet), []string{fmt.Sprint(unitbet)}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}

		//"0"是HIFG "1"是LIGHTNING
		if round.Position == 16 {
			round, err = mock_game.Spin("98", decimal.NewFromFloat(Totalbet), []string{"1"}, *round)
			if err != nil {
				t.Errorf("Respin error: %s", err.Error())
			}
			BGTIMES++
		}
		if round.Position == 33 || round.Position == 36 {
			BIGFGTIMES++
		}
		if round.Position == 129 || round.Position == 132 {
			JPTIMES++
		}

		/*j_round, _ := json.Marshal(round.Position)
		fmt.Printf("%v\n", string(j_round))*/
		Total = Total.Add(round.TotalPoint)
	}
	t.Logf("BIGFGTIMES: %v, BIGFGfreq: %v \n", BIGFGTIMES, float64(BIGFGTIMES)/float64(Rounds))
	t.Logf("JPTIMES: %v, JPfreq: %v \n", JPTIMES, float64(JPTIMES)/float64(Rounds))
	t.Logf("BGTIMES: %v, BGfreq: %v \n", BGTIMES, float64(BGTIMES)/float64(Rounds))
	t.Logf("RTP:%v \n", Total.Div(decimal.NewFromInt(int64(Rounds))).Div(decimal.NewFromFloat(Totalbet)))
}

// 關閉BONUS
// Test NG 連線分布
func Test_LinesSG001(t *testing.T) {
	Rounds := 1000000
	unitbet := 88
	// create a new slot machine

	///t.Logf("GameName: %s", mock_game.Name())
	//t.Logf("GameInfo: %s", mock_game.Info())
	//t.Logf("NGtest Spin %d times", Rounds)

	LinesRecord := [15][6]int{} //紀錄連線狀況
	for times := 0; times < Rounds; times++ {
		mock_game := games.NewGames(sg001.New())
		// spin
		round, err := mock_game.Spin("98", decimal.NewFromFloat(4400.0), []string{fmt.Sprint(unitbet)}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}

		Symbols := round.Result["0"].Symbols

		m1 := make(map[string]bool)
		//Print轉成INT後的真實滾輪//////////////////////////////////////////
		/*for i := 5; i < 10; i++ {
			fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))
		}
		fmt.Printf("\n")
		for i := 10; i < 15; i++ {
			fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))

		}
		fmt.Printf("\n")
		for i := 15; i < 20; i++ {
			fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))
		}
		fmt.Printf("\n")*/
		r1 := []string{} //紀錄第一輪的符號
		for i := 5; i <= 16; i = i + 5 {
			if !m1[Symbols[i]] {
				r1 = append(r1, Symbols[i])
			}
			m1[Symbols[i]] = true
		}

		/*fmt.Printf("Symbols[7] is %+v bool:%+v\n", Symbols[7], Symbols[7] == fmt.Sprint(0))
		fmt.Printf("Symbols[8] is %+v bool:%+v\n", Symbols[8], Symbols[8] == fmt.Sprint(0))
		fmt.Printf("Symbols[9] is %+v bool:%+v\n", Symbols[9], Symbols[9] == fmt.Sprint(0))*/ //check用
		//fmt.Printf("第一輪%+v\n", r1)
		for _, symbol := range r1 { //迴圈檢查與R1相同符號的個數
			Symbolarray := [6]int{}   //用ARRAY紀錄SYMBOL
			for i := 5; i < 10; i++ { // 檢查每一輪
				if Symbols[i] == symbol || Symbols[i] == fmt.Sprint(0) {
					Symbolarray[i-5]++
				}
				if Symbols[i+5] == symbol || Symbols[i+5] == fmt.Sprint(0) {
					Symbolarray[i-5]++
				}
				if Symbols[i+10] == symbol || Symbols[i+10] == fmt.Sprint(0) {
					Symbolarray[i-5]++
				}
			}
			///////////////////////////////
			L3 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2]                                   //3連線個數
			L4 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3]                  //4連線個數
			L5 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3] * Symbolarray[4] //5連線個數

			s := games.ToSymbol(symbol)
			//fmt.Printf("%+v\n", symbol) ///////////////////////////////////////////////
			//fmt.Printf("Symbolarray:%+v\n", Symbolarray)
			LinesRecord[s.Int()][5] += L5
			if L5 == 0 {
				LinesRecord[s.Int()][4] += L4
				if L4 == 0 {
					LinesRecord[s.Int()][3] += L3
				}

			}
		}
	}
	str := ""
	var TotolRTP decimal.Decimal
	for i, j := range LinesRecord { //i:符號編號 j:符號對應的陣列 k:k連線 l:有l條線
		str += fmt.Sprintf("%6v|%5v|", i, IntToSymbolString[i])
		for n := 0; n < 5; n++ { //pays
			str += fmt.Sprintf("%5v|", pays[i][n])
		}

		for l := 1; l < len(j); l++ { //連線數
			str += fmt.Sprintf("%12v|", decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds))))
		}
		CurRTP :=
			(decimal.NewFromInt(int64(j[1])).Mul(decimal.NewFromFloat(pays[i][0])).Add(decimal.NewFromInt(int64(j[2])).Mul(decimal.NewFromFloat(pays[i][1]))).Add(decimal.NewFromInt(int64(j[3])).Mul(decimal.NewFromFloat(pays[i][2]))).Add(decimal.NewFromInt(int64(j[4])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[5])).Mul(decimal.NewFromFloat(pays[i][4])))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds))))
		str += fmt.Sprintf("%10v", CurRTP)

		str += fmt.Sprintf("\n")
		TotolRTP = TotolRTP.Add(CurRTP)
	}
	str += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%v\n", "Totol", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", TotolRTP)

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////
	title := fmt.Sprintf("\n%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "index", "symbol", "pay-1", "pay-2", "pay-3", "pay-4", "pay-5", "hit_freq.-1", "hit_freq.-2", "hit_freq.-3", "hit_freq.-4", "hit_freq.-5", "RTP")
	title += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	/*j_string, err := json.Marshal(str)
	/*j_string, err
	 := json.Marshal(str)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("LinesRecord: %s", string(j_string))*/
	t.Logf(title + str)
}

// TEST LIGHTNING 1TIME
func Test_FlowLightning1(t *testing.T) {
	mock_game := sg001.New()
	var round *games.Rounds
	round, _ = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, r_default)
	var Total decimal.Decimal
	/*
		reels := make(games.Reels, lightning_game.GetReelDef())
		record := lightning_game.HoldNGReel(&reels, round.Result["0"])
		record.Point = record.Multiplier.Mul(bet)
		round.TotalPoint = round.TotalPoint.Add(record.Point)
		round.Result[id] = record
	*/

	//產出一個ROUND
	//round, err := mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, r_default)
	//round, err = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, *round)

	//更改ROUND中RESULT["0"]的SYMBOLS
	bgInttoStringMapping := map[int]string{
		1: "fg8",
		2: "fg18",
		3: "fg38",
		4: "fg68",
		5: "fg88",
		6: "fg888",
	}
	HowManySFWeightTable := []int{
		98480185442467,
		34525746602561,
		10112436761432,
		2616946587644,
		608852702467,
		116351250557,
		246355076852,
		2556596985,
		226093611,
		8695908}
	HowManySFObjectsTable := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	HowManySF := weights.NewGames(HowManySFWeightTable, HowManySFObjectsTable)
	dice := random.Intn(HowManySF.Sum())
	pick, _ := HowManySF.Pick(dice)

	for i := 0; i < pick; i++ {
		// 起始SF 權重選擇
		dice := random.Intn(lightning.LightningFreeGame.Sum())
		fgsymbol, _ := lightning.LightningFreeGame.Pick(dice)
		round.Result["0"].Symbols[i+5] = bgInttoStringMapping[fgsymbol]
	}
	//
	for i := pick; i < 15; i++ {
		round.Result["0"].Symbols[i+5] = "0"
	}

	fmt.Printf("\nround.Result0%+v\n", round.Result["0"])

	//將改完的ROUND代入
	r_default, err := mock_game.LightningFlow("", decimal.NewFromFloat(4.0), round)

	fmt.Printf("TotalPoint:%+v", r_default)
	for _, v := range r_default.Result {
		Total = Total.Add(v.Point)
	}
	fmt.Printf("Total: %+v ", Total)
	if err != nil {
		t.Errorf("Respin error: %s", err.Error())
	}

	// t.Logf("Respin reel: %v", reel)
	/*
		output, err := r_default.Result.Marshal()
		if err != nil {
			t.Errorf("Marshal error: %s", err.Error())
		}
		t.Logf("Respin record: %+v", r_default)
		t.Logf("Respin record: %+v", string(output))*/
}

func BenchmarkSG001(b *testing.B) {
	mock_game := games.NewGames(sg001.New())
	for i := 0; i < b.N; i++ {
		// mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"8"}, r_default)

		// spin
		round, _ := mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"8"}, r_default)

		// respin
		if round.Position == 16 {
			round, _ = mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"1"}, *round)
		}
	}
}

func BenchmarkSG001ThreadsafeParallel(b *testing.B) {
	mock_game := games.NewGames(sg001.New())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// spin
			round, _ := mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"8"}, r_default)

			// respin
			if round.Position == 16 {
				round, _ = mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"1"}, *round)
			}
		}
	})
}

// 僅計算POINTS
func Test_Points(t *testing.T) {
	// create a new slot machine
	TotalPoint := 0.0
	TotalLPoint := 0.0
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("CalcRollRate: %s", mock_game.CalcRollRate("8", "98", decimal.NewFromFloat(4)))
	JPTIME := 0
	BGTIME := 0
	BIGFGTIME := 0

	for i := 0; i < 10000; i++ {
		// spin
		round, err := mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, r_default)
		/*fmt.Printf("%v\n", round.Result["0"])
		fmt.Printf("%v\n", round.Position)
		if round.Position != 0 {
			fmt.Printf("%v\n", round.Position)
		}*/
		//fmt.Printf("%v\n", round.Position)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}
		if round.Position == 132 || round.Position == 129 {
			JPTIME++
		}
		if round.Position == games.FreeGame {
			BIGFGTIME++
		}
		// respin
		if round.Position == 16 {
			round, err = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"1"}, *round)
			if err != nil {
				t.Errorf("Respin error: %s", err.Error())
			}
			if round.Position == games.BonusGame2 {
				BGTIME++
			}
			fmt.Printf("%v\n", round.TotalPoint)
		}

		fmt.Printf("%v\n", round.Position)
		a, _ := round.TotalPoint.Float64()
		//PointRecord[a]++

		TotalPoint += a
		// clear round
		// round.RemoveNotStarted()

		/*
			if err != nil {
				t.Errorf("Marshal error: %s", err.Error())
			}
			t.Logf("Round: %s", string(j_round))

			// t.Logf("Round: %v", round)*/
	}
	//for k, v := range PointRecord {
	//	fmt.Printf("%+v, %+v\n", k, v)
	//}

	fmt.Printf("BonusTimes:JP %v   BIG %v    BG %v\n", JPTIME, BIGFGTIME, BGTIME)
	fmt.Printf("TotalPoint:%v\n", TotalPoint/10000/4)
	fmt.Printf("TotalLPoint:%v\n", TotalLPoint/100000/4)
}

/*
func Test_Lightning(t *testing.T) {
	mock_game := lightning.NewLightningGames(lightning.LightningGame, lightning.LightningFreeGame, 15)
	reel := make(games.Reels, 15)
	//////////////////////////////////////////////////////////////
	//依照EXCEL計算結果 隨機決定Lightning起始的SF(權重可調整精確的小數位)
	HowManySFWeightTable := []int{9848, 03453, 1011, 262, 61, 12, 25, 0, 0, 0}
	HowManySFObjectsTable := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	HowManySF := weights.NewGames(HowManySFWeightTable, HowManySFObjectsTable)
	dice := random.Intn(HowManySF.Sum())
	pick, _ := HowManySF.Pick(dice)

	for i := 0; i < pick; i++ {
		// SF 權重選擇
		dice := random.Intn(lightning.LightningFreeGame.Sum())
		pick, _ := lightning.LightningFreeGame.Pick(dice)
		reel[i] = games.Symbol(pick)
	}

	record := &games.Records{
		Brand:      "brand_test",
		Username:   "user_test",
		Case:       games.None,
		Stages:     3, // default 3
		Multiplier: decimal.Zero,
		Bet:        decimal.Zero,
		Point:      decimal.Zero,
	}
	var err error

	_, record, err = mock_game.Respin(&reel, record, "1", decimal.NewFromFloat(1.0))
	if err != nil {
		t.Errorf("Respin error: %s", err.Error())
	}
	// t.Logf("Respin reel: %v", reel)
	t.Logf("Respin record: %+v", record)
}*/

/*

func Test_SlotSG001Lightning(t *testing.T) {
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("CalcRollRate: %s", mock_game.CalcRollRate("8", "98", decimal.NewFromFloat(4)))

	// spin
	round, err := mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %s", err.Error())
	}

	// respin
	if round.Position == 16 {
		round, err = mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"1"}, *round)
		if err != nil {
			t.Errorf("Respin error: %s", err.Error())
		}
	}

	// clear round
	// round.RemoveNotStarted()

	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("Round: %s", string(j_round))

	// t.Logf("Round: %v", round)
}
*/

// 純計算 比對LIGHTNING用
func Test_onegai(t *testing.T) {
	SymbolNumMap := make(map[int]int)
	TotalPoint := 0
	Total := 0
	//TotalArray := make([]int, 500)
	rounds := 1000000
	BoardNumMap := make(map[int]int)
	T := 100.0 + 200.0 + 1200.0 + 600.0 + 200.0 + 10.0

	THEORY := []float64{100.0 / T, 200.0 / T, 1200.0 / T, 600.0 / T, 200.0 / T, 10.0 / T}

	for time := 0; time < rounds; time++ {
		TotalBet := 1
		TotalGoldPoint := 0
		curGreenNumber := 0
		TotalGreenPoint := 0

		board := make([]int, 15, 15)

		HowManySFWeightTable := []int{
			98480185442467,
			34525746602561,
			10112436761432,
			2616946587644,
			608852702467,
			116351250557,
			246355076852,
			2556596985,
			226093611,
			8695908}
		HowManySFObjectsTable := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		HowManySF := weights.NewGames(HowManySFWeightTable, HowManySFObjectsTable)
		dice := random.Intn(HowManySF.Sum())
		pick, _ := HowManySF.Pick(dice)

		LightningFGWeightTable := []int{100, 200, 1200, 600, 200, 10}
		LightningFGObjectTable := []int{8, 18, 38, 68, 88, 888}

		LightningFG := weights.NewGames(
			LightningFGWeightTable,
			LightningFGObjectTable,
		)

		for i := 0; i < pick; i++ {
			// 起始SF 權重選擇
			dice := random.Intn(LightningFG.Sum())
			fgsymbol, _ := LightningFG.Pick(dice)
			board[i] = fgsymbol
			TotalGoldPoint += fgsymbol
		}
		//
		for i := pick; i < 15; i++ {
			board[i] = 0
		}

		//fmt.Printf("%v\n", board)

		stage := 3
		bg_LightningWeightTable := []int{90, 8, 2}
		bg_LightningObjectTable := []int{0, 100, 1000}

		LightningGame := weights.NewGames(
			bg_LightningWeightTable,
			bg_LightningObjectTable,
		)
		//fmt.Printf("stage:%v: board:%v\n ", stage, board)
		A := 0
		for stage > 0 {
			for i := 0; i < 15; i++ {
				if board[i] == 0 {
					dice := random.Intn(LightningGame.Sum())
					FGColor, _ := LightningGame.Pick(dice)
					if FGColor == 100 {
						dice := random.Intn(LightningFG.Sum())
						fgsymbol, _ := LightningFG.Pick(dice)
						board[i] = fgsymbol
						TotalGoldPoint += fgsymbol
						SymbolNumMap[fgsymbol]++
						stage = 4
						A++
					}
					if FGColor == 1000 {
						curGreenNumber++
						board[i] = 1000
						stage = 4
						SymbolNumMap[1000]++
						A++
					}
				}
			}

			//fmt.Printf("stage:%v: board:%v\n ", stage, board)

			TotalGreenPoint += TotalGoldPoint * curGreenNumber
			//fmt.Printf("GOLD:%v: GREEN:%v\n ", TotalGoldPoint, TotalGreenPoint)
			stage--
			curGreenNumber = 0
			TotalPoint = (TotalGoldPoint + TotalGreenPoint) * TotalBet
		}
		BoardNumMap[A]++
		//fmt.Printf("TotalPoint:%v\n ", TotalPoint)
		Total += TotalPoint
		//TotalArray[TotalPoint/50]++
		//fmt.Printf("\n")
	}
	fmt.Printf("THEORY:%+v\n", THEORY)
	A := 0
	for _, v := range SymbolNumMap {
		A += v
	}
	fmt.Printf("\n%v\n", A)
	SymbolNumMap2 := make(map[int]float64)
	for k, v := range SymbolNumMap {
		SymbolNumMap2[k] = float64(v) / float64(A)
	}
	fmt.Printf("SymbolNumMap:%+v\n", SymbolNumMap)
	fmt.Printf("SymbolNumMap2:%+v\n", SymbolNumMap2)

	fmt.Printf("BoardNumMap:%+v\n", BoardNumMap)
	fmt.Printf("Total:%v\n", float64(Total)/float64(rounds))
	//fmt.Printf("TotalArray:%v\n", TotalArray)
	// t.Logf("Respin reel: %v", reel)
	/*
		output, err := r_default.Result.Marshal()
		if err != nil {
			t.Errorf("Marshal error: %s", err.Error())
		}
		t.Logf("Respin record: %+v", r_default)
		t.Logf("Respin record: %+v", string(output))*/
}

/*
func Test_BigFGFlow(t *testing.T) {
	Total := decimal.Zero
	var round *games.Rounds
	for i := 0; i < 1; i++ {
		mock_game := sg001.New()

		round, _ = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"8"}, r_default)
		//for i := 0; i < 100; i++ {
		round.TotalPoint = decimal.Zero
		round.Result["0"].Symbols[6] = "fg38"
		round.Result["0"].Symbols[7] = "fg38"
		round.Result["0"].Symbols[8] = "fg38"
		mock_game.LightningFlow("98", decimal.NewFromInt(4), round)
		fmt.Printf("mock_game.BigFGFlow: %+v\n", mock_game.BigFGFlow("98", decimal.NewFromInt(4), round))
		fmt.Printf("TotalPoint: %+v\n", round.TotalPoint)
		Total = Total.Add(round.TotalPoint)
		//fmt.Printf("round.TotalPoint: %+v\n", round.TotalPoint)

		//}
	}
	fmt.Printf("EXP: %+v\n", Total.Div(decimal.NewFromInt(1)))

}
*/
//func Test_jp
/////////////////////////////////////////////////////////////////////////////////////////////////

var IntToSymbolString map[int]string = map[int]string{
	0:  "WW",
	1:  "H1",
	2:  "H2",
	3:  "H3",
	4:  "H4",
	5:  "H5",
	6:  "LA",
	7:  "LK",
	8:  "LQ",
	9:  "LJ",
	10: "LT",
	11: "LN",
	12: "SE",
	13: "SF",
	14: "SB",
}

/*
	var pays = [...][5]float64{
		{0.0, 0.0, 0.0, 0.0, 0.0},   //WW
		{0.0, 0.0, 5.0, 10.0, 20.0}, //H1
		{0.0, 0.0, 5.0, 10.0, 20.0}, //H2
		{0.0, 0.0, 5.0, 10.0, 20.0}, //H3
		{0.0, 0.0, 5.0, 10.0, 20.0}, //H4
		{0.0, 0.0, 5.0, 10.0, 20.0}, //H5
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LA
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LK
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LQ
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LJ
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LT
		{0.0, 0.0, 5.0, 8.0, 10.0},  //LN
		{0.0, 0.0, 0.0, 0.0, 0.0},   //SE
		{0.0, 0.0, 0.0, 0.0, 0.0},   //SF
		{0.0, 0.0, 0.0, 0.0, 0.0},   //SB
	}
*/
var pays = [...][5]float64{
	{0.0, 0.0, 0.0, 0.0, 0.0}, //SE

	{20.0, 60.0, 250.0},
	{15.0, 50.0, 200.0},
	{15.0, 25.0, 150.0},
	{10.0, 20.0, 100.0},
	{5.0, 10.0, 20.0},
	{5.0, 8.0, 10.0},
	{5.0, 8.0, 10.0},
	{5.0, 8.0, 10.0},
	{5.0, 8.0, 10.0},
	{5.0, 8.0, 10.0},
	{5.0, 8.0, 10.0},
	{0.0, 0.0, 0.0, 0.0, 0.0}, //SE
	{0.0, 0.0, 0.0, 0.0, 0.0}, //SF
	{0.0, 0.0, 0.0, 0.0, 0.0}, //SB
}

// 調整sg001.go的pos到沒中獎的盤面 以方便令出始盤面
func Test_SlotSG001LightningRounds(t *testing.T) {
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	bet := decimal.NewFromFloat(4400)
	unitbet := "88"
	t.Logf("CalcRollRate: %s", mock_game.CalcRollRate("8", "98", bet))
	t.Logf("jackpot Base pool: %v", mock_game.BaseJackpot("8", "98", bet))
	Rounds := 1000000

	Total := decimal.Zero
	//用ARRAY紀錄出現的次數
	SymbolRecord := [10]int{}
	NumberRecord := [17]int{}
	//StageRecord := [100]int{}
	TotalNumberSubInitial := 0
	for i := 0; i < Rounds; i++ {
		// spin
		round, err := mock_game.Spin("98", bet, []string{unitbet}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}

		//更改ROUND中RESULT["0"]的SYMBOLS
		bgInttoStringMapping := map[int]string{
			1: "fg8",
			2: "fg18",
			3: "fg38",
			4: "fg68",
			5: "fg88",
			6: "fg888",
		}
		HowManySFWeightTable := []int{
			98480185442467,
			34525746602561,
			10112436761432,
			2616946587644,
			608852702467,
			116351250557,
			246355076852,
			2556596985,
			226093611,
			8695908}
		HowManySFObjectsTable := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		HowManySF := weights.NewGames(HowManySFWeightTable, HowManySFObjectsTable)
		dice := random.Intn(HowManySF.Sum())
		pick, _ := HowManySF.Pick(dice)

		for i := 0; i < pick; i++ {
			// 起始SF 權重選擇
			dice := random.Intn(lightning.LightningFreeGame.Sum())
			fgsymbol, _ := lightning.LightningFreeGame.Pick(dice)
			round.Result["0"].Symbols[i+5] = bgInttoStringMapping[fgsymbol]
		}
		//
		for i := pick; i < 15; i++ {
			round.Result["0"].Symbols[i+5] = "0"
		}
		TotalNumberSubInitial -= pick
		//fmt.Printf("\nround.Result0%+v\n", round.Result["0"])
		round.Position = 16
		// respin
		if round.Position == 16 {
			round, err = mock_game.Spin("98", bet, []string{"1"}, *round)
			if err != nil {
				t.Errorf("Respin error: %s", err.Error())
			}
		}
		//j_r, _ := json.Marshal(round)
		//fmt.Printf("\n%+v\n", string(j_r))
		FinalResult := round.Result[fmt.Sprint(round.Stages-1)]

		Total = Total.Add(round.TotalPoint)
		Total = Total.Sub(round.Result["0"].Point)

		SymbolRecord[7] = 0
		//用ARRAY紀錄出現的次數	0:fg8, 1:fg18, 2:fg38, 3:fg68, 4:fg88, 5:fg888 6:bg 7:總次數
		//fmt.Printf("%+v\n", FinalResult.Symbols)

		for _, v := range FinalResult.Symbols {
			if v == "fg8" {
				SymbolRecord[0]++
				SymbolRecord[7]++
			}
			if v == "fg18" {
				SymbolRecord[1]++
				SymbolRecord[7]++
			}
			if v == "fg38" {
				SymbolRecord[2]++
				SymbolRecord[7]++
			}
			if v == "fg68" {
				SymbolRecord[3]++
				SymbolRecord[7]++
			}
			if v == "fg88" {
				SymbolRecord[4]++
				SymbolRecord[7]++
			}
			if v == "fg888" {
				SymbolRecord[5]++
				SymbolRecord[7]++
			}
			if v == "bg" {
				SymbolRecord[6]++
				SymbolRecord[7]++
			}
		}
		NumberRecord[SymbolRecord[7]]++
		TotalNumberSubInitial += SymbolRecord[7]
	}

	t.Logf("Lightningtest %d times", Rounds)
	//fmt.Printf("Total %+v \n:", Total)
	//fmt.Printf("Total: %+v ", Total)
	fmt.Printf("平均倍數: %+v \n", Total.Div(decimal.NewFromInt(int64(Rounds))).Div(bet))
	//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//圖案分布
	title := fmt.Sprintf("\n%s|%s|%s|%s|%s|%s|%s|%s\n", "圖案", "fg8", "fg18", "fg38", "fg68", "fg88", "fg888", "fgbg")
	title += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s\n", "-", "-", "-", "-", "-", "-", "-", "-")
	str := ""
	str += fmt.Sprintf("%v次|%v|%v|%v|%v|%v|%v|%v\n", Rounds, SymbolRecord[0], SymbolRecord[1], SymbolRecord[2], SymbolRecord[3], SymbolRecord[4], SymbolRecord[5], SymbolRecord[6])
	str += fmt.Sprintf("%s|%v|%v|%v|%v|%v|%v|%v\n", "平均出現個數", decimal.NewFromInt(int64(SymbolRecord[0])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[1])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[2])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[3])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[4])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[5])).Div(decimal.NewFromInt(int64(Rounds))), decimal.NewFromInt(int64(SymbolRecord[6])).Div(decimal.NewFromInt(int64(Rounds))))
	A := SymbolRecord[0] + (SymbolRecord[1]) + (SymbolRecord[2]) + (SymbolRecord[3]) + (SymbolRecord[4]) + (SymbolRecord[5])
	str += fmt.Sprintf("%s|%v|%v|%v|%v|%v|%v|%v\n", "占比", decimal.NewFromInt(int64(SymbolRecord[0])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[1])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[2])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[3])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[4])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[5])).Div(decimal.NewFromInt(int64(A))), decimal.NewFromInt(int64(SymbolRecord[6])).Div(decimal.NewFromInt(int64(TotalNumberSubInitial))))
	str += fmt.Sprintf("\n")

	//最終盤面個數
	str += fmt.Sprintf("%s|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|\n", "最終盤面個數:", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	str += fmt.Sprintf("%s|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|%v|\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	str += fmt.Sprintf("%s", "次數:|")
	for i := 0; i < 16; i++ {
		str += fmt.Sprintf("%v|", NumberRecord[i])
	}
	str += fmt.Sprintf("\n")
	str += fmt.Sprintf("%s", "占比:|")
	for i := 0; i < 16; i++ {
		str += fmt.Sprintf("%v|", decimal.NewFromInt(int64(NumberRecord[i])).Div(decimal.NewFromInt(int64(Rounds))))
	}
	str += fmt.Sprintf("\n")
	/*
		str += fmt.Sprintf("\n%10s\n", "共轉了幾次:|")
		str += fmt.Sprintf("%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v|%10v\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")

		for r := 0; r < 5; r++ {
			for c := 0; c < 10; c++ {
				str += fmt.Sprintf("%1v|", r*20+c)
			}
			str += fmt.Sprintf("\n")
			for c := 0; c < 20; c++ {
				str += fmt.Sprintf("%1v|", decimal.NewFromInt(int64(StageRecord[r*20+c])))
			}
			str += fmt.Sprintf("\n")
		}*/
	t.Logf(title + str)

}

// HIFG 修改起始盤面再測
func Test_HiFGLinesSG001(t *testing.T) {
	Rounds := 1
	unitbet := 18
	Totalbet := decimal.NewFromFloat(4.0)
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("HIFGtest Spin %d times", Rounds)

	LinesRecord := [7][6]int{} //紀錄連線狀況
	for times := 0; times < Rounds; times++ {
		// spin
		round, err := mock_game.Spin("98", Totalbet, []string{fmt.Sprint(unitbet)}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}
		// respin
		if round.Position == 16 {
			round, err = mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{"0"}, *round)
			if err != nil {
				t.Errorf("Respin error: %s", err.Error())
			}
		}

		j_r, _ := json.Marshal(round)
		t.Logf("%+v\n", string(j_r))
		for _, res := range round.Result {
			Symbols := res.Symbols

			m1 := make(map[string]bool)
			//Print轉成INT後的真實滾輪//////////////////////////////////////////
			/*for i := 5; i < 10; i++ {
				fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))
			}
			fmt.Printf("\n")
			for i := 10; i < 15; i++ {
				fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))

			}
			fmt.Printf("\n")
			for i := 15; i < 20; i++ {
				fmt.Printf("%v ", int(games.ToSymbol(Symbols[i])))
			}
			fmt.Printf("\n")*/
			r1 := []string{} //紀錄第一輪的符號
			for i := 5; i <= 16; i = i + 5 {
				if !m1[Symbols[i]] {
					r1 = append(r1, Symbols[i])
				}
				m1[Symbols[i]] = true
			}

			/*fmt.Printf("Symbols[7] is %+v bool:%+v\n", Symbols[7], Symbols[7] == fmt.Sprint(0))
			fmt.Printf("Symbols[8] is %+v bool:%+v\n", Symbols[8], Symbols[8] == fmt.Sprint(0))
			fmt.Printf("Symbols[9] is %+v bool:%+v\n", Symbols[9], Symbols[9] == fmt.Sprint(0))*/ //check用
			//fmt.Printf("第一輪%+v\n", r1)
			for _, symbol := range r1 { //迴圈檢查與R1相同符號的個數
				Symbolarray := [6]int{}   //用ARRAY紀錄SYMBOL
				for i := 5; i < 10; i++ { // 檢查每一輪
					if Symbols[i] == symbol || Symbols[i] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
					if Symbols[i+5] == symbol || Symbols[i+5] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
					if Symbols[i+10] == symbol || Symbols[i+10] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
				}
				///////////////////////////////
				L3 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2]                                   //3連線個數
				L4 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3]                  //4連線個數
				L5 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3] * Symbolarray[4] //5連線個數

				s := games.ToSymbol(symbol)
				//fmt.Printf("%+v\n", symbol) ///////////////////////////////////////////////
				//fmt.Printf("Symbolarray:%+v\n", Symbolarray)
				LinesRecord[s.Int()][5] += L5
				if L5 == 0 {
					LinesRecord[s.Int()][4] += L4
					if L4 == 0 {
						LinesRecord[s.Int()][3] += L3
					}

				}
			}
		}
	}
	Rounds = Rounds * 8
	str := ""
	var TotolRTP decimal.Decimal
	for i, j := range LinesRecord { //i:符號編號 j:符號對應的陣列 k:k連線 l:有l條線
		str += fmt.Sprintf("%6v|%5v|", i, IntToSymbolString[i])
		for n := 0; n < 5; n++ { //pays
			str += fmt.Sprintf("%5v|", pays[i][n])
		}

		for l := 1; l < len(j); l++ { //連線頻率
			str += fmt.Sprintf("%12v|", decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds))))
		}
		CurRTP :=
			(decimal.NewFromInt(int64(j[1])).Mul(decimal.NewFromFloat(pays[i][0])).Add(decimal.NewFromInt(int64(j[2])).Mul(decimal.NewFromFloat(pays[i][1]))).Add(decimal.NewFromInt(int64(j[3])).Mul(decimal.NewFromFloat(pays[i][2]))).Add(decimal.NewFromInt(int64(j[4])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[5])).Mul(decimal.NewFromFloat(pays[i][4])))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds))))
		str += fmt.Sprintf("%10v", CurRTP)

		str += fmt.Sprintf("\n")
		TotolRTP = TotolRTP.Add(CurRTP)
	}
	str += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%v\n", "Totol", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", TotolRTP)

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////
	title := fmt.Sprintf("\n%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "index", "symbol", "pay-1", "pay-2", "pay-3", "pay-4", "pay-5", "hit_freq.-1", "hit_freq.-2", "hit_freq.-3", "hit_freq.-4", "hit_freq.-5", "RTP")
	title += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	j_string, err := json.Marshal(str)

	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("LinesRecord: %s", string(j_string))
	t.Logf(title + str)
}

func Test_BigFGSG001(t *testing.T) {
	Rounds := 1000000
	unitbet := 18
	// create a new slot machine
	mock_game := games.NewGames(sg001.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("BigFGtest Spin %d times", Rounds)

	LinesRecord := [15][6]int{} //紀錄連線狀況
	for times := 0; times < Rounds; times++ {
		// spin
		round, err := mock_game.Spin("98", decimal.NewFromFloat(4.0), []string{fmt.Sprint(unitbet)}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}

		/*j_r, _ := json.Marshal(round)
		t.Logf("%+v\n", string(j_r))*/
		for _, res := range round.Result {
			Symbols := res.Symbols
			m1 := make(map[string]bool)

			r1 := []string{} //紀錄第一輪的符號
			for i := 5; i <= 16; i = i + 5 {
				if !m1[Symbols[i]] {
					r1 = append(r1, Symbols[i])
				}
				m1[Symbols[i]] = true
			}

			for _, symbol := range r1 { //迴圈檢查與R1相同符號的個數
				Symbolarray := [6]int{}   //用ARRAY紀錄SYMBOL
				for i := 5; i < 10; i++ { // 檢查每一輪
					if Symbols[i] == symbol || Symbols[i] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
					if Symbols[i+5] == symbol || Symbols[i+5] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
					if Symbols[i+10] == symbol || Symbols[i+10] == fmt.Sprint(0) {
						Symbolarray[i-5]++
					}
				}
				///////////////////////////////
				L3 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2]                                   //3連線個數
				L4 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3]                  //4連線個數
				L5 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3] * Symbolarray[4] //5連線個數

				s := games.ToSymbol(symbol)
				//fmt.Printf("%+v\n", symbol) ///////////////////////////////////////////////
				//fmt.Printf("Symbolarray:%+v\n", Symbolarray)
				LinesRecord[s.Int()][5] += L5
				if L5 == 0 {
					LinesRecord[s.Int()][4] += L4
					if L4 == 0 {
						LinesRecord[s.Int()][3] += L3
					}

				}
			}
		}
	}
	Rounds = Rounds * 6
	str := ""
	var TotolRTP decimal.Decimal
	for i, j := range LinesRecord { //i:符號編號 j:符號對應的陣列 k:k連線 l:有l條線
		str += fmt.Sprintf("%v|%v|", i, IntToSymbolString[i])
		for n := 0; n < 5; n++ { //pays
			str += fmt.Sprintf("%v|", pays[i][n])
		}

		for l := 1; l < len(j); l++ { //連線數
			str += fmt.Sprintf("%v|", decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds))))
		}
		CurRTP :=
			(decimal.NewFromInt(int64(j[1])).Mul(decimal.NewFromFloat(pays[i][0])).Add(decimal.NewFromInt(int64(j[2])).Mul(decimal.NewFromFloat(pays[i][1]))).Add(decimal.NewFromInt(int64(j[3])).Mul(decimal.NewFromFloat(pays[i][2]))).Add(decimal.NewFromInt(int64(j[4])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[5])).Mul(decimal.NewFromFloat(pays[i][4])))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds))))
		str += fmt.Sprintf("%v", CurRTP)

		str += fmt.Sprintf("\n")
		TotolRTP = TotolRTP.Add(CurRTP)
	}
	str += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%v\n", "Totol", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", TotolRTP)

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////
	title := fmt.Sprintf("\n%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "index", "symbol", "pay-1", "pay-2", "pay-3", "pay-4", "pay-5", "hit_freq.-1", "hit_freq.-2", "hit_freq.-3", "hit_freq.-4", "hit_freq.-5", "RTP")
	title += fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	/*j_string, err := json.Marshal(str)
	/*j_string, err
	 := json.Marshal(str)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("LinesRecord: %s", string(j_string))*/
	t.Logf(title + str)
}
