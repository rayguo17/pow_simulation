package character

//every round done, should be able to manually start next round, after checking the output.
type Synchronizer struct {
	nodeNums           int
	hashDoneInformChan chan bool
	round              int
	receiveBlockChan   chan *Block
	broadCastBlockChan chan *BlockWrap
	summaryChan        chan *RoundSummary
	blockSeq           int
}

func NewSynchronizer(nodeNums int, hashDoneInformChan chan bool, receiveBlockChan chan *Block, broadCastBlockChan chan *BlockWrap, summaryChan chan *RoundSummary) *Synchronizer {
	return &Synchronizer{
		nodeNums:           nodeNums,
		hashDoneInformChan: hashDoneInformChan,
		receiveBlockChan:   receiveBlockChan,
		broadCastBlockChan: broadCastBlockChan,
		summaryChan:        summaryChan,
		round:              0,
		blockSeq:           0,
	}
}

func (s *Synchronizer) MainRoutine() {
	for {
		//every round
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
				//start next round

			}
		}
		//send the summary
		//calculate summary
		//should add some controllerable factor...
		s.round += 1
	}
}
