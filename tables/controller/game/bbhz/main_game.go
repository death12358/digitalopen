package bbhz

// import (
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bbhz/constants"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bbhz/prtable"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/internal"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/tools"
// 	pb_game "gitlab.com/bf_tech/template/slot/gametemplate/proto/game"
// )

// func (c *bbhzControl) processMainGame(prTable *prtable.PRTable, inRequest *pb_game.SlotGameRequest) (*pb_game.SlotGameResults, error) {
// 	var (
// 		slotGameResult pb_game.SlotGameResults
// 		slotEvent      pb_game.SlotResult_Event
// 		lockWild       []int
// 		stackH1        bool
// 		untilWin       bool
// 	)

// 	for isGetReSpin := true; isGetReSpin; {
// 		slotEvent = pb_game.SlotResult_MainEvent
// 		//產生一個盤面
// 		reels, err := tools.BuildReels(prTable.GetGameReels(), constants.ReelsLength, constants.ReelsHeight)
// 		if err != nil {
// 			return &slotGameResult, err
// 		}

// 		//重轉時觸發,將wild lock
// 		for _, col := range lockWild {
// 			for i := 0; i < len(reels[col]); i++ {
// 				reels[col][i] = internal.Wild
// 			}
// 			slotEvent = pb_game.SlotResult_RespinEvent
// 			untilWin = true
// 		}

// 		//特殊規則判斷
// 		for col := 0; col < len(reels); col++ {
// 			H1Count := tools.CalcSliceElementCount(reels[col], internal.H1)
// 			if H1Count == constants.ReelsHeight {
// 				lockWild = append(lockWild, col)
// 				for i := 0; i < len(reels[col]); i++ {
// 					reels[col][i] = internal.Wild
// 				}
// 				stackH1 = true
// 			}
// 		}

// 		reelsPayResult := c.Pay.Calculate(reels, inRequest.BetInfo.BetBase)

// 		slotRslt := pb_game.SlotResult{ //單一盤面該記錄的結果
// 			SlotEvent:  slotEvent,
// 			Reels:      internal.ParseToReelsPB(reels),
// 			PayDetails: internal.ParseToPayDetailPB(reelsPayResult.GetPayDetail()),
// 			// TriggerRound:    0,
// 			// BonusMultiplier: 0,
// 			Win: reelsPayResult.CalcWin(inRequest.GetBetInfo().Bet, inRequest.GetBetInfo().BetBase),
// 			// ExtraWin: int32(bonusPay) * inRequest.Bet,
// 			// CumWin:   0,
// 			// Customize:       []byte{},
// 		}

// 		slotRslt.CumWin = int64(slotRslt.Win) + int64(slotRslt.ExtraWin)

// 		slotGameResult.TotalWin += slotRslt.CumWin
// 		slotGameResult.AllSlotResults = append(slotGameResult.AllSlotResults, &slotRslt)

// 		if untilWin && slotRslt.CumWin > 0 {
// 			isGetReSpin = false
// 		}

// 		//全盤面Wild
// 		if len(lockWild) == constants.ReelsLength || !untilWin && !stackH1 {
// 			isGetReSpin = false
// 		}
// 	}

// 	return &slotGameResult, nil
// }
