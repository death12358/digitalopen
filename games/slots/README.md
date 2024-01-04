# slots

## Install

```bash
go get -u -v gitlab.com/gaas_module/games/slots
```

---

## 基本使用

```go
// 宣告轉輪表
var ngReelStrips = games.ReelStrips{
 {2, 2, 6, 13, 13, 5, 7, 4, 1, 5, 2, 2, 5, 4, 10, 2, 4, 6, 5, 3, 3, 4, 4, 4, 2, 1, 1, 3, 4, 5, 3, 4, 2, 9, 13, 13, 13, 11, 5, 4, 4, 3, 3, 2, 2, 2, 3, 3, 5, 5, 5, 8, 3, 3, 3, 7, 8, 9, 1, 1, 1, 3, 3},
 {3, 3, 2, 2, 2, 0, 1, 3, 3, 2, 2, 5, 3, 5, 2, 2, 2, 3, 3, 6, 2, 2, 10, 11, 8, 13, 13, 13, 7, 2, 2, 5, 3, 3, 9, 0, 7, 5, 4, 6, 5, 0, 2, 2, 3, 6, 7, 2, 3, 11, 13, 13, 4, 6, 3, 7, 4, 13, 13, 1, 1, 1, 3},
 {11, 5, 9, 8, 1, 1, 1, 4, 4, 0, 5, 9, 2, 13, 13, 1, 1, 5, 4, 8, 1, 10, 13, 13, 13, 4, 5, 1, 1, 4, 5, 1, 1, 4, 4, 4, 5, 1, 1, 0, 5, 7, 8, 5, 5, 4, 4, 1, 5, 5, 5, 2, 9, 3, 10, 4, 11, 3, 4, 5, 6, 4, 0},
 {5, 10, 4, 4, 13, 13, 13, 10, 2, 2, 6, 7, 10, 4, 4, 4, 6, 1, 2, 2, 2, 6, 9, 0, 10, 8, 5, 7, 4, 4, 4, 0, 4, 4, 10, 6, 2, 7, 0, 5, 7, 3, 10, 2, 2, 2, 11, 6, 1, 1, 1, 6, 7, 3, 7, 13, 13, 11, 2, 6, 7, 3, 3},
 {9, 2, 4, 9, 13, 13, 13, 5, 5, 9, 11, 3, 8, 9, 3, 11, 8, 1, 8, 5, 4, 3, 8, 9, 4, 8, 3, 3, 3, 11, 10, 5, 5, 1, 1, 1, 5, 5, 10, 8, 3, 11, 6, 2, 9, 6, 1, 10, 8, 5, 5, 5, 11, 6, 1, 7, 10, 2, 8, 13, 13, 7, 1, 1, 1, 6, 3, 3, 3, 6, 5, 7, 3},
}


// 定義轉輪為 3x3x3x3x3
reels_def := []int{3, 3, 3, 3, 3}

// 定位轉輪位置&正式遊戲使用mt19937亂數產生器
// ng_len := []int{len(ngReelStrips[0]), len(ngReelStrips[1]), len(ngReelStrips[2]), len(ngReelStrips[3]), len(ngReelStrips[4])}
// pos := random.Intsn(ng_len)
pos := []int{1, 1, 1, 1, 1}

// 取得 3x3x3x3x3 連續轉輪
spin := ngReelStrips.ContiguousReelStrips(reels_def, pos)
log.Printf("ng spin reels: %+v", spin)

```

output:

```bash
 ng spin reels: [[2 6 13] [3 2 2] [5 9 8] [10 4 4] [2 4 9]]
```

## 顯示輪每輪會多取上下兩個

```go
// 宣告轉輪表，參考上面
// var ngReelStrips = ...


// 定義轉輪為 3x3x3x3x3
reels_def := []int{3, 3, 3, 3, 3}

// 定位轉輪位置&正式遊戲使用mt19937亂數產生器
pos := []int{1, 1, 1, 1, 1}

// 取得 3x3x3x3x3 連續轉輪
spin := ngReelStrips.ShowReelStrips(reels_def, pos)
log.Printf("ng spin show reels: %+v", spin)
log.Printf("ng spin show reels反轉: %+v", spin.InvertRegularXYAxis())
```

output:

```bash
ng spin show reels: [[2 2 6 13 13] [3 3 2 2 2] [11 5 9 8 1] [5 10 4 4 13] [9 2 4 9 13]]
ng spin show reels反轉: [[2 3 11 5 9] [2 3 5 10 2] [6 2 9 4 4] [13 2 8 4 9] [13 2 1 13 13]]
```

## 大獎圖

```bash
// 宣告轉輪表，參考上面
// var ngReelStrips = ...


// 定義轉輪為 3x3x3x3x3
reels_def := []int{3, 3, 3, 3, 3}

// 定位轉輪位置&正式遊戲使用mt19937亂數產生器
pos := []int{1, 1, 1, 1, 1}

// 大獎圖，後面 bool 陣列是否為整輪
spin := ngReelStrips.RepeatedReelStrips(reels_set, pos, []bool{false, true, true, true, false})
log.Printf("big symbols: %+v", spin)

// 演出用大獎圖，後面 bool 陣列是否為整輪，最後兩個參數為演出用的轉輪數量
spin = ShowRepeatedReelStrips(reels_set, pos, []bool{false, true, true, true, false}, 1, 1)
log.Printf("big symbols 演出輪: %+v", spin)
```

output:

```bash
big symbols: [[2 6 13] [3 3 3] [5 5 5] [10 10 10] [2 4 9]]
big symbols 演出輪: [[2 2 6 13 13] [3 3 3 3 2] [11 5 5 5 1] [5 10 10 10 13] [9 2 4 9 13]]
```

## 獎圖個數計算

```golang
// 宣告轉輪表，參考上面
// var ngReelStrips = ...


// 定義轉輪為 3x3x3x3x3
reels_def := []int{3, 3, 3, 3, 3}

// 定位轉輪位置&正式遊戲使用mt19937亂數產生器
pos := []int{1, 1, 1, 1, 1}

// 取得 3x3x3x3x3 連續轉輪
spin := ngReelStrips.ContiguousReelStrips(reels_def, pos)
log.Printf("ng spin reels: %+v", spin)

// 計算從左到右得獎圖個數，返回每輪個有幾個目標獎圖，及算作多少way
match, multi := spin.CalcSymbolsMatchFromLeft(slots.H3, slots.WW)
log.Printf("CalcSymbolsMatchFromLeft: 每輪有: %+v，Way: %+v", match, multi)
```

output:

```bash
ng spin reels: [[2 6 13] [3 2 2] [5 9 8] [10 4 4] [2 4 9]]
CalcSymbolsMatchFromLeft: 每輪有: []，Way: 1
```
