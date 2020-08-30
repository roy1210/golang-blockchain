package main

import (
	"fmt"
	"log"
	"time"
)

// Block struct: interface of block
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// NewBlock func: return block
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Print func: custom print method to read easier
func (b *Block) Print() {
	fmt.Printf("timestamp     	 %d\n", b.timestamp)
	fmt.Printf("nonce     			 %d\n", b.nonce)
	fmt.Printf("previous_hash    %s\n", b.previousHash)
	fmt.Printf("transactions     %s\n", b.transactions)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}

// Memo init(): Define log prefix
// Memo: *: value of the memory address(&) 実体;  &: Pass the memory address.
// Memo NewBlock(): `new` will return the pointer of struct. So b can return pointer
