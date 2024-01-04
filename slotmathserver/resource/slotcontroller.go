package resource

import (
	"digitalopen/slotmathserver/resource/args"
	"encoding/json"
	"fmt"
	"strings"

	"digitalopen/games"

	"digitalopen/games/scriptor"

	"github.com/shopspring/decimal"
)

// SlotController is a data structure that represents the game flow of a slot machine.
type SlotController struct {
	rtp        map[string]string // RTP 值映射表
	game       *games.Games      // 遊戲對象
	code       string            // 遊戲代碼
	defaultRTP string            // 默認 RTP 值
	jp_def     []string          // JP 前綴列表，可為 nil
	jp_sorts   []string          // JP 排序規則
	jp_mapping map[string]string // JP 映射表
}

// New is a function that returns a pointer to a SlotController object.
func New(rtp map[string]string, game *games.Games, defRTP string) *SlotController {
	sf := &SlotController{
		rtp:        rtp,
		game:       game,
		code:       game.Name(),
		defaultRTP: defRTP,
		jp_def:     nil,
		jp_sorts:   nil,
		jp_mapping: nil,
	}

	return sf
}

// NewJP is a function that returns a pointer to a SlotController object.
// 如果遊戲不包含 jp 前綴，則將 jp_def 參數設為 nil。
func NewJP(rtp map[string]string, game *games.Games, defRTP string, jp_def *[]string, jp_sorts []string, jp_mapping map[string]string) *SlotController {
	sf := &SlotController{
		rtp:        rtp,
		game:       game,
		code:       game.Name(),
		defaultRTP: defRTP,
		jp_sorts:   jp_sorts,
		jp_mapping: jp_mapping,
	}

	// 只有在 jp_def 不為 nil 時，才設置 jp_def 屬性
	if jp_def != nil {
		sf.jp_def = *jp_def
	}

	return sf
}

// getRTP is a function that returns the RTP value for a given currency.
// It takes a currency string and a default RTP value as parameters.
// The function returns the RTP value for the given currency if it exists,
// otherwise it returns the default RTP value.
func (sf *SlotController) getRTP(currency string) string {
	if rtp, ok := sf.rtp[strings.ToUpper(currency)]; ok {
		return rtp
	}

	return sf.defaultRTP
}

// selectGameInSession selects a game in a session and returns the updated round.
func (sf *SlotController) selectGameInSession(cached *scriptor.RoundRecords, rtp string, bet decimal.Decimal, argv *args.SpinArgs) (*games.Rounds, error) {
	roundPos := cached.Round.Position
	stages := cached.Round.Stages

	// Check if the round is in a rest position and not in a not started yet position.
	// If so, spin the game and update the round in the cache.
	if roundPos.IsRest() && !roundPos.IsNotStartedYet() {
		// Spin the game with the given parameters and update the stages and position fields in the round.
		round, err := sf.game.Spin(rtp, bet, argv.Pickem, cached.Round)
		if err != nil {
			return nil, err
		}
		round.Stages = stages + 1
		round.Position = round.Position.Pop(games.Rest).Pop(games.NotStartedYet)

		// Update the round in the cache and remove the not started flag.
		cached, err = slot_Scriptor.RoundUpdate("0", sf.code, argv.Id, argv.Username, *round)
		if err != nil {
			return nil, err
		}

		cached.Round.RemoveNotStarted()
		return &cached.Round, nil
	}

	// If the round is not in a rest position or is in a not started yet position,
	// update the position field in the round and remove the not started flag.
	cached.Round.Position = cached.Round.Position.Pop(games.NotStartedYet)
	cached.Round.RemoveNotStarted()
	return &cached.Round, nil
}

// createNewSession creates a new session and returns the created session.
func (sf *SlotController) createNewSession(argv args.SpinArgs, rtp string, bet decimal.Decimal) (*games.Rounds, error) {
	// Create a new rounds object with the necessary fields.
	round := &games.Rounds{
		Id:       argv.Id,
		GameCode: sf.code,
		Brand:    argv.Brand,
		Username: argv.Username,
		Status:   games.State(0),
		Position: games.State(0),
		Stages:   0,
		Result:   games.NewResults(),
		Currency: argv.Currency,
	}

	// Spin the game with the given parameters.
	round, err := sf.game.Spin(rtp, bet, argv.Pickem, *round)
	if err != nil {
		return nil, err
	}

	// If the round is not started yet, remove the not started flag.
	if round.Position.IsNone() {
		round.RemoveNotStarted()
		return round, nil
	}

	// Set the position to not started yet, reset the stages to zero, and update the round in the cache.
	round.Position = round.Position.Pop(games.NotStartedYet)
	round.Stages = 0
	cached, err := slot_Scriptor.RoundUpdate("0", sf.code, argv.Id, argv.Username, *round)
	if err != nil {
		return nil, err
	}

	// Update the round object with the cached round, and remove the not started flag.
	round = &cached.Round
	round.RemoveNotStarted()
	return round, nil
}

// ExecuteSpin handles the spin flow for a slot game. It takes in the spin arguments and returns a pointer to a Rounds object and an error.
func (sf *SlotController) ExecuteSpin(argv args.SpinArgs) (*games.Rounds, error) {
	// Get the RTP value based on the currency.
	rtp := sf.getRTP(argv.Currency)
	// Get the cached round from the cache.
	cached, err := slot_Scriptor.RoundNext("0", sf.code, argv.Id, argv.Username)
	if err != nil {
		return nil, err
	}

	// Convert the bet value to a decimal.
	bet, err := decimal.NewFromString(argv.Bet.String())
	if err != nil {
		return nil, err
	}

	// Determine what action to take based on the status of the cached round.
	switch cached.Status {
	// If the cached round exists, select a game in the session and return the updated round.
	case scriptor.OK:
		round, err := sf.selectGameInSession(cached, rtp, bet, &argv)
		if err != nil {
			return nil, err
		}
		return round, nil

	// If the cached round does not exist, create a new session and return the new round.
	case scriptor.NOT_EXIST:
		round, err := sf.createNewSession(argv, rtp, bet)
		if err != nil {
			return nil, err
		}
		return round, nil

	// If the cached round has an unknown status, return an error.
	default:
		return nil, fmt.Errorf("UNKNOWN_ROUNDSTATUS_ERROR: %s", cached.Status)
	}
}

// GetJackpot retrieves the current jackpot values for the specified slot game.
// The function takes in the argv args.SpinArgs as input.
// The function returns a slice of decimal.Decimal values representing the jackpot values for each of the defined jackpot levels.
// The function also returns an error if the jackpot levels are not defined, or if there is an issue retrieving the jackpot values from the cache.
func (sf *SlotController) GetJackpot(argv args.SpinArgs) ([]decimal.Decimal, error) {
	// // Check if jackpot levels are defined
	// if sf.jp_def == nil {
	// 	return nil, fmt.Errorf("JACKPOT_NOT_DEFINED")
	// }

	// // Get RTP based on currency
	// rtp := sf.getRTP(argv.Currency)

	// // Convert bet value to decimal
	// bet, err := decimal.NewFromString(argv.Bet.String())
	// if err != nil {
	// 	return nil, err
	// }

	// // Get the base jackpot values for the specified game and bet level
	// basepool := sf.game.BaseJackpot(argv.Bet.String(), rtp, bet)

	// // Retrieve the current jackpot values from the cache
	// cache, err := slot_Scriptor.JackpotPeeks("0", argv.Brand+sf.code, argv.Currency)
	// if err != nil {
	// 	return nil, err
	// }

	// // If no jackpot values are found in the cache, return the base jackpot values
	// if len(cache.Pools) == 0 {
	// 	return basepool, nil
	// }

	// // Add the current jackpot values to the base jackpot values for each jackpot level
	// for i, v := range basepool {
	// 	pool, err := decimal.NewFromString(cache.Pools[sf.jp_def[i]].String())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	basepool[i] = v.Add(pool)
	// }

	return nil, nil
}

// GetJackpotPools is a SlotController method function that retrieves the jackpot pools
// for the game specified in the input arguments (argv). The method returns a
// map of pool values for each jackpot type.
func (sf *SlotController) GetJackpotPools(argv args.PoolArgs) (map[string][]json.Number, error) {
	// // Retrieve the jackpot pools for the game and currency specified in the input arguments.
	// jp, err := slot_Scriptor.JackpotCurrencyPeeks("0", argv.Brand+sf.code)
	// if err != nil {
	// 	return nil, err
	// }

	// // Retrieve the pool values for each jackpot type.
	// pools := jp.Pools
	// poolResp := map[string][]json.Number{}
	// if len(pools) != 0 {
	// 	for i, v := range pools {
	// 		// Sort the pool values according to the configured order of jackpot types.
	// 		poolResp[i] = make([]json.Number, 4)
	// 		for v_i, v_v := range sf.jp_sorts {
	// 			poolResp[i][v_i] = v[v_v]
	// 		}
	// 	}
	// }

	return nil, nil
}

func (sf *SlotController) PushJackpot(unitbet string, jpname string, currency string, totalBet decimal.Decimal) error {
	// rtp := sf.getRTP(currency)
	// jp_roll_rate := sf.game.CalcRollRate(unitbet, rtp, totalBet)

	// for i, v := range sf.jp_def {
	// 	_, err := slot_Scriptor.JackpotPush("0", jpname, currency, v, jp_roll_rate[i])
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// JPCalc calculates the jackpot points and updates the rounds accordingly.
func (sf *SlotController) JPCalc(rounds *games.Rounds) (*games.Rounds, error) {

	return nil, nil
}

// JackpotSweep sweeps the jackpot pool.
func (sf *SlotController) JackpotSweep() (*scriptor.JackpotPool, error) {

	return nil, nil
}
