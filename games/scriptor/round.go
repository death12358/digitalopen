package scriptor

import (
	"encoding/json"

	"digitalopen/games"
)

// RoundRecords
type RoundRecords struct {
	Status string       `json:"status"`
	Round  games.Rounds `json:"round"`
}

// Marshal
func (r *RoundRecords) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Unmarshal
func (r *RoundRecords) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}
