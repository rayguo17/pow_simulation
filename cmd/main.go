package main

import (
	"bufio"
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/character"
	"github.com/rayguo17/pow/cmd/controller"
	"math/rand"
	"os"
	"time"
)

func main() {
	nodeNum := 5
	evilNum := 0
	difficulty := 0.005
	hashRound := 10

	broadcastChan := make(chan *character.BlockWrap)
	roundEndChan := make(chan *character.RoundSummary)
	hashDoneInformChan := make(chan bool)
	nodeList := make([]*character.Node, 0)

	initBlock := block.NewBlock(false, false, nil, -1, -1, -1)
	for i := 0; i < nodeNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())

		r1 := rand.New(s1)
		node := character.NewNode(i, broadcastChan, difficulty, false, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock)
		nodeList = append(nodeList, node)
		go node.MainRoutine()
	}
	for i := 0; i < evilNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		evilNode := character.NewNode(nodeNum+i, broadcastChan, difficulty, true, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock)
		nodeList = append(nodeList, evilNode)
		go evilNode.MainRoutine()
	}
	//run
	console := controller.NewConsole()
	sync := character.NewSynchronizer(nodeNum+evilNum, hashDoneInformChan, broadcastChan, roundEndChan, initBlock, console)
	go sync.MainRoutine()

	//cmdline input
	inputReader := bufio.NewReader(os.Stdin)
	for {
		str, _ := inputReader.ReadString('\n')
		switch str {
		case "1\n":
			console.Config(5)
		case "2\n":
			sync.PrintChain()
		case "3\n":
			return
		}

	}

}
