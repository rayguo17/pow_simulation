package character

import (
	"github.com/rayguo17/pow/cmd/block"
	"github.com/rayguo17/pow/cmd/common"
	"github.com/rayguo17/pow/cmd/controller"
	"math/rand"
)

//every Round done, should be able to manually start next Round, after checking the output.
type Synchronizer struct {
	nodeNums           int
	hashDoneInformChan chan bool
	round              int
	console            *controller.Console
	broadCastBlockChan chan *common.BlockWrap
	summaryChan        chan *common.RoundSummary
	blockSeq           int
	chain              *block.Tree
}

func NewSynchronizer(nodeNums int, hashDoneInformChan chan bool, broadCastBlockChan chan *common.BlockWrap, summaryChan chan *common.RoundSummary, initBlock *block.Node, console *controller.Console) *Synchronizer {
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
		receivedBlock := make([]*common.BlockWrap, 0)
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
		rs := s.handleSummary(receivedBlock)
		s.handleControl(rs)
		for i := 0; i < s.nodeNums; i++ {
			s.summaryChan <- rs
		}
		s.round += 1
	}
}
func (s *Synchronizer) GetLatestPrevBlock() *block.Node {
	bl := s.chain.GetLongestChain()
	if len(bl) == 1 {
		return bl[0].GetPrevBlock()
	}
	if len(bl) > 1 {
		nums := len(bl)
		index := rand.Intn(nums)
		return bl[index].GetPrevBlock()
	}
	return nil
}
func (s *Synchronizer) handleControl(rs *common.RoundSummary) {
	rs.EvilTargetSeq = -2
	s.console.Controlled(rs)

}
func (s *Synchronizer) GetBlockByIndex(seq int) *block.Node {
	return s.chain.GetBlockBySeq(seq)
}
func (s *Synchronizer) PrintChain() {
	s.chain.PrintChain()
}
func (s *Synchronizer) handleSummary(wrap []*common.BlockWrap) *common.RoundSummary {
	if len(wrap) == 0 {
		rs := &common.RoundSummary{
			RoundType: common.VOID,
		}
		return rs
		for i := 0; i < s.nodeNums; i++ {
			s.summaryChan <- rs
		}

	} else {
		//generate block with sequence
		blockList := make([]*block.Node, 0, len(wrap))
		for i := 0; i < len(wrap); i++ {
			nBlock := block.NewBlock(wrap[i].IsFromEvil, wrap[i].PrevEvil, wrap[i].Prev, wrap[i].Owner, s.round, s.blockSeq, wrap[i].IsEvilPurpose, wrap[i].EvilTarget)
			s.blockSeq += 1
			blockList = append(blockList, nBlock)
			s.console.AddCount()
			nBlock.PrintBlock()
			s.chain.AddNode(nBlock)
		}
		rs := &common.RoundSummary{
			RoundType: common.PROD,
			Blocks:    blockList,
		}
		return rs
		for i := 0; i < s.nodeNums; i++ {
			s.summaryChan <- rs
		}
	}
	return nil
}
