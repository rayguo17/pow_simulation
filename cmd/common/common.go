package common

import "github.com/rayguo17/pow/cmd/block"

const (
	VOID int = 1
	PROD     = 2
)

type BlockWrap struct {
	IsFromEvil    bool //tag only for simulation
	PrevEvil      bool
	Prev          *block.Node //from which
	Owner         int         //id of node
	Round         int         //which Round
	IsEvilPurpose bool
	EvilTarget    int
}

type RoundSummary struct {
	Blocks        []*block.Node
	RoundType     int
	EvilTargetSeq int // attack target, if <-1 do nothing... else do start to work on it, set target should have different behaviour
}
