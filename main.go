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

	firstBlock := blockchain.Blocks[0]
	firstBytes ,err :=firstBlock.Serialize()
	if err != nil {
		panic(err.Error())
	}
	//反序列化，验证逆过程
	deFirstBlock , err :=chain.DeSerialize(firstBytes)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(deFirstBlock.Data))
}