package character

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/controller"
)

//every Round done, should be able to manually start next Round, after checking the output.
type Synchronizer struct {
	nodeNums           int
	hashDoneInformChan chan bool
	round              int
	console            *controller.Console
	broadCastBlockChan chan *BlockWrap
	summaryChan        chan *RoundSummary
	blockSeq           int
	chain              *block.Tree
}

func NewSynchronizer(nodeNums int, hashDoneInformChan chan bool, broadCastBlockChan chan *BlockWrap, summaryChan chan *RoundSummary, initBlock *block.Node, console *controller.Console) *Synchronizer {
	chain := block.NewTree(initBlock)
	return &Synchronizer{
		nodeNums:           nodeNums,
		hashDoneInformChan: hashDoneInformChan,
		broadCastBlockChan: broadCastBlockChan,
		summaryChan:        summaryChan,
		round:              0,
		blockSeq:           0,
		chain:              chain,
		console:            console,
	}
}

func (s *Synchronizer) MainRoutine() {
	for {
		//every Round
		receiverCnt := 0
		receivedBlock := make([]*BlockWrap, 0)
		for {
			select {
			case block := <-s.broadCastBlockChan:
				receivedBlock = append(receivedBlock, block)
			case <-s.hashDoneInformChan:
				receiverCnt += 1
			}
			if receiverCnt+len(receivedBlock) == s.nodeNums {
				break
				//start next Round

			}
		}
		//send the summary
		//calculate summary
		//should add some controllerable factor...
		s.handleSummary(receivedBlock)
		s.handleControl()
		s.round += 1
	}
}
func (s *Synchronizer) handleControl() {
	s.console.Controlled()
}
func (s *Synchronizer) PrintChain() {
	s.chain.PrintChain()
}
func (s *Synchronizer) handleSummary(wrap []*BlockWrap) {
	if len(wrap) == 0 {
		rs := &RoundSummary{
			RoundType: VOID,
		}
		for i := 0; i < s.nodeNums; i++ {
			s.summaryChan <- rs
		}

	} else {
		//generate block with sequence
		blockList := make([]*block.Node, 0, len(wrap))
		for i := 0; i < len(wrap); i++ {
			nBlock := block.NewBlock(wrap[i].IsEvil, wrap[i].PrevEvil, wrap[i].Prev, wrap[i].Owner, s.round, s.blockSeq)
			s.blockSeq += 1
			blockList = append(blockList, nBlock)
			s.console.AddCount()
			nBlock.PrintBlock()
			s.chain.AddNode(nBlock)
		}
		rs := &RoundSummary{
			RoundType: PROD,
			Blocks:    blockList,
		}
		for i := 0; i < s.nodeNums; i++ {
			s.summaryChan <- rs
		}
	}
}
