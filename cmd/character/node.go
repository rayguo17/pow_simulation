package character

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/common"
	"log"
	"math/rand"
)

const (
	Clean       = 0
	Sybil   int = 1 //51% attack
	Selfish     = 2 //Selfish mining
)

type Node struct {
	probability        float64
	id                 int
	broadcastChan      chan *common.BlockWrap //when block mined broadcast
	isEvil             bool                   //behave differently
	hashRound          int                    //how many hash function could be done in one round?
	roundEndChan       chan *common.RoundSummary
	hashDoneInformChan chan bool //tell sychronizer done with no block associated
	round              int
	random             *rand.Rand //not safe for goroutine so each different
	chain              *block.Tree
	evilTarget         int //if <-2 do nothing
	//selfish mining and 51% attack cannot be implement at the same time!
	evilMode int
	nodeNum  int
}

func NewNode(id int, broadcastChan chan *common.BlockWrap, probability float64, isEvil bool, hashRound int, roundEndChan chan *common.RoundSummary, hashDoneInformChan chan bool, random *rand.Rand, initBlock *block.Node, evilMode int, nodeNum int) *Node {
	chain := block.NewTree(initBlock)
	return &Node{
		id:                 id,
		broadcastChan:      broadcastChan,
		probability:        probability,
		isEvil:             isEvil,
		hashRound:          hashRound,
		roundEndChan:       roundEndChan,
		round:              0,
		hashDoneInformChan: hashDoneInformChan,
		random:             random,
		chain:              chain,
		evilTarget:         -2,
		evilMode:           evilMode,
		nodeNum:            nodeNum,
	}
}

//evil node have two type of attack:
//1. selfish mining
//2. branch diver

//should be able to control params... especially in the case of building.
//every Round the probability of having new Blocks...
//every Round could access q times hash, everytime hash has a probability... -> relate to difficulty
func (n *Node) MainRoutine() {

	//for every Round listen to different chain and also generate own block,
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
func (n *Node) handleSum(sum *common.RoundSummary) {
	if n.isEvil {
		//do parse summary
		n.evilTarget = sum.EvilTargetSeq
	}
	if sum.RoundType == common.VOID {
		return
	} else {
		for i := 0; i < len(sum.Blocks); i++ {
			n.chain.AddNode(sum.Blocks[i])
		}
	}
}
func (n *Node) packageBlock(prevBlock *block.Node) {
	//create new block
	purpose := false
	evilTarget := -2
	if n.isEvil && n.evilMode == Sybil && n.evilTarget >= -1 {
		purpose = true
		evilTarget = n.evilTarget
	} else {
		evilTarget = prevBlock.GetEvilTarget()
	}
	bw := &common.BlockWrap{
		Owner:         n.id,
		IsFromEvil:    n.isEvil,
		IsEvilPurpose: purpose,
		Prev:          prevBlock,
		PrevEvil:      prevBlock.InheritEvil(),
		Round:         n.round,
		EvilTarget:    evilTarget,
	}
	n.broadcastChan <- bw
}
func (n *Node) calculate() {
	//decide main block here??
	mainBlock := n.getMainBlock()
	if mainBlock == nil {
		log.Fatal("main block should not be empty")
	}
	for i := 0; i < n.hashRound; i++ {
		res := rand.Float64() < n.probability
		if res == true {
			//handle package block
			n.packageBlock(mainBlock)
			return
		}
	}
	//if none of the block could be done then tell the synchronizer
	n.hashDoneInformChan <- true
}

func (n *Node) getMainBlock() *block.Node {
	if n.isEvil && n.evilMode == Sybil && n.evilTarget >= -1 {
		//do something else????
		//find evil leaves or evil target
		bl := n.chain.FindEvilTarget(n.evilTarget)
		return bl
	}
	bl := n.chain.GetLongestChain()
	if len(bl) == 1 {
		return bl[0]
	}
	if len(bl) > 1 {
		nums := len(bl)
		index := n.random.Intn(nums)
		return bl[index]
	}
	return nil
}
func (n *Node) IsEvil() bool {
	return n.isEvil
}
