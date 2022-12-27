package character

import (
	"github.com/rayguo17/pow/cmd/block"
	"math/rand"
)

type Node struct {
	probability         float64
	id                  int
	receiveBlockChan    chan *block.Node
	broadcastChan       chan *BlockWrap //when block mined broadcast
	isEvil              bool            //behave differently
	hashRound           int             //how many hash function could be done in one round?
	roundEndChan        chan *RoundSummary
	informRoundDoneChan chan bool //tell sychronizer done with no block associated
	round               int
	blockStorage        []*Block   //storage data structure should be different
	random              *rand.Rand //not safe for goroutine so each different
	chain               *block.Tree
}

func NewNode(id int, receiveBlockChan chan *block.Node, broadcastChan chan *BlockWrap, probability float64, isEvil bool, hashRound int, roundEndChan chan *RoundSummary, informRoundDoneChan chan bool, random *rand.Rand, initBlock *block.Node) *Node {
	chain := block.NewTree(initBlock)
	return &Node{
		id:                  id,
		receiveBlockChan:    receiveBlockChan,
		broadcastChan:       broadcastChan,
		probability:         probability,
		isEvil:              isEvil,
		hashRound:           hashRound,
		roundEndChan:        roundEndChan,
		round:               0,
		informRoundDoneChan: informRoundDoneChan,
		random:              random,
		chain:               chain,
	}
}

//evil node have two type of attack:
//1. selfish mining
//2. branch diver

//should be able to control params... especially in the case of building.
//every round the probability of having new blocks...
//every round could access q times hash, everytime hash has a probability... -> relate to difficulty
func (n *Node) MainRoutine() {

	//for every round listen to different chain and also generate own block,
	for {
		//do calculation routine
		n.calculate()
		select {
		case sum := <-n.roundEndChan:
			n.handleSum(sum)
		}
		n.round++
	}
}

/*summary have
1. nobody create block
2. have one block created
3. have multiple block created: then we will just randomly accept one (if both valid)

question:
1. should we have a "errCount" to indicate we are how many block behind it??? or we accept both???
2. how does normal node react to invalid or previously valid block.
3. how does each node store block tree?
*/
func (n *Node) handleSum(sum *RoundSummary) {

}
func (n *Node) packageBlock() {
	//create new block
}
func (n *Node) calculate() {
	for i := 0; i < n.hashRound; i++ {
		res := rand.Float64() < n.probability
		if res == true {
			//handle package block
			n.packageBlock()
		}
	}
	//if none of the block could be done then tell the synchronizer
}
func (n *Node) verifyBlock(b *Block) {

}
