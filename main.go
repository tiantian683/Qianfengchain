package main

import (
	"XianfenChain04/chain"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	blockchain := chain.CreateChianWithGensis([]byte("hello world"))
	blockchain.CreateNewBlcok([]byte("hello wordl"))
			fmt.Println("区块链中的区块个数",len(blockchain.Blocks))
			fmt.Println("区块0的哈希值",blockchain.Blocks[0])
	        fmt.Println("区块1的哈希值：",blockchain.Blocks[1])
	//fmt.Println("区块0的哈希值",block0.Hash)
	//fmt.Println("区块1的哈希值",block1.Hash)
}