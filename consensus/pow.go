package consensus

import (
	"XianfenChain04/chain"
	"XianfenChain04/chain/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const DIFFICULTY  = 10 //难度值系数

type PoW struct {
	Block chain.Block
	Target *big.Int
}


func (pow PoW)FindNonce() int64{
	fmt.Println("这里是共识机制pow")
	//1、给定一个nonce值，计算区块的hash
	var nonce int64
	nonce = 0
	//for无限循环
	for {
		hash := Calculate(pow.Block,nonce)
		//2、拿到系统的目标值  难度
		target := pow.Target
		//3、比较大小
		result := bytes.Compare(hash[:], target.Bytes())
		//4、判断结果
		if result == 1 {
			return nonce
		}
		nonce++
	}
	return 0
}
//根据区块已有的信息和当前nonce的赋值，计算区块hash
func Calculate(block chain.Block,nonce int64)[32]byte  {
	heightByte, _ := utils.Int2Byte(block.Height)
	versionByte, _ := utils.Int2Byte(block.Version)
	timeByte, _ := utils.Int2Byte(block.TimeStamp)
	nonceByte, _ := utils.Int2Byte(nonce)
	blockByte := bytes.Join([][]byte{heightByte, versionByte, block.PrevHash[:], timeByte, nonceByte, block.Data}, []byte{})
	//计算区块的hash
	hash := sha256.Sum256(blockByte)
	return hash
}