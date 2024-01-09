package dcg014_test

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/shopspring/decimal"

	"github.com/death12358/digitalopen/games"
	"github.com/death12358/digitalopen/games/random"
	"github.com/death12358/digitalopen/slot/dcg014"
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

// print reelstrips.
func Test_Ge014(t *testing.T) {
	dcg014.Generatecsv()
}

func freq(freqquan [11]int, a decimal.Decimal, bet decimal.Decimal) [11]int {
	freqbar := [10]int64{0, 1, 3, 5, 10, 20, 50, 100, 200, 500}
	for i := 0; i < len(freqbar)-1; i++ {
		if a.GreaterThan(decimal.NewFromInt(freqbar[i]).Mul(bet)) && a.LessThanOrEqual(decimal.NewFromInt(freqbar[i+1]).Mul(bet)) {
			freqquan[i+1] += 1
		}
	}
	if a.GreaterThan(decimal.NewFromInt(freqbar[9]).Mul(bet)) {
		freqquan[10] += 1
	}
	if a.IsZero() {
		freqquan[0] += 1
	}
	return freqquan
}

var (
	r_default = games.Rounds{
		Id:         "🔴",
		GameCode:   "dcg014",
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

// Test One Spin.
func Test_Slotdcg014(t *testing.T) {
	// create a new slot machine
	s := games.NewGames(dcg014.New())
	round, err := s.Spin("98", decimal.NewFromFloat(1.0), []string{"100"}, r_default)
	if err != nil {
		t.Errorf("Spin error: %6s", err.Error())
	}
	j_round, err := json.Marshal(round)
	if err != nil {
		t.Errorf("Marshal error: %6s", err.Error())
	}
	t.Logf("Round: %6s", string(j_round))
	//t.Logf("Round: %v", round)
}

// Test NormalGame expection &&bonus game &&free game triggered Frequency;
// Record the RTP contribution of each elimination during the normal game.
func Test_Roundsdcg014(t *testing.T) {
	bg, fg := 0, 0
	el1, el2, el3 := decimal.Zero, decimal.Zero, decimal.Zero
	// create a new slot machine+
	mock_game := dcg014.New()
	rtp := "98"
	Total := decimal.Zero
	unitbet := "100"
	rounds := 1000000
	bet := decimal.NewFromFloat(1.0)
	freqquan := [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ngReelsLen, _ := mock_game.GetNGReelsLen(unitbet, rtp)
	for time := 0; time < rounds; time++ {
		r_default = games.Rounds{
			Id:         "1234897684",
			GameCode:   "dcg014",
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
		pos := random.Intsn(ngReelsLen)
		round, reels := mock_game.NormalGameFlow(bet, unitbet, rtp, pos, &r_default)
		if mock_game.IsFGWin(unitbet, reels) {
			fg++
		}
		if mock_game.IsBonusWin(unitbet, reels) {
			bg++
		}
		//print morethan 100 mul bet.
		// if round.TotalPoint.GreaterThanOrEqual(bet.Mul(decimal.NewFromInt(100))) {
		// 	j_r, _ := json.Marshal(round)
		// 	t.Logf("%+v\n", string(j_r))
		// }
		//frequency
		freqquan = freq(freqquan, round.TotalPoint, bet)
		Total = Total.Add(round.TotalPoint)
		for _, res := range round.Result {
			var EXTRA ExtraStruct
			json.Unmarshal([]byte(res.Extra[0]), &EXTRA)
			res.Extra = EXTRA.CascadingPoints
			for k, v := range res.Extra {
				ex, _ := decimal.NewFromString(v)
				if k == 0 {
					el1 = el1.Add(ex)
				}
				if k == 1 {
					el2 = el2.Add(ex)
				}
				if k >= 2 {
					el3 = el3.Add(ex)
				}
			}
		}
	}
	t.Logf("NGRTP: %v", Total.Div(decimal.NewFromInt(int64(rounds))))
	t.Logf("NGel1rtp: %v", el1.Div(decimal.NewFromInt(int64(rounds))))
	t.Logf("NGel2rtp: %v", el2.Div(decimal.NewFromInt(int64(rounds))))
	t.Logf("NGel3rtp: %v", el3.Div(decimal.NewFromInt(int64(rounds))))
	t.Logf("FREQUNET_FG: %v", fg)
	t.Logf("FREQUNET_BG: %v", bg)
	str := ""
	title := fmt.Sprintf("\n%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|\n",
		"0", "(0-1]", "(1,3]", "(3,5]", "(5,10]", "(10,20]", "(20,50]", "(50,100]", "(100,200]", "(200,500]", "500up",
	)
	for _, v := range freqquan {
		str += fmt.Sprintf("%9v|", v)
	}
	t.Logf(title + str)
}

// FreeGame floor Spin test
func Test_dcgfloor014(t *testing.T) {
	round := 1000000
	level := 4
	rtp := "98"
	unitbet := "100"
	freqquan := [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	respincount, hitrate, next0, next1, next2, next3 := 0, 0, 0, 0, 0, 0
	fg_game := dcg014.Dcg014fg
	bet := decimal.NewFromInt(1.0)
	total := decimal.Zero
	fglen := fg_game.GetReelsLen(1, "100", rtp)
	r_default = games.Rounds{
		Id:         "1234897684",
		GameCode:   "dcg014",
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
	file, err := os.Create("output.csv")
	file2, err2 := os.Create("dataoutput.csv")
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	defer file.Close()
	defer file2.Close()
	// 创建 CSV writer
	writer := csv.NewWriter(file)
	writer2 := csv.NewWriter(file2)
	defer writer.Flush()
	defer writer2.Flush()

	for i := 0; i < round; i++ {
		position := random.Intsn(fglen)
		reelStrips, initialposition, show_reelStrips, _ := fg_game.FirstSpinReelStripts(level, unitbet, rtp, position)
		fg_result, _, _, _, nextLevel, respin := fg_game.CalcFGResult(level, unitbet, rtp, *reelStrips, initialposition, r_default.Id, r_default.Brand, r_default.Username, show_reelStrips, bet)
		total = total.Add(fg_result.Point)
		// LET ME SEE SEE WHO IS THE LUCKY COOKEE XD.
		if fg_result.Point.GreaterThanOrEqual(bet.Mul(decimal.NewFromInt(500))) {
			j_r, _ := json.Marshal(fg_result)
			j_d, _ := json.Marshal(fg_result.Symbols)
			//t.Logf("%+v\n", string(j_r))
			writer.Write([]string{string(j_r)})
			writer2.Write([]string{string(j_d)})
		}

		freqquan = freq(freqquan, fg_result.Point, bet)
		if respin > 0 {
			respincount++
		}
		if nextLevel == level {
			next0++
		}
		if nextLevel == level+1 {
			next1++
		}
		if nextLevel == level+2 {
			next2++
		}
		if nextLevel == level+3 {
			next3++
		}
		if fg_result.Point.IsPositive() {
			hitrate++
		}
	}
	t.Logf("FG Expection: %+v\n", total.Div(decimal.NewFromInt(int64(round))))
	t.Logf("hitrate: %+v\n", (decimal.NewFromInt(int64(hitrate))).Div(decimal.NewFromInt(int64(round))))
	t.Logf("respincount: %+v\n", (decimal.NewFromInt(int64(respincount))).Div(decimal.NewFromInt(int64(round))))
	t.Logf("next0: %+v\n", (decimal.NewFromInt(int64(next0))).Div(decimal.NewFromInt(int64(round))))
	t.Logf("next1: %+v\n", (decimal.NewFromInt(int64(next1))).Div(decimal.NewFromInt(int64(round))))
	t.Logf("next2: %+v\n", (decimal.NewFromInt(int64(next2))).Div(decimal.NewFromInt(int64(round))))
	t.Logf("next3: %+v\n", (decimal.NewFromInt(int64(next3))).Div(decimal.NewFromInt(int64(round))))
	str := ""
	title := fmt.Sprintf("\n%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|\n",
		"0", "(0-1]", "(1,3]", "(3,5]", "(5,10]", "(10,20]", "(20,50]", "(50,100]", "(100,200]", "(200,500]", "500up",
	)
	for _, v := range freqquan {
		str += fmt.Sprintf("%9v|", v)
	}
	t.Logf(title + str)
}

// FreeGame Expection
func Test_dcg014(t *testing.T) {
	round := 10000000
	s := dcg014.New()
	bet := decimal.NewFromInt(1.0)
	total := decimal.Zero
	for i := 0; i < round; i++ {
		r_default = games.Rounds{
			Id:         "1234897684",
			GameCode:   "dcg014",
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
		r, _ := s.FreeGameFlow(bet, "100", "98", &r_default, 5)
		total = total.Add(r.TotalPoint)
		//	LogDebug(r)
		// byteRound, _ := json.Marshal(r)
		// fmt.Println(string(byteRound))
	}
	t.Logf("FG Expection: %+v\n", total.Div(decimal.NewFromInt(int64(round))))
}

// Test Spin : NG/FG/BG  contribution in RTP.
func Test_SlotRTPdcg014(t *testing.T) {
	// create a new slot machine
	Total, FGTotal, NGTotal, BGTotal := decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero
	freqquan := [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	BET := decimal.NewFromInt(100)
	Time := 10000000
	j := 0
	k := 0
	for i := 0; i < Time; i++ {
		s := dcg014.New()
		round, _ := s.Spin("98", BET, []string{"100"}, r_default)
		Total = Total.Add(round.TotalPoint)
		NGTotal = NGTotal.Add(round.Result["0"].Point)
		// LET ME SEE SEE WHO IS THE LUCKY COOKEE XD.
		// if round.TotalPoint.GreaterThanOrEqual(BET.Mul(decimal.NewFromInt(500))) {
		// 	j_r, _ := json.Marshal(round)
		// 	t.Logf("%+v\n", string(j_r))
		// }
		if int(round.Status) >= int(games.FreeGame) && int(round.Status) < int(games.Bonus) {
			j++
			FGTotal = FGTotal.Add(round.TotalPoint.Add((round.Result["0"].Point).Neg()))

		} else if int(round.Status) >= int(games.Bonus) {
			k++
			BGTotal = BGTotal.Add(round.TotalPoint.Add((round.Result["0"].Point).Neg()))
			//Record a case where a bonus and free game were triggered at the same time.
		}
		//  else if int(round.Status) >= 350 {
		// 	j_round, err := json.Marshal(round)
		// 	if err != nil {
		// 		t.Errorf("Marshal error: %6s", err.Error())
		// 	}
		// 	t.Logf("Round: %6s", string(j_round))
		// }
		freqquan = freq(freqquan, round.TotalPoint, BET)
	}
	t.Logf("RTP: %v", Total.Div(decimal.NewFromInt(int64(Time))))
	t.Logf("EXPECTION_NG: %v", NGTotal.Div(decimal.NewFromInt(int64(Time))))
	t.Logf("EXPECTION_FG: %v", FGTotal.Div(decimal.NewFromInt(int64(Time))))
	t.Logf("EXPECTION_BG: %v", BGTotal.Div(decimal.NewFromInt(int64(Time))))
	t.Logf("FREQUNET_FG: %v", j)
	t.Logf("FREQUNET_BG: %v", k)
	str := ""
	title := fmt.Sprintf("\n%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|%9s|\n",
		"0", "(0-1]", "(1,3]", "(3,5]", "(5,10]", "(10,20]", "(20,50]", "(50,100]", "(100,200]", "(200,500]", "500up",
	)
	for _, v := range freqquan {
		str += fmt.Sprintf("%9v|", v)
	}
	t.Logf(title + str)
	//t.Logf("Frequency distribution: %+v\n", freqquan)
}
