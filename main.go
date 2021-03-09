package main

import (
	"XianfenChain04/chain"
	"fmt"
	"github.com/bolt-master"
)

const BLOCKS   = "xianfengchain04.db"
func main() {
	//打开数据库文件
	db,err := bolt.Open(BLOCKS,0600,nil)
	if err!=nil {
		panic(err.Error())
	}
	defer db.Close()

	blockChian := chain.CreateChian(db)
	//创世区块
	err = blockChian.CreateGensis([]byte("Hello World!!"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//新增一个区块
	err = blockChian.CreateNewBlcok([]byte("Hello"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//测试
	lastBlock,err := blockChian.GetLastBlock()
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("最新区块是",lastBlock)

	blocks,err := blockChian.GetAllBlocks()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	for  index,block := range blocks  {
		fmt.Printf("第%d个区块",index)
		fmt.Println(block)
	}
	//blockchain := chain.CreateChianWithGensis([]b	//fmt.Println("hello world")yte("hello world"))
	//
	//blockchain.CreateNewBlcok([]byte("hello wordl"))
	//		fmt.Println("区块链中的区块个数",len(blockchain.Blocks))
	//		fmt.Println("区块0的哈希值",blockchain.Blocks[0])
	//        fmt.Println("区块1的哈希值：",blockchain.Blocks[1])
	//
	//firstBlock := blockchain.Blocks[0]
	//firstBytes ,err :=firstBlock.Serialize()
	//if err != nil {
	//	panic(err.Error())
	//}
	////反序列化，验证逆过程
	//deFirstBlock , err := chain.DeSerialize(firstBytes)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(string(deFirstBlock.Data))


}