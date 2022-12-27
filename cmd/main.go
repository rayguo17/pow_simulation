package main

import (
	"bufio"
	"fmt"
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/character"
	"github.com/rayguo17/pow/cmd/common"
	"github.com/rayguo17/pow/cmd/controller"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	nodeNum := 10
	evilNum := 10
	difficulty := 0.0005
	hashRound := 10
	evilMode := character.Sybil
	broadcastChan := make(chan *common.BlockWrap)
	roundEndChan := make(chan *common.RoundSummary)
	hashDoneInformChan := make(chan bool)
	nodeList := make([]*character.Node, 0)

	initBlock := block.NewBlock(false, false, nil, -1, -1, -1, false, -2)
	for i := 0; i < nodeNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		node := character.NewNode(i, broadcastChan, difficulty, false, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock, character.Clean, nodeNum+evilNum)
		nodeList = append(nodeList, node)
		go node.MainRoutine()
	}
	for i := 0; i < evilNum; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		evilNode := character.NewNode(nodeNum+i, broadcastChan, difficulty, true, hashRound, roundEndChan, hashDoneInformChan, r1, initBlock, evilMode, nodeNum+evilNum)
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
			console.Config(5, -3)
		case "2\n":
			sync.PrintChain()
		case "3\n":
			tB := sync.GetLatestPrevBlock()
			if tB == nil {
				log.Fatal("latest block nil")
			}
			console.Config(5, tB.GetSeq())
		case "4\n":
			return
		default:
			num, err := strconv.Atoi(strings.Trim(str, "\n"))
			if err != nil {
				fmt.Println(err)
			}
			bl := sync.GetBlockByIndex(num)
			bl.PrintBlock()
		}

	}

}
