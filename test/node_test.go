package test

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/character"
	"github.com/rayguo17/pow/cmd/common"
	"math/rand"
	"testing"
)

func TestNode(t *testing.T) {
	nodeNum := 25
	hashRound := 10
	probility := 0.000015

	node := character.NewNode(0, make(chan *common.BlockWrap), probility, false, hashRound, make(chan *common.RoundSummary), make(chan bool), rand.New(rand.NewSource(1)), block.NewBlock(false, false, nil, -1, -1, -1, false, -2), 0, nodeNum)
	node.CalExpectation()
}
