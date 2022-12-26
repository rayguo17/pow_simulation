package main

import (
	"github.com/rayguo17/pow/cmd/character"
)

func main() {
	nodeNum := 20
	evilNum := 5
	difficulty := 0.00005
	hashRound := 10
	receiveBlockChan := make(chan *character.Block)
	broadcastChan := make(chan *character.BlockWrap)
	roundEndChan := make(chan *character.RoundSummary)
	hashDoneInformChan := make(chan bool)
	nodeList := make([]*character.Node, 0)
	for i := 0; i < nodeNum; i++ {
		node := character.NewNode(i, receiveBlockChan, broadcastChan, difficulty, false, hashRound, roundEndChan, hashDoneInformChan)
		nodeList = append(nodeList, node)
		go node.MainRoutine()

	}
	for i := 0; i < evilNum; i++ {
		evilNode := character.NewNode(nodeNum+i, receiveBlockChan, broadcastChan, difficulty, true, hashRound, roundEndChan, hashDoneInformChan)
		nodeList = append(nodeList, evilNode)
		go evilNode.MainRoutine()
	}
	//run
	sync := character.NewSynchronizer(nodeNum+evilNum, hashDoneInformChan, receiveBlockChan, broadcastChan, roundEndChan)
	go sync.MainRoutine()
	select {}
}
