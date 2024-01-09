package sg006

import (
	"github.com/death12358/digitalopen/games"
	"github.com/shopspring/decimal"
)

var payTable_unitbet100 = &games.PayTable{
	// WW
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// H1
	games.Pays{decimal.Zero, decimal.NewFromInt(50), decimal.NewFromInt(200), decimal.NewFromInt(900), decimal.NewFromInt(5000)},
	// H2
	games.Pays{decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(100), decimal.NewFromInt(600), decimal.NewFromInt(3000)},
	// H3
	games.Pays{decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(70), decimal.NewFromInt(300), decimal.NewFromInt(2000)},
	// H4
	games.Pays{decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(70), decimal.NewFromInt(300), decimal.NewFromInt(2000)},
	// H5
	games.Pays{decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(60), decimal.NewFromInt(200), decimal.NewFromInt(1500)},
	// LA
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(40), decimal.NewFromInt(150), decimal.NewFromInt(400)},
	// LK
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(40), decimal.NewFromInt(150), decimal.NewFromInt(400)},
	// LQ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(25), decimal.NewFromInt(90), decimal.NewFromInt(300)},
	// LJ
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(25), decimal.NewFromInt(90), decimal.NewFromInt(300)},
	// LT
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(70), decimal.NewFromInt(180)},
	// LN
	games.Pays{decimal.Zero, decimal.Zero, decimal.NewFromInt(20), decimal.NewFromInt(70), decimal.NewFromInt(180)},
	// SE
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	// SF
	//	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
	games.Pays{decimal.Zero, decimal.NewFromInt(100), decimal.NewFromInt(550), decimal.Zero, decimal.Zero},
	// SB
	games.Pays{decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero, decimal.Zero},
}
