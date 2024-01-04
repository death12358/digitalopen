package games_test

import (
	"encoding/json"
	"testing"

	"games"

	"github.com/shopspring/decimal"
)

func Test_RoundFormat_Marshal(t *testing.T) {
	test_round := &games.Rounds{
		Id:         "1234567890",
		GameCode:   "TS001",
		Brand:      "brand_test",
		Username:   "user_test",
		Status:     games.State(0),
		Position:   games.State(0),
		Stages:     0,
		Result:     games.NewResults(),
		Start:      0,
		Fisish:     0,
		TotalBet:   decimal.Zero,
		TotalPoint: decimal.Zero,
	}

	j_r, err := json.Marshal(test_round)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("%+v\n", string(j_r))
}

func Test_ResultsFormat(t *testing.T) {
	test_result := games.NewResults()

	test_result["aa"] = &games.Records{
		Id: "AA",
	}

	j_r, err := json.Marshal(test_result)
	if err != nil {
		t.Errorf("Marshal error: %s", err.Error())
	}
	t.Logf("%+v\n", string(j_r))

	var result_unmarshal games.Results
	err = result_unmarshal.Unmarshal(j_r)
	if err != nil {
		t.Errorf("Unmarshal error: %s", err.Error())
	}
	t.Logf("%+v\n", result_unmarshal)
}
