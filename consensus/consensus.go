package consensus

import "XianfenChain04/chain"

type Consensus interface {
	FindNonce()int64
}

func NewPoW(block chain.Block) Consensus {
	return PoW{}
}

func NewPoS(block chain.Block) Consensus  {
	return PoS{}
}