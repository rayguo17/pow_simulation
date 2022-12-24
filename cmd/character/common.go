package character

type BlockWrap struct {
	owner int
	block *Block
}

type RoundSummary struct {
	haveBlock bool
	block     *Block
}
