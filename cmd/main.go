package main

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/character"
	"math/rand"
	"time"
)

func main() {
	nodeNum := 20
	evilNum := 5
	difficulty := 0.00005
	hashRound := 10
	receiveBlockChan := make(chan *block.Node)
	broadcastChan := make(chan *character.BlockWrap)
	roundEndChan := make(chan *character.RoundSummary)
	hashDoneInformChan := make(chan bool)
	nodeList := make([]*character.Node, 0)
	initBlock := block.NewBlock(false, false, nil, -1, -1, -1)
	for i := 0; i < nodeNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		node := character.NewNode(i, receiveBlockChan, broadcastChan, difficulty, false, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock)
		nodeList = append(nodeList, node)
		go node.MainRoutine()

	}
	for i := 0; i < evilNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		evilNode := character.NewNode(nodeNum+i, receiveBlockChan, broadcastChan, difficulty, true, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock)
		nodeList = append(nodeList, evilNode)
		go evilNode.MainRoutine()
	}
	//run
	sync := character.NewSynchronizer(nodeNum+evilNum, hashDoneInformChan, receiveBlockChan, broadcastChan, roundEndChan, initBlock)
	go sync.MainRoutine()
	select {}
}
