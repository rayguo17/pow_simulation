package controller

import (
	"github.com/rayguo17/pow/cmd/common"
)

type Console struct {
	count      int
	tmpCount   int
	informChan chan bool
	evilTarget int
}

func NewConsole() *Console {
	return &Console{

		informChan: make(chan bool),
		count:      1,
		tmpCount:   0,
		evilTarget: -2,
	}
}

func (c *Console) AddCount() {
	c.tmpCount += 1
}
func (c *Console) Controlled(rs *common.RoundSummary) {
	//what's next model
	rs.EvilTargetSeq = c.evilTarget
	if c.tmpCount < c.count {
		return
	} else {
		<-c.informChan
	}
}

func (c *Console) Config(count int, evilTarget int) {
	c.count = count
	c.tmpCount = 0
	if evilTarget > -3 {
		c.evilTarget = evilTarget
	}
	//-2 clear -3 maintain, >-2 clear
	c.informChan <- true
}
