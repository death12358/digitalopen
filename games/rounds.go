package games

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

// Rounds - 遊戲回合單
//
//	一個 Round 包含一個下注到最後贏分
//	@Identifier	- 回合識別碼
//	@Brand		- 品牌
//	@User		- 使用者
//	@Status		- 回合狀態
//	@Position	- 回合狀態播放旗標，0 表示未開始、或結束
//	@Satges		- 回合狀態階段或局數
//	@Result		- 回合結果
//	@Currency	- 幣種
//	@Start		- 回合開始時間
//	@Fisish		- 回合結束時間
//	@TotalBet	- 總下注
//	@TotalPoint	- 總贏分
type Rounds struct {
	Id         string          `json:"id"`
	GameCode   string          `json:"game_code"`
	Brand      string          `json:"brand"`
	Username   string          `json:"username"`
	Status     State           `json:"status"`
	Position   State           `json:"position"`
	Stages     int64           `json:"stages"`
	Result     Results         `json:"result"`
	Currency   string          `json:"currency"`
	Start      int64           `json:"start"`
	Fisish     int64           `json:"finish"`
	TotalBet   decimal.Decimal `json:"total_bet"`
	TotalPoint decimal.Decimal `json:"total_point"`
}

// NewRounds - 建立回合單
func NewRounds() *Rounds {
	return &Rounds{
		Id:         "",
		GameCode:   "",
		Brand:      "",
		Username:   "",
		Status:     State(0),
		Position:   State(0),
		Stages:     0,
		Result:     NewResults(),
		Currency:   "",
		Start:      0,
		Fisish:     0,
		TotalBet:   decimal.Zero,
		TotalPoint: decimal.Zero,
	}
}

// Marshal
func (r Rounds) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *Rounds) Unmarshal(roundsBytes []byte) error {
	return json.Unmarshal(roundsBytes, r)
}

// RemoveNotStarted - Remove case NotStartedYet
func (r *Rounds) RemoveNotStarted() {
	for k, v := range r.Result {
		if v.Case.IsNotStartedYet() {
			delete(r.Result, k)
		}
	}
}

// RemoveStageNotStarted - Remove stages NotStartedYet
func (r *Rounds) RemoveStageNotStarted() {
	for k, v := range r.Result {
		if State(v.Stages).IsNotStartedYet() {
			delete(r.Result, k)
		}
	}
}
