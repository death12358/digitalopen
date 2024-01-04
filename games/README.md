# games

## Install

```bash
go get -u -v .../games
```

---

## Dependency

<details>
<summary><b>github.com/shopspring/decimal</b></summary>

  ⠿ _浮點數計算相關皆使用 decimal 套件，避免 IEEE 754 精度遺失問題_

- **Install**

    ```bash
    go get -u -v github.com/shopspring/decimal
    ```

</details>

---

## Rounds

  ⠿ _遊戲回合。_

  **Rounds 資料格式**:

```go
type Rounds struct {
    Identifier string            // 訂單編號
    GameCode   string            // 遊戲代號
    Brand      string            // 廠商
    User       string            // 使用者
    Status     State             // 遊戲狀態
    Position   State             // 遊戲播放旗標
    Satges     int64             // 遊戲階段
    Results     map[string]Results // 遊戲紀錄
    Currency   string            // 代幣種類
    TotalBet   decimal.Decimal   // 總投注額
    TotalPoint decimal.Decimal   // 總贏分
}
```

  **Results 資料格式**:

```go
type Results struct {
    Identifier string          `json:"identifier"`
    Brand      string          `json:"brand"`
    User       string          `json:"user"`
    Case       State           `json:"case"`
    Satges     int64           `json:"stages"`
    Pickem     []string        `json:"pickem"`
    Symbols    []string        `json:"symbols"`
    Multiplier decimal.Decimal `json:"multiplier"`
    Bet        decimal.Decimal `json:"bet"`
    Point      decimal.Decimal `json:"point"`
}
```
