package chain

//定义区块链结构体，该结构用于管理区块
type BlockChain struct {
	Blocks []Block
}
//创建一个区块链对象，包含一个创世区块
func CreateChianWithGensis(data []byte)BlockChain  {
	gensis := CreateGensis(data)
	blocks := make([]Block,0)
	blocks  = append(blocks,gensis)
	return BlockChain{blocks}
}
//生成一个新区块，并将新区块添加到链上
func (chain *BlockChain)CreateNewBlcok(data []byte)  {
	blocks := chain.Blocks//获取到当前所有的区块
	lastBlock := blocks[len(blocks)-1]//最后最新的区块
	newBlock := NewBlock(lastBlock.Height,lastBlock.Hash,lastBlock.Data)
	chain.Blocks = append(chain.Blocks,newBlock)
}
