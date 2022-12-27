package block

import (
	"errors"
	"fmt"
)

type Tree struct {
	storage map[int]*Node //[seq]*node
	leaves  []int         //seq
}

func NewTree(initBlock *Node) *Tree {
	tree := &Tree{
		storage: make(map[int]*Node),
		leaves:  make([]int, 0, 10),
	}
	tree.AddNode(initBlock)
	return tree
}
func (t *Tree) PrintChain() {
	bl := make([][]int, 0, len(t.leaves))
	for i := 0; i < len(t.leaves); i++ {
		tchain := make([]int, 0)
		tmpBlock := t.storage[t.leaves[i]]
		for tmpBlock != nil {
			tchain = append(tchain, tmpBlock.seq)
			tmpBlock = tmpBlock.prev
		}
		bl = append(bl, tchain)
		for j := len(tchain) - 1; j >= 0; j-- {
			fmt.Printf("%d ", tchain[j])
		}
		fmt.Println("")
	}

}

func (t *Tree) AddNode(node *Node) error {
	//see if able to add
	if _, ok := t.storage[node.seq]; ok {
		return errors.New("inserting existed node")
	}
	if len(t.leaves) == 0 {
		//initial
		t.leaves = append(t.leaves, node.seq)
		t.storage[node.seq] = node
		return nil
	}
	//have value
	if t.inLeaves(node.prev.seq) {
		t.updateLeaves(node.prev.seq, node.seq)
		t.storage[node.seq] = node
		return nil
	}
	//
	if _, ok := t.storage[node.prev.seq]; ok {
		t.leaves = append(t.leaves, node.seq)
		t.storage[node.seq] = node
		return nil
	}

	return errors.New("leaves not empty and not in leaf")
}
func (t *Tree) calChainLen(seq int) int {
	//for each leaves, calculate their chain length...
	length := 0
	tmpblock := t.storage[seq]
	for tmpblock != nil {
		length += 1
		tmpblock = tmpblock.prev
	}
	return length
}

//
func (t *Tree) GetLongestChain() []*Node {
	lenMap := make(map[int]int) //chain length map
	longest := 0
	for i := 0; i < len(t.leaves); i++ {
		chainLength := t.calChainLen(t.leaves[i])
		lenMap[t.leaves[i]] = chainLength
		if chainLength > longest {
			longest = chainLength
		}
	}
	res := make([]*Node, 0)
	for i := 0; i < len(t.leaves); i++ {
		if lenMap[t.leaves[i]] == longest {
			res = append(res, t.storage[t.leaves[i]])
		}
	}
	return res
}

func (t *Tree) updateLeaves(oldSeq int, newSeq int) error {
	for i := 0; i < len(t.leaves); i++ {
		if t.leaves[i] == oldSeq {
			t.leaves[i] = newSeq
			return nil
		}
	}
	return errors.New("cannot find old seq")
}
func (t *Tree) inLeaves(seq int) bool {
	for i := 0; i < len(t.leaves); i++ {
		if seq == t.leaves[i] {
			return true
		}
	}
	return false
}
