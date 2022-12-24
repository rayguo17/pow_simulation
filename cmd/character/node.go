package character

import "math/rand"

type Node struct {
	probability             float64
	id                      int
	receiveBlockChan        chan *BlockWrap
	calculateEndRoutineChan chan bool
	broadcastChan           chan *BlockWrap //when block mined broadcast
	isEvil                  bool            //behave differently
	nodeNums                int             // number of node in the network
	hashRound               int             //how many hash function could be done in one round?
	stopCalChan             chan bool
	calSuccessChan          chan bool
	roundEndChan            chan *RoundSummary
	informRoundDoneChan     chan bool
	round                   int
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
func (n *Node) handleSum(sum *RoundSummary) {

}
func (n *Node) packageBlock() {
	//create new block
}
func (n *Node) calculate() {
	for i := 0; i < n.hashRound; i++ {
		res := rand.Float64() < n.probability
		if res == true {
			n.calSuccessChan <- true
		}
	}
	//if none of the block could be done then tell the synchronizer
}
func (n *Node) verifyBlock(b *Block) {

}
