package sg006_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/sg006"
	"github.com/shopspring/decimal"
)

var (
	r_default = games.Rounds{
		Id:         "1234567890",
		GameCode:   "SG006",
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

// 一次Spin
func Test_SpinSlotSG006(t *testing.T) {
	// create a new slot machine
	mock_game := games.NewGames(sg006.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())

	// spin
	round, err := mock_game.Spin("98", decimal.NewFromFloat(10.0), []string{"100"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %s", err.Error())
	}

	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("Round: %s", string(j_round))
	t.Logf("Round.Position: %v", round.Position)
	t.Logf("Round.Status: %v", round.Status)

	t.Logf("Round.Symbols: \n%v\n%v\n%v", round.Result["0"].Symbols[5:10], round.Result["0"].Symbols[10:15], round.Result["0"].Symbols[15:20])
	t.Logf("Round.R[0] point: %v", round.Result["0"].Point)
	// t.Logf("Round: %v", round)
}

// 多次Spin驗證RTP
func Test_SpinRoundsSG006(t *testing.T) {
	// create a new slot machine+
	mock_game := games.NewGames(sg006.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())

	// spin
	Total := decimal.Zero
	TotalNG := decimal.Zero
	TotalBG1 := decimal.Zero
	TotalBG2 := decimal.Zero
	Totalbet := decimal.NewFromFloat(500.0)
	rounds := 50000000
	rtp := "98"
	BG1Times, BG2Times, BG1andBG2Times, WinTimes := 0, 0, 0, 0
	t.Logf("rtp=%v rounds=%v", rtp, rounds)

	for time := 0; time < rounds; time++ {
		var (
			r_default = games.Rounds{
				Id:         "1234567890",
				GameCode:   "SG006",
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

		mock_game := games.NewGames(sg006.New())
		// spin
		round, err := mock_game.Spin(rtp, Totalbet, []string{"100"}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}
		//fmt.Printf("%v\n", round.Position)
		if round.Position.IsBonusGame1() {
			BG1Times++
			TotalBG1 = TotalBG1.Add(round.TotalPoint).Sub(round.Result["0"].Point)
		}
		if round.Position.IsBonusGame2() {
			BG2Times++
			TotalBG2 = TotalBG2.Add(round.TotalPoint).Sub(round.Result["0"].Point)
		}
		if round.Status.IsBonusGame2() && round.Status.IsBonusGame2() {
			BG1andBG2Times++

		}

		TotalNG = TotalNG.Add(round.Result["0"].Point)
		if round.TotalPoint.GreaterThan(decimal.Zero) {
			WinTimes++
		}
		Total = Total.Add(round.TotalPoint)
	}
	fmt.Printf("rounds: %v\n", rounds)
	//fmt.Printf("Total: %v\n", Total)
	fmt.Printf("EXP: %v\n", Total.Div(decimal.NewFromInt(int64(rounds))).Div(Totalbet))
	fmt.Printf("WinRate: %v\n", float64(WinTimes)/float64(rounds))
	fmt.Printf("NGEXP: %v\n", TotalNG.Div(decimal.NewFromInt(int64(rounds))).Div(Totalbet))

	fmt.Printf("VSHitRate: %v\n", float64(BG1Times)/float64(rounds))
	fmt.Printf("VSEXP: %v\n", TotalBG1.Div(decimal.NewFromInt(int64(rounds))).Div(Totalbet))

	fmt.Printf("FGTimesHitRate: %v\n", float64(BG2Times)/float64(rounds))
	fmt.Printf("FGEXP: %v\n", TotalBG2.Div(decimal.NewFromInt(int64(rounds))).Div(Totalbet))

	fmt.Printf("FGinVSHitRate: %v\n", float64(BG1andBG2Times)/float64(rounds))
}

func BenchmarkSG006(b *testing.B) {
	mock_game := games.NewGames(sg006.New())
	for i := 0; i < b.N; i++ {
		// mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"8"}, r_default)

		// spin
		round, _ := mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"100"}, r_default)

		// respin
		if round.Position == 16 {
			round, _ = mock_game.Spin("98", decimal.NewFromFloat(8.0), []string{"1"}, *round)
		}
	}
}

func BenchmarkSG006ThreadsafeParallel(b *testing.B) {
	mock_game := games.NewGames(sg006.New())
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

var pays = [...][5]float64{
	{0.0, 0.0, 0.0, 0.0, 0.0},         //WW
	{0.0, 50.0, 200.0, 900.0, 5000.0}, //H1
	{0.0, 20.0, 100.0, 600.0, 3000.0}, //H2
	{0.0, 20.0, 70.0, 300.0, 2000.0},  //H3
	{0.0, 20.0, 70.0, 300.0, 2000.0},  //H4
	{0.0, 20.0, 60.0, 200.0, 1500.0},  //H5
	{0.0, 0.0, 40.0, 150.0, 400.0},    //LA
	{0.0, 0.0, 40.0, 150.0, 400.0},    //LK
	{0.0, 0.0, 25.0, 90.0, 300.0},     //LQ
	{0.0, 0.0, 25.0, 90.0, 300.0},     //LJ
	{0.0, 0.0, 20.0, 70.0, 180.0},     //LT
	{0.0, 0.0, 20.0, 70.0, 180.0},     //LN
	{0.0, 0.0, 0.0, 0.0, 0.0},         //SE
	{0.0, 100.0, 550.0, 0.0, 0.0},     //SF
	{0.0, 0.0, 0.0, 0.0, 0.0},         //SB
}

func Test_LinesSG006(t *testing.T) {
	Rounds := 10000000
	Totalbet := decimal.NewFromFloat(10.0)
	// create a new slot machine
	mock_game := games.NewGames(sg006.New())
	t.Logf("GameName: %s", mock_game.Name())
	t.Logf("GameInfo: %s", mock_game.Info())
	t.Logf("NGtest Spin %d times", Rounds)
	TrueTotal := decimal.Zero
	LinesRecord := [15][6]int{} //紀錄連線狀況
	for times := 0; times < Rounds; times++ {
		// spin
		round, err := mock_game.Spin("98", Totalbet, []string{"100"}, r_default)
		if err != nil {
			t.Errorf("Spin error: %s", err.Error())
		}
		TrueTotal = TrueTotal.Add(round.Result["0"].Point)
		//Print真實滾輪/////////////////////////////////////////////////////////////////
		/*fmt.Printf("滾輪:%+v\n", round.Result["0"].Symbols[5:10])
		fmt.Printf("滾輪:%+v\n", round.Result["0"].Symbols[10:15])
		fmt.Printf("滾輪:%+v\n", round.Result["0"].Symbols[15:20])*/

		Symbols := round.Result["0"].Symbols
		m1 := make(map[string]bool)

		SFNumber := [5]int{}
		//計算有多少ＳＦ連線
		SFLengh := 0
		for i := 5; i < 10; i++ { // 檢查每一輪
			if Symbols[i] == fmt.Sprint(13) {
				SFNumber[i-5]++
			}
			if Symbols[i+5] == fmt.Sprint(13) {
				SFNumber[i-5]++
			}
			if Symbols[i+10] == fmt.Sprint(13) {
				SFNumber[i-5]++
			}
		}
		for i := 1; i < 4; i++ {
			if SFNumber[i] == 0 {
				SFNumber[i] = 1
			} else {
				SFLengh++
			}
		}
		SFLines := SFNumber[1] * SFNumber[2] * SFNumber[3]

		if SFLengh == 2 {
			LinesRecord[13][2] += SFLines
		} else if SFLengh == 3 {
			LinesRecord[13][3] += SFLines
		}

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
			Symbolarray := [6]int{} //用ARRAY紀錄SYMBOL
			if symbol == "1" {
				for i := 5; i < 10; i++ { // 檢查每一輪
					if Symbols[i] == symbol || Symbols[i] == fmt.Sprint(13) {
						Symbolarray[i-5]++
					}
					if Symbols[i+5] == symbol || Symbols[i+5] == fmt.Sprint(13) {
						Symbolarray[i-5]++
					}
					if Symbols[i+10] == symbol || Symbols[i+10] == fmt.Sprint(13) {
						Symbolarray[i-5]++
					}
				}
			} else {
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
			}
			///////////////////////////////
			L2 := Symbolarray[0] * Symbolarray[1]
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
					if L3 == 0 {
						LinesRecord[s.Int()][2] += L2
					}
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
			(decimal.NewFromInt(int64(j[1])).Mul(decimal.NewFromFloat(pays[i][0])).Add(decimal.NewFromInt(int64(j[2])).Mul(decimal.NewFromFloat(pays[i][1]))).Add(decimal.NewFromInt(int64(j[3])).Mul(decimal.NewFromFloat(pays[i][2]))).Add(decimal.NewFromInt(int64(j[4])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[5])).Mul(decimal.NewFromFloat(pays[i][4])))).Div(decimal.NewFromInt(int64(100)).Mul(decimal.NewFromInt(int64(Rounds))))
		str += fmt.Sprintf("%v", CurRTP)

		str += fmt.Sprintf("\n")
		TotolRTP = TotolRTP.Add(CurRTP)
	}
	str += fmt.Sprintf("%5s|%6s|%5s|%5s|%5s|%5s|%5s|%12s|%12s|%12s|%12s|%12s|%v\n", "Totol", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", TotolRTP)

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////
	title := fmt.Sprintf("\n%5s|%6s|%5s|%5s|%5s|%5s|%5s|%12s|%12s|%12s|%12s|%12s|%10s\n", "index", "symbol", "pay-1", "pay-2", "pay-3", "pay-4", "pay-5", "hit_freq.-1", "hit_freq.-2", "hit_freq.-3", "hit_freq.-4", "hit_freq.-5", "RTP")
	title += fmt.Sprintf("%5s|%6s|%5s|%5s|%5s|%5s|%5s|%12s|%12s|%12s|%12s|%12s|%10s\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	/*j_string, err := json.Marshal(str)
	/*j_string, err
	 := json.Marshal(str)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("LinesRecord: %s", string(j_string))*/
	t.Logf(title + str)
	fmt.Printf("\n%v", TrueTotal.Div(Totalbet).Div(decimal.NewFromInt(int64(Rounds))))
}
