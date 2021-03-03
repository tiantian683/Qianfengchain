package main

import (
	"XianfenChain04/chain"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	block0 := chain.CreateGensis([]byte("hello world"))
	block1 := chain.NewBlock(block0.Height,block0.Hash,[]byte("hello"))
			fmt.Println(block0)
	        fmt.Println(block1)
	fmt.Println("区块0的哈希值",block0.Hash)
	fmt.Println("区块1的哈希值",block1.Hash)
}