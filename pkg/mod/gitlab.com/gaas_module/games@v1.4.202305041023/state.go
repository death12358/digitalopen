package games

// State - The state of the game.
//  遊戲狀態使用 bitwise 來表示，且可複合多個狀態，再複合的狀態中，最高位元的狀態優先。
//  @None		- 0x0000 無中獎
//  @Win		- 0x0001 中獎
//  @Draw		- 0x0002 和局
//  @Lose		- 0x0004 輸
//  @Rest		- 0x0008 休息 or 跳過 Or 未開始
//  @FreeGame	- 0x0010 免費遊戲
//  @Bonus		- 0x0020 獎勵遊戲
//  @Jackpot		- 0x0040 特別獎
type State uint16

// Game State - 遊戲狀態
const (
	None     = State(0x0000)
	Win      = State(0x0001)
	Draw     = State(0x0002)
	Lose     = State(0x0004)
	Skip     = State(0x0008)
	Rest     = State(0x0010)
	FreeGame = State(0x0020)
	Bonus    = State(0x0040)
	Jackpot  = State(0x0080)

	BonusGame1 = State(0x0100)
	BonusGame2 = State(0x0200)
	BonusGame3 = State(0x0400)
	BonusGame4 = State(0x0800)
	
	FreeGame1 = State(0x1000)
	FreeGame2 = State(0x2000)
	// Unknown7 = State(0x4000)
	NotStartedYet = State(0x8000)
)

// Push -
// Push the state.
func (s State) Push(state State) State {
	return s | state
}

// Pop -
// Pop the state.
func (s State) Pop(state State) State {
	return s &^ state
}

// HighestOrderSet
// Get the highest order set of the state.
func (s State) HighestOrderSet() State {
	// Benchmark_HighestOrderSet-12    	59311336	        20.06 ns/op	       0 B/op	       0 allocs/op
	// if s == 0 {
	// 	return 0
	// }
	// return 1 << uint16(math.Log2(float64(s)))

	// 0xFFFF: Benchmark_HighestOrderSet-12    	1000000000	         0.2579 ns/op	       0 B/op	       0 allocs/op
	// 0x0001: Benchmark_HighestOrderSet-12    	179729894	         6.567 ns/op	       0 B/op	       0 allocs/op
	if s == 0 {
		return State(0x0000)
	}
	for h := State(0x8000); h > 0; h = h >> 1 {
		sh := s & h
		if sh > 0 {
			return sh
		}
	}
	return State(0x0000)
}

// IsNone -
// Check if the state is a none.
func (s State) IsNone() bool {
	return s == None
}

// IsWin -
// Check if the state is a win.
func (s State) IsWin() bool {
	return (s & Win) == Win
}

// IsDraw -
// Check if the state is a draw.
func (s State) IsDraw() bool {
	return (s & Draw) == Draw
}

// IsLose -
// Check if the state is a lose.
func (s State) IsLose() bool {
	return (s & Lose) == Lose
}

// IsRest -
// Check if the state is a rest.
func (s State) IsRest() bool {
	return (s & Rest) == Rest
}

// IsFreeGame -
// Check if the state is a free game.
func (s State) IsFreeGame() bool {
	return (s & FreeGame) == FreeGame
}

// IsBonus -
// Check if the state is a bonus.
func (s State) IsBonus() bool {
	return (s & Bonus) == Bonus
}

// IsJackpot -
// Check if the state is a jackpot.
func (s State) IsJackpot() bool {
	return (s & Jackpot) == Jackpot
}

// IsBonusGame1 -
// Check if the state is a bonus game 1.
func (s State) IsBonusGame1() bool {
	return (s & BonusGame1) == BonusGame1
}

// IsBonusGame2 -
// Check if the state is a bonus game 2.
func (s State) IsBonusGame2() bool {
	return (s & BonusGame2) == BonusGame2
}

// IsBonusGame3 -
// Check if the state is a bonus game 3.
func (s State) IsBonusGame3() bool {
	return (s & BonusGame3) == BonusGame3
}

// IsBonusGame4 -
// Check if the state is a bonus game 4.
func (s State) IsBonusGame4() bool {
	return (s & BonusGame4) == BonusGame4
}

// IsNotStartedYet -
// Check if the state is a not started yet.
func (s State) IsNotStartedYet() bool {
	return (s & NotStartedYet) == NotStartedYet
}
