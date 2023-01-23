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
	h := sha256.New()

	_, er := h.Write(block)

	if er != nil {
		fmt.Println("error hash : ", err.Error())
		return
	}
	hash := h.Sum(nil)
	//fmt.Println(string(hash))
	fmt.Printf("hash : %x \n", hash)
	fmt.Printf("%+v \n", b)

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
