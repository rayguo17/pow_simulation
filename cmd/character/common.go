package character

const (
	VOID int = 1
	PROD     = 2
)

type BlockWrap struct {
	isEvil bool   //tag only for simulation
	prev   *Block //from which
	owner  int    //id of node
	round  int    //which round
}

type RoundSummary struct {
	haveBlock bool
	blocks    []*Block
	roundType int
}

func tester() {

}
