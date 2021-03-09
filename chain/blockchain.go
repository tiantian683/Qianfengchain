package chain

import (
	"errors"
	"github.com/bolt-master"
)

const BLOCKS  = "blocks"
const LASTHASH  = "lasthash"
//定义区块链结构体，该结构用于管理区块
type BlockChain struct {
	//Blocks []Block
	DB *bolt.DB
}

func CreateChian(db *bolt.DB) BlockChain  {
	return BlockChain{db}
}

//创建一个区块链对象，包含一个创世区块
func (chain *BlockChain) CreateGensis(data []byte) error {
	var err error
	//gensis持久化到db中去
	engine := chain.DB
	engine.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket==nil {
			bucket,err = tx.CreateBucket([]byte(BLOCKS))
			if err != nil {
				return err
				//panic("保存操作区块存储文件失败")
			}
		}
		//先查看
		lastHash := bucket.Get([]byte(LASTHASH))
		if len(lastHash)==0 {
			gensis := CreateGensis(data)
			genSerBytes,_ :=gensis.Serialize()
			//bucket已经存在
			//key -> value
			//blockHash->区块序列化以后的数据
			bucket.Put(gensis.Hash[:],genSerBytes)
			//使用一个标记，用来记录最新区块的哈希，以标明当前文件中存储到了最新的哪个区块
			bucket.Put([]byte(LASTHASH),gensis.Hash[:])
		}else {
			//lasthash 有值，长度为0，啥都不用干
		}
		return nil
	})
	return err
}
//生成一个新区块，并将新区块添加到链上
func (chain *BlockChain)CreateNewBlcok(data []byte)  error{
	//目的：生成一个新区块，并存到blot.DB文件中
	//手段(步骤):
	//1、从文件中查到当前存储的最新区块数据
	db := chain.DB
	var err error
	var lastBlock  Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil{
			err = errors.New("区块数据库操作失败，请重试！")
			return err
		}
		lastHash :=bucket.Get([]byte(LASTHASH))
		lastBlockBytes := bucket.Get(lastHash)
		//2、反序列化得到
		lastBlock,err = DeSerialize(lastBlockBytes)
		if err != nil {
			return err
		}
		return nil
	})
	//3、区块根据获取的最新区块生成一个新区块
	newBlock := NewBlock(lastBlock.Height,lastBlock.Hash,data)
	//4、将最新区块序列化，得到序列化数据
	newBlockSerBytes,err := newBlock.Serialize()
	if err != nil {
		return err
	}
	//5、将序列化数据存储到文件、同时还有更新最新区块的标记lastHahs，更新为最新区块的哈希
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil {
			 err = errors.New("区块数据库操作失败")
		}
		//将新生成的区块保存到文件中去
		bucket.Put(newBlock.Hash[:],newBlockSerBytes)
		//更新标记最新区块的标记lasthash,更新为最新区块的hash
		bucket.Put([]byte(LASTHASH),newBlock.Hash[:])
		return nil
	})
		return err
}

	//	获取最新的区块数据
func (chain *BlockChain)GetLastBlock()(Block,error) {
	db:=chain.DB
	var err error
	var lastBlock Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket==nil {
			return errors.New("区块数据区操作失败，重试")
		}
		lastHash := bucket.Get([]byte(LASTHASH))
		lastBlockBytes := bucket.Get(lastHash)
		lastBlock,err = DeSerialize(lastBlockBytes)
		if err!= nil {
			return err
		}
		return nil
	})
	return lastBlock,err
}
	//获取所有的区块数据
func (chain *BlockChain)GetAllBlocks() ([]Block,error) {
	//目的：获取所有的区块
	//手段（步骤）:
		//1、找到最后一个区块
		db := chain.DB
		var err error
		blocks := make([]Block,0)
		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(BLOCKS))
			if bucket == nil {
				err = errors.New("区块数据库操作失败" )
				return err
			}
			var currentHash []byte
			currentHash = bucket.Get([]byte(LASTHASH))
			//2、根据最后一个区块依次往前找
			for  {
				currentBlockBytes := bucket.Get(currentHash)
				currentBlock,err := DeSerialize(currentBlockBytes)
				if err!= nil {
					break
				}
				//3、每次找到的区块放入到一个[]Block容器中
				blocks = append(blocks,currentBlock)
				//4、找到最开始的创世区块是，结束
				if currentBlock.Height==0 {
					break
				}
				currentHash = currentBlock.PrevHash[:]
			}
			return nil
		})
	return blocks,err
}