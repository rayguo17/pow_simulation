package character

//every round done, should be able to manually start next round, after checking the output.
type Synchronizer struct {
	clockChan          chan bool //clock for synchronization
	nodeNums           int
	hashDoneInformChan chan bool
	round              int
	receiveBlockChan   chan *BlockWrap
	broadCastBlockChan chan *BlockWrap
}

func (s *Synchronizer) MainRoutine() {
	for {
		//every round
		receiverCnt := 0
		receivedBlock := make([]*BlockWrap, 0)
		for {
			select {
			case block := <-s.receiveBlockChan:
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

		s.round += 1
	}
}
