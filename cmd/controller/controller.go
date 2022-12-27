package controller

type Console struct {
	count      int
	tmpCount   int
	informChan chan bool
}

func NewConsole() *Console {
	return &Console{

		informChan: make(chan bool),
		count:      1,
		tmpCount:   0,
	}
}

func (c *Console) AddCount() {
	c.tmpCount += 1
}
func (c *Console) Controlled() {
	//what's next model
	if c.tmpCount < c.count {
		return
	} else {
		<-c.informChan
	}
}

func (c *Console) Config(count int) {
	c.count = count
	c.tmpCount = 0
	c.informChan <- true
}
