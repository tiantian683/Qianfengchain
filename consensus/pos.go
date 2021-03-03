package consensus

import (
	"XianfenChain04/chain"
	"fmt"
)

type PoS struct {
	Block chain.Block
}

func (pos PoS)FindNonce() int64{
	fmt.Println("这里是共识机制pow")
	return 0
}
