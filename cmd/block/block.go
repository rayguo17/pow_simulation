package block

import "fmt"

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
func (n *Node) PrintBlock() {
	fmt.Printf("round: %d, seq: %d, prevSeq: %d, owner: %d, isEvil: %v, prevEvil: %v\n", n.round, n.seq, n.prev.seq, n.owner, n.isEvil, n.prevEvil)
}

func (n *Node) InheritEvil() bool {
	return n.isEvil || n.prevEvil
}

func (n *Node) GetPrevBlock() *Node {
	return n.prev
}
