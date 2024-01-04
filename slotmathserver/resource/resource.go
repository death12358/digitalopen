package resource

import (
	"log"

	"github.com/shopspring/decimal"
	. "gitlab.com/gaas_module/games/scriptor"
)

func init() {
	decimal.DivisionPrecision = 16
}

var slot_Scriptor *Scriptor

func NewSlotScriptor(opts *Option) {
	var err error
	slot_Scriptor, err = NewScriptor(opts)
	if err != nil {
		// panic(err.Error(), opts)
		log.Panicf("can't connect: %v, error is %v", opts.Host, err.Error())
	}
	log.Printf("SlotScriptor initialized connect: %v:%v", opts.Host, opts.Port)
}
