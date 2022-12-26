package character

type Block struct {
	isEvil bool   //tag only for simulation
	prev   *Block //from which
	owner  int    //id of node
	round  int    //which round
	seq    int
}

func NewBlock(isEvil bool, prev *Block, owner int, round int) *Block {
	return &Block{
		isEvil: isEvil,
		prev:   prev,
		owner:  owner,
		round:  round,
	}
}
