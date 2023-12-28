package blockchain

import (
	"LR-Chain/constcoe"
	"LR-Chain/utils"
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

/*
构建一个可以返回目标难度值的函数
这里使用的之前设定的一个常量Difficulty来构造目标难度值
但是在实际的区块链中目标难度值会根据网络情况定时进行调整，且能够保证各节点在同一时间在同一难度下进行竞争
故这里的GetTarget可以理解为预留API
*/
func (b *Block) GetTarget() []byte {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-constcoe.DIFFICULTY))
	//Lsh函数就是向左移位，移的越多目标难度值越大，哈希取值落在的空间就更多也就越容易找到符合条件的nonce
	return target.Bytes()
}

// 每次输入一个nonce对应的区块的哈希值都会改变
func (b *Block) GetBase4Nonce(nonce int64) []byte {
	data := bytes.Join([][]byte{
		utils.ToHexInt(b.Timestamp),
		b.PrevHash,
		utils.ToHexInt(int64(nonce)),
		b.Target,
		b.Data,
	},
		[]byte{},
	)
	return data
}

// 对于任意一个区块，都能去寻找一个合适的nonce
func (b *Block) FindNonce() int64 {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	var nonce int64
	nonce = 0
	intTarget.SetBytes(b.Target)

	for nonce < math.MaxInt64 {
		data := b.GetBase4Nonce(nonce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(&intTarget) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce
}

func (b *Block) ValidatePow() bool {
	var intHash big.Int
	var intTarget big.Int
	var hash [32]byte
	intTarget.SetBytes(b.Target)
	data := b.GetBase4Nonce(b.Nonce)
	hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	if intHash.Cmp(&intTarget) == -1 {
		return true
	}
	return false
}
