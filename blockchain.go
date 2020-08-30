package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Block struct: interface of block
type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []string
}

// NewBlock func: return block
func NewBlock(nonce int, previousHash [32]byte) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Print func: custom print method for Block
func (b *Block) Print() {
	fmt.Printf("timestamp     	 %d\n", b.timestamp)
	fmt.Printf("nonce     	 %d\n", b.nonce)
	fmt.Printf("previous_hash    %x\n", b.previousHash)
	fmt.Printf("transactions     %s\n", b.transactions)
}

// Hash returns [32]byte due to Sum256
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

// MarshalJSON func(): Customized overwrite method
func (b *Block) MarshalJSON() ([]byte, error) {
	// Memo: Change struct due to block struct is private struct
	return json.Marshal(struct {
		Timestamp    int64    `json: "timestamp"`
		Nonce        int      `json: "nonce"`
		PreviousHash [32]byte `json: "previous_hash"`
		Transactions []string `json: "transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Print func: custom print method for BlockChain
func (bc *Blockchain) Print() {
	// Memo: loop the `chain` array
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockChain()
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}

// Memo init(): Define log prefix
// Memo: *: value of the memory address(&) 実体;  &: Pass the memory address.
// Memo NewBlock(): `new` will return the pointer of struct. So b can return pointer
