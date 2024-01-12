package bbhz

// import (
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bbhz/constants"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bbhz/prtable"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/internal"
// 	pb_game "gitlab.com/bf_tech/template/slot/gametemplate/proto/game"
// )

// var bbhz *bbhzControl

// type bbhzControl struct {
// 	internal.BaseGameControl[*prtable.PRTable]
// }

// func GetController() *bbhzControl {

// 	if bbhz == nil {
// 		bbhz = &bbhzControl{}
// 		bbhz.InitController()
// 	}

// 	return bbhz
// }

// func (c *bbhzControl) InitController() error {

// 	bbhz.BaseGameControl = internal.NewBGC[*prtable.PRTable]()

// 	if err := c.InitPay(constants.FolderName, int(constants.GameType)); err != nil {
// 		return err
// 	}

// 	if err := c.InitProbabilitySetting(constants.FolderName); err != nil {
// 		return err
// 	}

// 	for _, file := range c.ProbabilitySettings {
// 		if err := c.InsertProbabilityTable(constants.FolderName, file, &prtable.PRTable{}); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (c *bbhzControl) Spin(inRequest *pb_game.SlotGameRequest) (*pb_game.SlotGameResults, error) {

// 	var (
// 		slotGameResult *pb_game.SlotGameResults
// 		spinErr        error
// 	)

// 	//抓取機率表
// 	probabilityName, prId, prtable, err := c.GetProbabilityTable(int(inRequest.GetProbInfo().ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	respinTimes := 0
// 	probabilityLimit := prtable.Limit.CalculateMergedProbabilityLimit(int(inRequest.ProbInfo.Upper))
// 	for isOverProbabilityLimit := true; isOverProbabilityLimit; {

// 		if respinTimes >= internal.LimitRespinTimes { //到達重開次數
// 			probabilityName, prId, prtable, err = c.GetProbabilityTable(internal.ZeroProbTableID)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}

// 		slotGameResult, spinErr = c.processMainGame(prtable, inRequest)

// 		if inRequest.GetBetInfo().Bet != 0 {
// 			winM := slotGameResult.TotalWin / int64(inRequest.GetBetInfo().Bet) //得分倍數
// 			isOverProbabilityLimit = probabilityLimit.IsOutOfLimit(int(winM))
// 		}

// 		if respinTimes >= internal.LimitRespinTimes {
// 			break
// 		}

// 		respinTimes++
// 	}

// 	slotGameResult.UsedProbabilityName = probabilityName
// 	slotGameResult.UsedProbabilityID = int32(prId)

// 	return slotGameResult, spinErr
// }

// // GetPay() 取賠付資訊
// func (c *bbhzControl) GetPay() *pb_game.PayData {
// 	payResult := c.Pay.ConvPBPayData()
// 	return payResult
// }
// func (c *bbhzControl) GetProbabilitySetting() map[int]string {
// 	return c.ProbabilitySettings
// }
