package variance

import (
	"sync"

	"github.com/shopspring/decimal"
)

/*
公式：
變異數 = Σ( (數值 - 平均值)^2 ) / N

其中，Σ代表對所有數值進行求和，數值是數據集中的每個數值，N是數據集中的數值總數。

現在，讓我們進行一個實際的例子：

假設有一個數據集：[10, 15, 20, 25, 30]

Step 1: 計算平均值（Mean）
平均值 = (10 + 15 + 20 + 25 + 30) / 5 = 20

Step 2: 計算每個數值與平均值的差異
差異 = [10 - 20, 15 - 20, 20 - 20, 25 - 20, 30 - 20] = [-10, -5, 0, 5, 10]

Step 3: 計算差異的平方
差異的平方 = [(-10)^2, (-5)^2, 0^2, 5^2, 10^2] = [100, 25, 0, 25, 100]

Step 4: 計算平方差異的平均值
變異數 = (100 + 25 + 0 + 25 + 100) / 5 = 250 / 5 = 50

所以，這個數據集的變異數為50。
*/

// VarianceCalculator 變異數計算器
type VarianceCalculator struct {
	valueList []decimal.Decimal
	total     decimal.Decimal
}

// NewVarianceCalculator new 建立
func NewVarianceCalculator() *VarianceCalculator {
	return &VarianceCalculator{}
}

// Add 加入參數
func (vc *VarianceCalculator) Add(x decimal.Decimal) {
	vc.valueList = append(vc.valueList, x)
	vc.total = vc.total.Add(x)
}

// Variance 計算變異數
func (vc *VarianceCalculator) Variance() decimal.Decimal {
	if len(vc.valueList) < 2 {
		return decimal.Zero
	}

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		workers int = 10 //併發數量
	)

	chunkSize := (len(vc.valueList) + workers - 1) / workers
	n := decimal.NewFromInt(int64(len(vc.valueList)))
	avg := vc.total.Div(n)
	variance := decimal.Zero

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(vc.valueList) {
			end = len(vc.valueList)
		}
		wg.Add(1)
		go func(s, e int) {
			defer wg.Done()
			tempVariance := decimal.Zero
			for _, val := range vc.valueList[s:e] {
				delta := val.Sub(avg)
				tempVariance = tempVariance.Add(delta.Mul(delta))
			}
			mu.Lock()
			variance = variance.Add(tempVariance)
			mu.Unlock()

		}(start, end)
	}
	wg.Wait()

	return variance.Div(n)
}
