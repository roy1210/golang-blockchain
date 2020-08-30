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
	transactions []*Transaction
}

// NewBlock func: return block
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

// Print func: custom print method for Block
func (b *Block) Print() {
	fmt.Printf("timestamp     	 %d\n", b.timestamp)
	fmt.Printf("nonce     	 %d\n", b.nonce)
	fmt.Printf("previous_hash    %x\n", b.previousHash)
	// fmt.Printf("transactions     %s\n", b.transactions)
	for _, t := range b.transactions {
		t.Print()
	}
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
		Timestamp    int64          `json: "timestamp"`
		Nonce        int            `json: "nonce"`
		PreviousHash [32]byte       `json: "previous_hash"`
		Transactions []*Transaction `json: "transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool []*Transaction
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
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

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_address  %s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_address  %s\n", t.recipientBlockchainAddress)
	fmt.Printf("value  %.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockChain()

	blockchain.AddTransaction("A", "B", 1.0)

	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)

	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}

// Memo init(): Define log prefix
// Memo: *: value of the memory address(&) 実体;  &: Pass the memory address.
// Memo NewBlock(): `new` will return the pointer of struct. So b can return pointer
