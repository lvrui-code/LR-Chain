package blockchain

import (
	"LR-Chain/utils"
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp int64  //时间戳
	Hash      []byte // 本身的哈希值
	PrevHash  []byte // 指向上一区块的哈希值， 前三个属性构成区块头部信息
	Target    []byte // 目标难度值
	Nonce     int64  // 节点寻找到的作为卷王的证据
	Data      []byte // 区块中的数
}

// 构造哈希函数11111
// 对于一个区块而言，可以通过哈希算法概括其所包含的所有信息，哈希值就相当于区块的ID值，同时也可以用来检查区块所包含信息的完整性。
func (b *Block) SetHash() {
	information := bytes.Join([][]byte{utils.ToHexInt(b.Timestamp), b.PrevHash, b.Target, utils.ToHexInt(b.Nonce), b.Data}, []byte{})
	// bytes.Join可以将多个字节串连接，第二个参数是将字节串连接时的分隔符，这里设置为[]byte{}即为空

	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

// 创建区块
func CreateBlock(prevHash, data []byte) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevHash, []byte{}, 0, data}
	block.Target = block.GetTarget()
	block.Nonce = block.FindNonce()
	block.SetHash()
	return &block
}

// 创世区块
func GenesisBlock() *Block {
	gensisWords := "Hello, LR-Chain!!!"
	return CreateBlock([]byte{}, []byte(gensisWords))
}
