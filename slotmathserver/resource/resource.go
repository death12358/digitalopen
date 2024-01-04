package resource

import (
	"digitalopen/games/scriptor"
	"log"

	"github.com/shopspring/decimal"
)

func init() {
	decimal.DivisionPrecision = 16
}

var slot_Scriptor *scriptor.Scriptor

func NewSlotScriptor(opts *scriptor.Option) {
	var err error
	slot_Scriptor, err = scriptor.NewScriptor(opts)
	if err != nil {
		// panic(err.Error(), opts)
		log.Panicf("can't connect: %v, error is %v", opts.Host, err.Error())
	}
	log.Printf("SlotScriptor initialized connect: %v:%v", opts.Host, opts.Port)
}
