package character

import "github.com/rayguo17/pow/cmd/block"

const (
	VOID int = 1
	PROD     = 2
)

type BlockWrap struct {
	IsEvil   bool //tag only for simulation
	PrevEvil bool
	Prev     *block.Node //from which
	Owner    int         //id of node
	Round    int         //which Round
}

type RoundSummary struct {
	Blocks    []*block.Node
	RoundType int
}
