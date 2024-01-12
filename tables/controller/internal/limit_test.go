package internal

import (
	"fmt"
	"math"
	"testing"
)

func TestProbabilityLimit_CalculateMergedProbabilityLimit(t *testing.T) {
	fmt.Println("機率表沒設定, rtp傳upper:10")
	probA := ProbabilityLimit{
		Upper: math.MaxInt,
		Lower: -1,
	}
	nProbA := probA.CalculateMergedProbabilityLimit(10)
	fmt.Printf("結果:Upper:%v,Lower:%v\n\n", nProbA.Upper, nProbA.Lower)

	fmt.Println("機率表設定upper:1000, lower:0, rtp傳upper:100")
	probB := ProbabilityLimit{
		Upper: 1000,
		Lower: 0,
	}
	nProbB := probB.CalculateMergedProbabilityLimit(100)
	fmt.Printf("結果:Upper:%v,Lower:%v\n\n", nProbB.Upper, nProbB.Lower)

	fmt.Println("機率表有設定upper:1000, lower:10,  rtp傳upper:1100")
	probC := ProbabilityLimit{
		Upper: 1000,
		Lower: 10,
	}
	nProbC := probC.CalculateMergedProbabilityLimit(1100)
	fmt.Printf("結果:Upper:%v,Lower:%v\n\n", nProbC.Upper, nProbC.Lower)

	fmt.Println("機率表有設定upper:1000, lower:-1, rtr沒設定")
	probD := ProbabilityLimit{
		Upper: 1000,
		Lower: -1,
	}
	nProbD := probD.CalculateMergedProbabilityLimit(-1)
	fmt.Printf("結果:Upper:%v,Lower:%v\n\n", nProbD.Upper, nProbD.Lower)
}
