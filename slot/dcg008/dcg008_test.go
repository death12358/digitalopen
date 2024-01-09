package dcg008_test

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"testing"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/slot/dcg008"
	"github.com/shopspring/decimal"
)

type ExtraStruct struct {
	CascadingPoints []string `json:"cascadingpoints"`
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
	15: "NA",
}
var pays = [...][5]float64{
	{0.0, 0.0, 0.0, 0.0, 0.0},      //WW
	{0.0, 0.0, 80.0, 100.0, 120.0}, //H1
	{0.0, 0.0, 60.0, 80.0, 100.0},  //H2
	{0.0, 0.0, 45.0, 60.0, 100.0},  //H3
	{0.0, 0.0, 45.0, 60.0, 100.0},  //H4
	{0.0, 0.0, 0.0, 0.0, 0.0},      //H5
	{0.0, 0.0, 15.0, 30.0, 45.0},   //LA
	{0.0, 0.0, 10.0, 15.0, 30.0},   //LK
	{0.0, 0.0, 10.0, 15.0, 30.0},   //LQ
	{0.0, 0.0, 10.0, 15.0, 30.0},   //LJ
	{0.0, 0.0, 5.0, 10.0, 15.0},    //LT
	{0.0, 0.0, 5.0, 10.0, 15.0},    //LN
	{0.0, 0.0, 0.0, 0.0, 0.0},      //SE
	{0.0, 0.0, 0.0, 0.0, 0.0},      //SF
	{0.0, 0.0, 0.0, 0.0, 0.0},      //SB
	{0.0, 0.0, 0.0, 0.0, 0.0},      //NA
}

func Test_eee008(t *testing.T) {
	dcg008.Generatecsv()
}
func LogDebug(a ...any) {

	pc, file, lineNo, ok := runtime.Caller(1)
	if !ok {

		fmt.Println("runtime.Caller() failed")
	}

	// str := fmt.Sprintf("%6s", a)

	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素

	prefix := fmt.Sprintf("%6s---%6s(...)---lines:%d", fileName, funcName, lineNo)
	prefix = fmt.Sprintf("\x1b[0;33;m%6s\x1b[0m", prefix)
	// result := fmt.Sprintf("%6s, DEBUG:%6s", prefix, str)

	fmt.Print(prefix)
	fmt.Print(a...)
	fmt.Print("\n\r")
}

var (
	r_default = games.Rounds{
		Id:         "1234897684",
		GameCode:   "dcg008",
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

// Spin for NG
func Test_dcg008(t *testing.T) {
	Time := 5000000
	s := dcg008.New()
	bet := decimal.NewFromInt(1.0)
	total := decimal.Zero
	for i := 0; i < Time; i++ {
		pos := random.Intsn([]int{157, 157, 157, 157, 157})
		r, _ := s.NormalGameFlow(bet, "100", "98", pos, &r_default)
		// j_round, err := json.Marshal(r)
		// if err != nil {
		// 	t.Errorf("Marshal error: %6s", err.Error())
		// }
		// t.Logf("Round: %6s", string(j_round))
		total = total.Add(r.TotalPoint)
	}
	t.Logf("Taltol: %+v\n", total.Div(decimal.NewFromInt(int64(Time))))
}

func Test_dcgRTP008(t *testing.T) {
	Time := 50000000
	s := dcg008.New()
	bet := decimal.NewFromInt(1.0)
	total := decimal.Zero
	for i := 0; i < Time; i++ {

		r, _ := s.Spin("98", bet, []string{"100"}, r_default)
		// j_round, err := json.Marshal(r)
		// if err != nil {
		// 	t.Errorf("Marshal error: %6s", err.Error())
		// }
		// t.Logf("Round: %6s", string(j_round))
		total = total.Add(r.TotalPoint)
	}
	t.Logf("Taltol: %+v\n", total.Div(decimal.NewFromInt(int64(Time))))
}

// 一次Spin
func Test_Slotdcg008(t *testing.T) {
	// create a new slot machine
	s := games.NewGames(dcg008.New())
	//s := dcg008.New()
	//t.Logf("GameName: %6s", s.Name())
	//t.Logf("GameInfo: %6s", s.Info())
	//pos := [][]int{{155, 156, 0}, {1, 2, 3}, {73, 74, 75}, {45, 46, 47}, {138, 139, 140}}
	// spin

	round, err := s.Spin("40", decimal.NewFromFloat(10.0), []string{"100"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %6s", err.Error())
	}

	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %6s", err.Error())
	}
	t.Logf("Round: %6s", string(j_round))

	// t.Logf("Round: %v", round)
}

// 多次Spin驗證消除第I次RTP
func Test_Roundsdcg008(t *testing.T) {
	// create a new slot machine+
	Total := decimal.Zero
	NGTotal1, NGTotal2, NGTotal3, FGTotal1, FGTotal2, FGTotal3, BGTotal :=
		decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero
	rounds := 10000000
	i := 0
	k := 0
	for time := 0; time < rounds; time++ {
		mock_game := dcg008.New()
		// spin
		round, err := mock_game.Spin("92", decimal.NewFromFloat(100.0), []string{"100"}, r_default)
		if err != nil {
			t.Errorf("Spin error: %6s", err.Error())
		}
		if round.Status == 356 {
			j_round, err := json.Marshal(round)
			if err != nil {
				t.Errorf("Marshal error: %6s", err.Error())
			}
			t.Logf("Round: %6s", string(j_round))

			t.Logf("Round: %v", round)
		}
		Total = Total.Add(round.TotalPoint)
		if round.Stages > 0 && int(round.Status) < int(games.Bonus) {
			i++
		}
		if round.Stages > 0 && round.Status > games.Bonus {
			k++
		}
		for _, v := range round.Result {
			var EXTRA ExtraStruct
			json.Unmarshal([]byte(v.Extra[0]), &EXTRA)
			v.Extra = EXTRA.CascadingPoints

			if round.Stages > 0 && int(round.Status) < int(games.Bonus) && v.Stages > int64(0) {

				if len(v.Extra) > 0 {

					//Total = Total.Add(v.Point)
					r, _ := decimal.NewFromString((v.Extra)[0])
					// //t.Logf("r: %6s", r)
					FGTotal1 = FGTotal1.Add(r)

				}
				if len(v.Extra) > 1 {

					//Total = Total.Add(v.Point)
					r, _ := decimal.NewFromString((v.Extra)[1])
					// //t.Logf("r: %6s", r)
					FGTotal2 = FGTotal2.Add(r)

				}
				if len(v.Extra) > 2 {
					rr := decimal.Zero
					//Total = Total.Add(v.Point)
					for j := 2; j < len(v.Extra); j++ {
						r, _ := decimal.NewFromString((v.Extra)[j])
						rr = rr.Add(r)
					}
					// //t.Logf("r: %6s", r)
					FGTotal3 = FGTotal3.Add(rr)

				}
			}

			if v.Stages == 0 {
				if len(v.Extra) > 0 {

					//Total = Total.Add(v.Point)
					r, _ := decimal.NewFromString((v.Extra)[0])
					// //t.Logf("r: %6s", r)
					NGTotal1 = NGTotal1.Add(r)

				}
				if len(v.Extra) > 1 {

					//Total = Total.Add(v.Point)
					r, _ := decimal.NewFromString((v.Extra)[1])
					// //t.Logf("r: %6s", r)
					NGTotal2 = NGTotal2.Add(r)

				}
				if len(v.Extra) > 2 {
					rr := decimal.Zero
					//Total = Total.Add(v.Point)
					for j := 2; j < len(v.Extra); j++ {
						r, _ := decimal.NewFromString((v.Extra)[j])
						rr = rr.Add(r)
					}
					// //t.Logf("r: %6s", r)
					NGTotal3 = NGTotal3.Add(rr)

				}
			}
			if round.Stages > 0 && round.Status > games.Bonus {

				BGTotal = BGTotal.Add(v.Point)
			}
		}

	}

	fmt.Printf("FG hit_fq: %v\n", decimal.NewFromInt(int64(i)).Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("BG hit_fq: %v\n", decimal.NewFromInt(int64(k)).Div(decimal.NewFromInt(int64(rounds))))
	//fmt.Printf("EXP: %v\n", Total.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("FG_1: %v\n", FGTotal1.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("FG_2: %v\n", FGTotal2.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("FG_3+: %v\n", FGTotal3.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("NG_1: %v\n", NGTotal1.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("NG_2: %v\n", NGTotal2.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("NG_3+: %v\n", NGTotal3.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("BG: %v\n", BGTotal.Div(decimal.NewFromInt(int64(rounds))))
	fmt.Printf("RTP: %v\n", Total.Div(decimal.NewFromInt(int64(rounds))))
}

// NG统计
func Test_NGDCG008(t *testing.T) {
	Rounds := 1000000
	unitbet := 100
	EXPANDINGWILD := true
	sf7, sf10, sf15, bb := 0, 0, 0, 0
	//f1, f2, f3, f4, f5, f6, f7, f8, f9 := 0, 0, 0, 0, 0, 0, 0, 0, 0
	freqbar := [10]int64{0, 1, 3, 5, 10, 20, 50, 100, 200, 500}

	freqquan := [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// create a new slot machine
	mock_game := dcg008.New()
	t.Logf("GameName: %6s", mock_game.Name())
	t.Logf("GameInfo: %6s", mock_game.Info())
	t.Logf("BigFGtest Spin %d times", Rounds)
	Tatol := decimal.Zero
	totalsf, totalbb := decimal.Zero, decimal.Zero
	LinesRecord := [15][14]int{} //紀錄連線狀況
	BET := decimal.NewFromFloat(100.0)
	for times := 0; times < Rounds; times++ {
		// spin
		round, err := mock_game.Spin("96", BET, []string{"100"}, r_default)
		if err != nil {
			t.Errorf("Spin error: %6s", err.Error())
		}
		// if round.Status >= 289 {
		// 	fmt.Print("both feature game awarded.")
		// }
		if (round.Status == 33 || round.Status == 36) && round.Stages == 7 {
			//
			sf7++
		}
		if (round.Status == 33 || round.Status == 36) && round.Stages == 10 {
			sf10++
		}
		if (round.Status == 33 || round.Status == 36) && round.Stages == 15 {
			sf15++
		}
		if round.Status > 320 {
			bb++
		}
		if round.TotalPoint.IsZero() {
			freqquan[0] += 1
		}
		for i := 0; i < len(freqbar)-1; i++ {
			if round.TotalPoint.GreaterThan(decimal.NewFromInt(freqbar[i]).Mul(BET)) && round.TotalPoint.LessThanOrEqual(decimal.NewFromInt(freqbar[i+1]).Mul(BET)) {
				freqquan[i+1] += 1
			}
		}
		if round.TotalPoint.GreaterThan(decimal.NewFromInt(freqbar[9]).Mul(BET)) {
			freqquan[10] += 1
		}
		Tatol = Tatol.Add(round.TotalPoint.Div(decimal.NewFromInt(int64(100 * Rounds))))

		//see result////////////////////////////
		//j_r, _ := json.Marshal(round)
		//t.Logf("%+v\n", string(j_r))
		///////////////////////////////////////
		//for Statistacs_Scatte
		//Statistacs_SF := true
		//Statistacs_SB := true
		for Cth, res := range round.Result {
			var EXTRA ExtraStruct
			json.Unmarshal([]byte(res.Extra[0]), &EXTRA)
			res.Extra = EXTRA.CascadingPoints
			//t.Logf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			if res.Case.IsNotStartedYet() && round.Stages > 0 && int(round.Status) < int(games.Bonus) && res.Stages > int64(0) {
				totalsf = totalsf.Add(res.Point)
				//fmt.Println(res.Point)
			}
			if round.Stages > 0 && round.Status > games.Bonus {
				totalbb = totalbb.Add(res.Point)
				//fmt.Println(res.Point)
			}
			if Cth == "0" {
				Symbols := res.Symbols

				//紀錄第一輪的符號
				left := 0
				right := 0

				l := len(res.Extra)
				for k := 1; k <= l; k++ {
					if k == 1 {
						left = 5
						right = 15
					}
					if k > 1 {
						left = 30*k - 20
						right = 30*k - 10
					}
					r1 := []string{}
					m1 := make(map[string]bool)
					for i := left; i <= right; i = i + 5 {
						if !m1[Symbols[i]] {
							r1 = append(r1, Symbols[i])
							m1[Symbols[i]] = true
						}
					}
					//test
					//t.Logf("%+v\n", r1)
					for _, symbol := range r1 { //迴圈檢查與R1相同符號的個數
						Symbolarray := [6]int{}
						//expanding      //用ARRAY紀錄SYMBOL
						for i := left; i < left+5; i++ { // 檢查每一輪
							if EXPANDINGWILD && (Symbols[i] == "0" || Symbols[i+5] == "0" || Symbols[i+10] == "0") {
								if Symbols[i] != "13" && Symbols[i] != "14" {
									Symbols[i] = "0"
								}
								if Symbols[i+5] != "13" && Symbols[i+5] != "14" {
									Symbols[i+5] = "0"
								}
								if Symbols[i+10] != "13" && Symbols[i+10] != "14" {
									Symbols[i+10] = "0"
								}
							}
							//t.Logf("%+v\n", Symbols)
							if Symbols[i] == symbol || Symbols[i] == fmt.Sprint(0) {
								Symbolarray[i-left]++
							}
							if Symbols[i+5] == symbol || Symbols[i+5] == fmt.Sprint(0) {
								Symbolarray[i-left]++
							}
							if Symbols[i+10] == symbol || Symbols[i+10] == fmt.Sprint(0) {
								Symbolarray[i-left]++
							}
						}
						//test
						//t.Logf("Symbolarray: %+v\n", Symbolarray)
						///////////////////////////////
						L3 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2]                                   //3連線個數
						L4 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3]                  //4連線個數
						L5 := Symbolarray[0] * Symbolarray[1] * Symbolarray[2] * Symbolarray[3] * Symbolarray[4] //5連線個數

						s := games.ToSymbol(symbol)
						//fmt.Printf("%+v\n", symbol) ///////////////////////////////////////////////
						//fmt.Printf("Symbolarray:%+v\n", Symbolarray)
						//	key, _ := strconv.Atoi(Cth)
						//Record "Cascading 3+
						key := k
						if k > 2 {
							key = 3
						}
						//for normal symbol
						if s != games.ToSymbol("13") || s != games.ToSymbol("14") {
							LinesRecord[s.Int()][(key-1)*4+2] += L5
							if L5 == 0 {
								LinesRecord[s.Int()][(key-1)*4+1] += L4
								if L4 == 0 {
									LinesRecord[s.Int()][(key-1)*4] += L3
								}
							}
							//t.Logf("LinesRecord: %+v\n", LinesRecord)
						}

					}
				}
			}
		}
	}
	//t.Logf("LinesRecord: %+v\n", LinesRecord)

	//Rounds = Rounds * 6
	//strscatter := ""
	str := ""
	var TotalRTP, NGRTP1, NGRTP2, NGRTP3 decimal.Decimal
	for i, j := range LinesRecord { //i:符號編號 j:符號對應的陣列 k:k連線 l:有l條線
		str += fmt.Sprintf("%6v|%6v|", i, IntToSymbolString[i])
		for n := 2; n < 5; n++ { //pays
			str += fmt.Sprintf("%6v|", pays[i][n])
		}

		for l := 0; l < 3; l++ { //連線數
			str += fmt.Sprintf("%6v|", (decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds)))).Round(4))
		}
		CurRTP1 := (decimal.NewFromInt(int64(j[0])).Mul(decimal.NewFromFloat(pays[i][2])).Add(decimal.NewFromInt(int64(j[1])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[2])).Mul(decimal.NewFromFloat(pays[i][4]))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds)))).Round(4))
		str += fmt.Sprintf("%6v|", CurRTP1)
		for l := 4; l < 7; l++ { //連線數
			str += fmt.Sprintf("%6v|", (decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds)))).Round(4))
		}
		CurRTP2 := (decimal.NewFromInt(int64(j[4])).Mul(decimal.NewFromFloat(pays[i][2])).Add(decimal.NewFromInt(int64(j[5])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[6])).Mul(decimal.NewFromFloat(pays[i][4]))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds)))).Round(4))
		str += fmt.Sprintf("%6v|", CurRTP2)
		for l := 8; l < 11; l++ { //連線數
			str += fmt.Sprintf("%6v|", (decimal.NewFromInt(int64(j[l])).Div(decimal.NewFromInt(int64(Rounds)))).Round(4))
		}
		CurRTP3 := (decimal.NewFromInt(int64(j[8])).Mul(decimal.NewFromFloat(pays[i][2])).Add(decimal.NewFromInt(int64(j[9])).Mul(decimal.NewFromFloat(pays[i][3]))).Add(decimal.NewFromInt(int64(j[10])).Mul(decimal.NewFromFloat(pays[i][4]))).Div(decimal.NewFromInt(int64(unitbet)).Mul(decimal.NewFromInt(int64(Rounds)))).Round(4))
		str += fmt.Sprintf("%6v|", CurRTP3)
		CurRTP := CurRTP1.Add(CurRTP2).Add(CurRTP3)

		str += fmt.Sprintf("%v", CurRTP)
		str += fmt.Sprintf("\n")
		TotalRTP = TotalRTP.Add(CurRTP)
		NGRTP1 = NGRTP1.Add(CurRTP1)
		NGRTP2 = NGRTP2.Add(CurRTP2)
		NGRTP3 = NGRTP3.Add(CurRTP3)
	}

	str += fmt.Sprintf("%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%s|%6s\n",
		"Totol", "-", "-", "-", "-", "-", "-", "-", NGRTP1, "-", "-", "-", NGRTP2, "-", "-", "-", NGRTP3, TotalRTP)

	//fmt.Printf("%+v\n", LinesRecord) //////////////////////////////////////////////////////////////////////////////////////////////////////////
	title := fmt.Sprintf("\n%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s\n",
		"index", "Symbol", "pay-3", "pay-4", "pay-5",
		"hf_13", "hf_14", "hf_15", "RTP1",
		"hf_23", "hf_24", "hf_25", "RTP2",
		"hf_33+", "hf_34+", "hf_35+", "RTP3",
		"NGRTP",
	)
	title += fmt.Sprintf("%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s|%6s\n", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-")
	/*j_string, err := json.Marshal(str)
	/*j_string, err
	 := json.Marshal(str)
	if err != nil {
		t.Errorf("Marshal error: %6s", err.Error())
	}
	t.Logf("LinesRecord: %6s", string(j_string))*/
	t.Logf(title + str)

	/*str2 := fmt.Sprintf("%20v|", "BigFree Game")
	str2 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(sf7)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8))
	str2 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(sf10)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8))
	str2 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(sf15)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8))
	str2 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(sf7 + sf10 + sf15)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8))
	str2 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(1)).Div(decimal.NewFromInt(int64(sf7 + sf10 + sf15)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8)))
	str2 += fmt.Sprintf("%20v|", (totalsf.Div((decimal.NewFromInt(int64(sf7 + sf10 + sf15)))).Round(8)))
	str2 += fmt.Sprintf("%20v|", (totalsf.Div((decimal.NewFromInt(int64(Rounds * 100)))).Round(8)))
	str2 += fmt.Sprintf("\n")
	title2 := fmt.Sprintf("\n%20s|%20s|%20s|%20s|%20s|%20s|%20s|%20s|\n",
		"game_name", "hf7", "hf10", "hf15", "hit_freq", "Averageinter", "EX", "sfrtp",
	)
	t.Logf(title2 + str2)

	str3 := fmt.Sprintf("%20v|", "Bonus Game")
	str3 += fmt.Sprintf("%20v|", ((decimal.NewFromInt(int64(bb))).Div(decimal.NewFromInt(int64(Rounds))).Round(8)))
	str3 += fmt.Sprintf("%20v|", (decimal.NewFromInt(int64(1)).Div(decimal.NewFromInt(int64(bb)).Div(decimal.NewFromInt(int64(Rounds)))).Round(8)))
	str3 += fmt.Sprintf("%20v|", (totalbb.Div((decimal.NewFromInt(int64(bb)))).Round(8)))
	str3 += fmt.Sprintf("%20v|", (totalbb.Div((decimal.NewFromInt(int64(Rounds * 100)))).Round(8)))
	str3 += fmt.Sprintf("\n")
	title3 := fmt.Sprintf("\n%20s|%20s|%20s|%20s|%20s|\n",
		"game_name", "hf", "Averageinter", "EX", "bbrtp",
	)
	t.Logf(title3 + str3)
	// if Tatol != TotalRTP {
	// 	fmt.Printf("??????????????????????????????????????????????")
	// } else {
	// 	fmt.Printf("TRUE")
	// }*/
	str4 := ""
	title4 := fmt.Sprintf("\n%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|\n",
		"0", "(0-1]", "(1,3]", "(3,5]", "(5,10]", "(10,20]", "(20,50]", "(50,100]", "(100,200]", "(200,500]", "500up",
	)
	for _, v := range freqquan {
		str4 += fmt.Sprintf("%9v|", v)
	}
	t.Logf(title4 + str4)
	t.Logf("Total: %+v\n", Tatol)
	t.Logf("Frequency distribution: %+v\n", freqquan)
}

// 破产
func Test_SpinSlotdcg008(t *testing.T) {
	Balance := decimal.NewFromInt(2000)
	Bet := decimal.NewFromInt(10)
	Obj := decimal.NewFromInt(60000)
	Effective := decimal.Zero
	k := 0
	for Balance.GreaterThanOrEqual(Bet) {
		k++

		mock_game := dcg008.New()
		round, err := mock_game.Spin("96", Bet, []string{"100"}, r_default)
		if err != nil {
			t.Errorf("Spin error: %6s", err.Error())
		}
		Effective = Effective.Add(Bet)
		Balance = Balance.Add(round.TotalPoint).Add(Bet.Neg())
		t.Logf("Round: %+v, Balance: %+v, Effective: %+v\n", k, Balance, Effective)
		if Balance.GreaterThanOrEqual(Obj) {
			break
		}
	}

}
