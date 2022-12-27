package block

import "fmt"

type Node struct {
	isEvil        bool //tag only for simulation
	prevEvil      bool
	prev          *Node //from which
	owner         int   //id of node
	round         int   //which round
	seq           int
	isEvilPurpose bool
	evilTarget    int
}

func NewBlock(e bool, pe bool, prev *Node, owner int, round int, seq int, isEvilPurpose bool, evilTarget int) *Node {
	return &Node{
		isEvil:        e,
		prevEvil:      pe,
		prev:          prev,
		owner:         owner,
		round:         round,
		seq:           seq,
		isEvilPurpose: isEvilPurpose,
		evilTarget:    evilTarget,
	}
}
func (n *Node) PrintBlock() {
	fmt.Printf("round: %d, seq: %d, prevSeq: %d, owner: %d, isEvil: %v, prevEvil: %v, evilTarget: %d\n", n.round, n.seq, n.prev.seq, n.owner, n.isEvil, n.prevEvil, n.evilTarget)
}

func (n *Node) InheritEvil() bool {
	return n.isEvilPurpose || n.prevEvil
}
func (n *Node) GetSeq() int {
	return n.seq
}
func (n *Node) GetPrevEvil() bool {
	return n.prevEvil
}
func (n *Node) GetEvilTarget() int {
	return n.evilTarget
}
func (n *Node) GetPrevBlock() *Node {
	return n.prev
}
