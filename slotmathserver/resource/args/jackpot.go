package args

import (
	"encoding/json"

	"github.com/Bofry/arg"
)

type PoolArgs struct {
	Brand string `json:"brand"`
}

// Validate validates the SpinArgs
func (v *PoolArgs) Validate() error {
	return arg.Assert(
		arg.Strings.NonEmpty(v.Brand, "brand"),
	)
}

type JackpotArgs struct {
	Brand    string      `json:"brand"`
	Username string      `json:"username"`
	Currency string      `json:"currency"`
	Bet      json.Number `json:"bet"`
}

// Validate validates the SpinArgs
func (v *JackpotArgs) Validate() error {
	// return assert.Assert(
	// 	assert.NonEmptyString(v.Brand, "brand"),
	// 	assert.NonEmptyString(v.Username, "username"),
	// 	assert.JsonNumber(v.Bet, "bet",
	// 		assert.NonNanNorInf,
	// 		assert.NonNegativeNumber,
	// 	),
	// 	assert.NonEmptyString(v.Currency, "currency"),
	// )
	return nil
}
