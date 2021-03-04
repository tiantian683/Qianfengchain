package consensus

import (
	"XianfenChain04/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)
//目的：拿到区块的属性数据(属性值 )
    //1、通过结构体引用，引用block
const DIFFICULTY  = 10 //难度值系数

type PoW struct {
	Block BlockInterface
	Target *big.Int
}


func (pow PoW)FindNonce() int64{
	fmt.Println("这里是共识机制pow")
	//1、给定一个nonce值，计算区块的hash
	var nonce int64
	nonce = 0
	//for无限循环
	hashBig := new(big.Int)
	for {
		hash := Calculate(pow.Block,nonce)
		//2、拿到系统的目标值  难度
		target := pow.Target
		//3、比较大小
		hashBig = hashBig.SetBytes(hash[:])
		result := hashBig.Cmp(target)
		//4、判断结果
		if result == -1 {
			return nonce
		}
		nonce++
	}
	return 0
}
//根据区块已有的信息和当前nonce的赋值，计算区块hash
func Calculate(block BlockInterface,nonce int64)[32]byte  {
	heightByte, _ := utils.Int2Byte(block.GetHeight())
	versionByte, _ := utils.Int2Byte(block.GetTimeStamp())
	timeByte, _ := utils.Int2Byte(block.GetTimeStamp())
	nonceByte, _ := utils.Int2Byte(nonce)
	perv := block.GetPrevHash()
	blockByte := bytes.Join([][]byte{heightByte,
		versionByte, perv[:],
		timeByte, nonceByte,
		block.GetData()},
		[]byte{})
	//计算区块的hash
	hash := sha256.Sum256(blockByte)
	return hash
}