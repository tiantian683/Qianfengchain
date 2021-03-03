package chain

import (
	"XianfenChain04/chain/utils"
	"XianfenChain04/consensus"
	"bytes"
	"crypto/sha256"
	"time"
)

const VERSION  = 0
/*
区块的结构体定义
 */
type Block struct {
	Height int64 //高度
	Version int64
	PrevHash [32]byte
	Hash [32]byte
	//默克尔根
	TimeStamp int64
	//Difficulty int64
	Nonce int64
	//区块体
	Data []byte
}
//计算哈希并赋值
func (block *Block)CalculateBlockHash()  {
	heightByte,_ := utils.Int2Byte(block.Height)
	versionByte,_ :=utils.Int2Byte(block.Version)
	timeByte,_ := utils.Int2Byte(block.TimeStamp)
	nonceByte,_:=utils.Int2Byte(block.Nonce)

	blockByte :=  bytes.Join([][]byte{heightByte,versionByte,block.PrevHash[:],timeByte,nonceByte,block.Data},[]byte{})
	block.Hash = sha256.Sum256(blockByte)
}

/*
生成创世区块的函数
 */
func CreateGensis(data []byte)Block  {
	genesis := Block{
		Height:    0,
		Version:   VERSION,
		PrevHash:  [32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		TimeStamp: time.Now().Unix(),
		Data:      data,
	}
	//todo 设置哈希，寻找并设置NONCE
	//计算并设置哈希值
	genesis.CalculateBlockHash()
	//
	proof := consensus.NewPoW(genesis)
	genesis.Nonce =proof.FindNonce()
	return genesis
}

/**
*生成新区块的功能函数
 */
func NewBlock(height int64,Prev [32]byte,data []byte) Block {
	newBlock := Block{
		Height:    height+1,
		Version:   VERSION,
		PrevHash:  Prev,
		TimeStamp: time.Now().Unix(),
		Data:      data,
	}
	 //todo  设置哈希、寻找并设置nonce
	 //设置区块hash
	 newBlock.CalculateBlockHash()
	proof := consensus.NewPoW(newBlock)
	newBlock.Nonce =proof.FindNonce()
	return newBlock
}
