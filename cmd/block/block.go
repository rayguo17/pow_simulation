package block

type Node struct {
	isEvil   bool //tag only for simulation
	prevEvil bool
	prev     *Node //from which
	owner    int   //id of node
	round    int   //which round
	seq      int
}

func NewBlock(e bool, pe bool, prev *Node, owner int, round int, seq int) *Node {
	return &Node{
		isEvil:   e,
		prevEvil: pe,
		prev:     prev,
		owner:    owner,
		round:    round,
		seq:      seq,
	}
}

func (n *Node) InheritEvil() bool {
	return n.isEvil || n.prevEvil
}
