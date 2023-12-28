package main

import (
	"LR-Chain/blockchain"
	"fmt"
	"time"
)

func main() {
	blockChain := blockchain.CreateBlockChain()
	time.Sleep(time.Second)
	blockChain.AddBlock("After genesis, I have something to say.")
	time.Sleep(time.Second)
	blockChain.AddBlock("I'm LVRUI!!!")
	time.Sleep(time.Second)
	blockChain.AddBlock("I love BlockChain")
	time.Sleep(time.Second)

	for _, block := range blockChain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println("Proof of Work validation:", block.ValidatePow())
	}
}
