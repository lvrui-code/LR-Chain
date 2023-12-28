package blockchain

// 区块链结构体
type BlockChain struct {
	Blocks []*Block // 区块链就是区块构成的集合
}

// 添加区块， 区块链可以根据其它信息创建区块进行储存
func (bc *BlockChain) AddBlock(data string) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 构建一个区块链初始化函数，使其返回一个包含创始区块的区块链
func CreateBlockChain() *BlockChain {
	blockChain := BlockChain{Blocks: nil}
	blockChain.Blocks = append(blockChain.Blocks, GenesisBlock())
	return &blockChain
}
