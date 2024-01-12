package internal

const (
	Wild     = "W"
	Scatter  = "Sc"
	Mystery  = "M"
	Interval = "I"
	H1       = "H1"
)

type Folder string

type PrizeLevel int

const (
	Mini PrizeLevel = iota
	Minor
	Major
	Grand
)

type GameStand int

const (
	Line GameStand = iota + 1
	Way
)

const (
	LimitRespinTimes = 10
	ZeroProbTableID  = 0
)
