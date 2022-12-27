package test

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/character"
	"github.com/rayguo17/pow/cmd/common"
	"github.com/rayguo17/pow/cmd/controller"
	"math/rand"
	"testing"
	"time"
)

func TestOneNode(t *testing.T) {
	nodeNum := 1
	evilNum := 0
	difficulty := 0.5
	hashRound := 1
	broadcastChan := make(chan *common.BlockWrap)
	roundEndChan := make(chan *common.RoundSummary)
	hashDoneInformChan := make(chan bool)
	nodeList := make([]*character.Node, 0)
	initBlock := block.NewBlock(false, false, nil, -1, -1, -1, false, -2)
	for i := 0; i < nodeNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		node := character.NewNode(i, broadcastChan, difficulty, false, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock, 0, nodeNum+evilNum)
		nodeList = append(nodeList, node)
		go node.MainRoutine()
	}
	for i := 0; i < evilNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		evilNode := character.NewNode(nodeNum+i, broadcastChan, difficulty, true, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock, character.Clean, nodeNum+evilNum)
		nodeList = append(nodeList, evilNode)
		go evilNode.MainRoutine()
	}
	//run
	console := controller.NewConsole()

	sync := character.NewSynchronizer(nodeNum+evilNum, hashDoneInformChan, broadcastChan, roundEndChan, initBlock, console)
	go sync.MainRoutine()
	select {}
}
