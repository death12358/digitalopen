package weights

import "fmt"

// Games - WeightGames 結構
type Games struct {
	weights []int
	objects []int
	sum     int
}

// NewWeightGames - 建立 WeightGames
//  @param weights	權重
//  @param objects	物件
//  @return *WeightGames	WeightGames 物件
func NewGames(weights []int, objects []int) *Games {
	// 檢查權重與物件數量是否相同
	if len(weights) != len(objects) {
		fmt.Println("weights and objects length not equal")
		return nil
	}

	// 檢查權重是否為正整數 & 計算總權重
	sum := 0
	for _, weight := range weights {
		if weight < 0 {
			fmt.Println("weight must be positive integer")
			return nil
		}
		sum += weight
	}

	return &Games{
		weights: weights,
		objects: objects,
		sum:     sum,
	}
}

// Get - 取得物件
//  @param index	索引
//  @return decimal.Decimal	物件
func (w *Games) Get(index int) int {
	return w.objects[index]
}

// Len - 取得物件數量
//  @return int	物件數量
func (w *Games) Len() int {
	return len(w.objects)
}

// Sum - 取得總權重加總
//  @return int	總權重
func (w *Games) Sum() int {
	return w.sum
}

// Pick - 隨機取得物件
//  @param random	隨機數
//  @return decimal.Decimal	物件
//  @return int	索引
func (w *Games) Pick(random int) (int, int) {
	// 檢查隨機數是否為正整數
	if random < 0 {
		fmt.Println("random must be positive integer")
	}

	// 取得物件
	for i, weight := range w.weights {
		if random < weight {
			return w.objects[i], i
		}
		random -= weight
	}

	fmt.Println("random out of range")
	return 0, 0
}

// Picks - 隨機取得多個物件
//  @param randoms	隨機數
//  @return []int	物件
//  @return []int	索引
func (w *Games) Picks(randoms []int) ([]int, []int) {
	objects := make([]int, len(randoms))
	indices := make([]int, len(randoms))

	for i, random := range randoms {
		objects[i], indices[i] = w.Pick(random)
	}

	return objects, indices
}
