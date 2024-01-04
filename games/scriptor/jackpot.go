package scriptor

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type JackpotPool struct {
	Status       string          `json:"status"`
	Jackpot_Pool decimal.Decimal `json:"jackpot_pool"`
}

// Marshal
func (r *JackpotPool) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *JackpotPool) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}

// JackpotPools -
type JackpotPools struct {
	Status string                 `json:"status"`
	Pools  map[string]json.Number `json:"pools"`
}

// Marshal
func (r *JackpotPools) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *JackpotPools) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}

// JackpotCurrencyPools -
//  有夠醜，待優化
type JackpotCurrencyPools struct {
	Status string                            `json:"status"`
	Pools  map[string]map[string]json.Number `json:"pools"`
}

// Marshal
func (r *JackpotCurrencyPools) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *JackpotCurrencyPools) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}
