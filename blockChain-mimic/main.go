package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/larrybattle/nonce-golang"
)

type Block struct {
	Sender          string
	Receiver        string
	Amount          int8
	TimeStamp       time.Time
	TransactionData string
	Difficulty      string
	Nonce           string
	Hash            []byte
}

func main() {
	b := &Block{}
	b = b.New()
	block, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error marshal : ", err.Error())
		return
	}
	value := createHash(block)
	fmt.Println(value)
}

func (b *Block) New() *Block {
	if b.Difficulty == "" {
		b.Difficulty = "Hard"
	}
	if b.Nonce == "" {
		b.Nonce = nonce.NewToken()
	}
	b.TimeStamp = time.Now()
	return b

}
func createHash(data []byte) string {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		fmt.Println(err.Error())
	}
	hash := h.Sum(nil)
	return fmt.Sprintf("%x", hash)
}
