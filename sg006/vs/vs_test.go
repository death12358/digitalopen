package vs

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_PickReward(t *testing.T) {
	TotalPoint := decimal.Zero
	Rounds := 100000000

	for i := 0; i < Rounds; i++ {
		vs := NewGame(VS_reel_98, VSPayTable)
		// PickReward - 轉轉輪決定獎勵
		a := vs.PickReward()
		TotalPoint = TotalPoint.Add(vs.GetVSPay(a))
	}
	fmt.Printf("EXP:= %v", TotalPoint.Div(decimal.NewFromInt(int64(Rounds))))
}
