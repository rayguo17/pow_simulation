package test

import (
	"github.com/k0kubun/pp/v3"
	"github.com/rayguo17/pow/cmd/block"
	"testing"
)

func TestStorage(t *testing.T) {
	initBlock := block.NewBlock(false, false, nil, -1, -1, -1)
	tree := block.NewTree(initBlock)
	prevBlock := initBlock
	for i := 0; i < 10; i++ {
		tmpBlock := block.NewBlock(false, false, prevBlock, 0, i, i)
		tree.AddNode(tmpBlock)
		prevBlock = tmpBlock
	}
	for i := 0; i < 3; i++ {
		prevBlock = prevBlock.GetPrevBlock()
	}
	for i := 0; i < 5; i++ {
		tmpBlock := block.NewBlock(false, false, prevBlock, 0, 10+i, 10+i)
		tree.AddNode(tmpBlock)
		prevBlock = tmpBlock
	}
	longestBlock := tree.GetLongestChain()
	pp.Println(longestBlock)
	t.Log("Done")
}
