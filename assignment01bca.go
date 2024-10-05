package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	Timestamp    time.Time
}

type Blockchain struct {
	Blocks []Block
}

func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	newBlock := &Block{transaction, nonce, previousHash, "", time.Now()}

	blockData := fmt.Sprintf("%d%s%s%s", newBlock.Nonce, newBlock.Transaction, newBlock.PreviousHash, newBlock.Timestamp.String())
	newBlock.Hash = CalculateHash(blockData)

	return newBlock
}

func ChangeBlock(blockchain *Blockchain, index int, newTransaction string) error {
	if index < 0 || index >= len(blockchain.Blocks) {
		return fmt.Errorf("index out of range")
	}

	blockchain.Blocks[index].Transaction = newTransaction

	// Recalculate the block's hash after changing the transaction
	block := &blockchain.Blocks[index]
	blockData := fmt.Sprintf("%d%s%s%s", block.Nonce, block.Transaction, block.PreviousHash, block.Timestamp.String())
	block.Hash = CalculateHash(blockData)

	return nil
}

func VerifyChain(blockchain *Blockchain) bool {
	if len(blockchain.Blocks) == 0 {
		return true // An empty blockchain is valid
	}

	for i := 1; i < len(blockchain.Blocks); i++ {
		currentBlock := &blockchain.Blocks[i]
		prevBlock := &blockchain.Blocks[i-1]

		blockData := fmt.Sprintf("%d%s%s%s", currentBlock.Nonce, currentBlock.Transaction, currentBlock.PreviousHash, currentBlock.Timestamp.String())
		calculatedHash := CalculateHash(blockData)

		if currentBlock.Hash != calculatedHash {
			fmt.Printf("Block %d has been tampered with!\n", i)
			return false
		}

		if currentBlock.PreviousHash != prevBlock.Hash {
			fmt.Printf("Block %d's previous hash is incorrect!\n", i)
			return false
		}
	}

	fmt.Println("Blockchain is valid.")
	return true
}

func ListBlocks(blockchain *Blockchain) {
	fmt.Println("List of Blocks")

	for i, blk := range blockchain.Blocks {
		fmt.Printf("Block %d:\n", i+1)
		fmt.Println("Transaction: ", blk.Transaction)
		fmt.Println("Nonce: ", blk.Nonce)
		fmt.Println("Previous Hash: ", blk.PreviousHash)
		fmt.Println("Hash: ", blk.Hash)
		fmt.Println("Timestamp: ", blk.Timestamp)
	}
}
