package weights_test

import (
	"fmt"
	"testing"

	"digitalopen/games/random"
	weights "digitalopen/games/weight"
)

func Test_Weight(t *testing.T) {
	wg := weights.NewGames([]int{2, 8, 90}, []int{
		1,
		2,
		3,
	})

	stat := make([]float64, wg.Len())
	round := 10000000
	for i := 0; i < round; i++ {
		rng := random.Intn(wg.Sum())
		_, idx := wg.Pick(rng)
		stat[idx]++
	}

	for _, v := range stat {
		fmt.Printf("%.4f \t\n", v/float64(round))
	}

	fmt.Printf("wc == %.2f \t\n", stat)
}
