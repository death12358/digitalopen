package game

// import (
// 	"fmt"
// 	"strconv"
// 	"sync"

// 	shared_game "github.com/adimax2953/Shared/backend/game"
// 	"github.com/adimax2953/Shared/shared/gamecode"
// 	LogTool "github.com/adimax2953/log-tool"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bbhz"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bxnw"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/bxnw2"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/dfdc"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/dscj"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/sgpd"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/controller/game/thh"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/currency"
// 	"gitlab.com/bf_tech/template/slot/gametemplate/pkg/slotgames/slotrtp"
// 	pb_game "gitlab.com/bf_tech/template/slot/gametemplate/proto/game"
// )

// var allGameControllerMap sync.Map

// type IGameControl interface {
// 	InitController() error
// 	Spin(inRequest *pb_game.SlotGameRequest) (*pb_game.SlotGameResults, error)
// 	GetPay() *pb_game.PayData
// 	GetProbabilitySetting() map[int]string
// }

// func registerController() {
// 	allGameControllerMap.Store(gamecode.DFDC.String(), dfdc.GetController())
// 	allGameControllerMap.Store(gamecode.DSCJ.String(), dscj.GetController())
// 	allGameControllerMap.Store(gamecode.BXNW.String(), bxnw.GetController())
// 	allGameControllerMap.Store(gamecode.BXNW2.String(), bxnw2.GetController())
// 	allGameControllerMap.Store(gamecode.SGPD.String(), sgpd.GetController())
// 	allGameControllerMap.Store(gamecode.THH.String(), thh.GetController())
// 	allGameControllerMap.Store(gamecode.BBHZ.String(), bbhz.GetController())
// }

// func InitAll() {
// 	registerController()
// 	errorMap := make(map[any]string)
// 	allGameControllerMap.Range(func(key, value any) bool {
// 		c, ok := value.(IGameControl)
// 		if !ok {
// 			// LogTool.LogWarningf("InitAll", "Controller init fail. ID:[%v]", key)
// 			errorMap[key] = fmt.Sprintf("Controller init fail. ID:[%v]", key)
// 		}
// 		err := c.InitController()
// 		if err != nil {
// 			// LogTool.LogWarningf("InitAll", "Controller init fail. ID:[%v], %v", key, err)
// 			errorMap[key] = fmt.Sprintf("Controller init fail. ID:[%v], %v", key, err)
// 		}
// 		LogTool.LogInfof("InitAll", "GameCode:[%s], init ok!!", key)
// 		return true
// 	})

// 	if len(errorMap) != 0 {
// 		for key, msg := range errorMap {
// 			LogTool.LogInfof("InitAll", "GameCode:[%s], init fail. error:[%v]", key, msg)
// 		}
// 	}
// }

// func LogProbabilityMapping() {

// 	allGameControllerMap.Range(func(key, value any) bool {
// 		gameCode := key.(string)
// 		shared_game.GetInstance().ClearSlotGameMappingLog(gameCode)

// 		c, _ := value.(IGameControl)
// 		mappingData := c.GetProbabilitySetting()
// 		argsMap := make(map[string]interface{})
// 		for k, v := range mappingData {
// 			argsMap[strconv.Itoa(k)] = v
// 		}
// 		shared_game.GetInstance().SetSlotGameMappingLog(gameCode, argsMap)
// 		return true
// 	})
// }

// // InitSlotGame for online
// func InitSlotGame() {
// 	slotrtp.InitSlotRTP()
// 	currency.InitCurrencyConfig()
// 	InitAll()
// 	LogProbabilityMapping()
// }

// func GetGameControl(inGameID string) IGameControl {
// 	controller, ok := allGameControllerMap.Load(inGameID)
// 	if ok {
// 		return controller.(IGameControl)
// 	}
// 	return nil
// }
