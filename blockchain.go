package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Version      string        `json:"version"`
	PreviousHash string        `json:"previousHash"`
	MerkleRoot   string        `json:"merkleRoot"`
	Timestamp    int64         `json:"timestamp"`
	Nonce        int           `json:"nonce"`
	BlockNumber  int           `json:"blockNumber"`
	Transactions []Transaction `json:"transactions"`
	Hash         string        `json:"hash"`
}

type Transaction struct {
	TransactionID   string       `json:"transactionId"`
	ProductID       string       `json:"productId"`
	ProductName     string       `json:"productName"`
	ProductType     string       `json:"productType"`
	Quantity        int          `json:"quantity"`
	UnitOfMeasure   string       `json:"unitOfMeasure"`
	TransactionType string       `json:"transactionType"`
	FromParty       Party        `json:"fromParty"`
	ToParty         Party        `json:"toParty"`
	Timestamp       int64        `json:"timestamp"`
	LocationData    Location     `json:"locationData"`
	Conditions      Conditions   `json:"conditions"`
	Notes           string       `json:"notes"`
	Attachments     []Attachment `json:"attachments"`
	Signature       string       `json:"signature"`
}

type Party struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Conditions struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

type Attachment struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Hash string `json:"hash"`
}

func (b *Block) CalculateHash() string {
	record := b.Version + b.PreviousHash + b.MerkleRoot + string(b.Timestamp) + string(b.Nonce) + string(b.BlockNumber)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func NewBlock(transactions []Transaction, previousHash string, blockNumber int) *Block {
	block := &Block{
		Version:      "1.0",
		PreviousHash: previousHash,
		MerkleRoot:   "", // Implement Merkle tree calculation
		Timestamp:    time.Now().Unix(),
		Nonce:        0,
		BlockNumber:  blockNumber,
		Transactions: transactions,
		Hash:         "", // Initialize the hash field
	}
	block.Hash = block.CalculateHash() // Calculate and assign the block hash
	return block
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock([]Transaction{}, "", 0)
	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(transactions []Transaction) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, previousBlock.Hash, len(bc.Blocks))
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
