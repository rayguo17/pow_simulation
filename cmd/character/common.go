package character

import "github.com/rayguo17/pow/cmd/block"

const (
	VOID int = 1
	PROD     = 2
)

type BlockWrap struct {
	isEvil   bool //tag only for simulation
	prevEvil bool
	prev     *block.Node //from which
	owner    int         //id of node
	round    int         //which round
}

type RoundSummary struct {
	blocks    []*block.Node
	roundType int
}
