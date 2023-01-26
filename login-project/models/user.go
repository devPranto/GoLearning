package models

import "time"

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" bson:"_id"`
	Gender    string `json:"gender"`
	Password  []byte `json:"password"`
	Path      string `json:"path" bson:"path"`
	JWT       string
	HasBlock  bool   `json:"hasBlock" bson:"has_block"`
	LastHash  string `json:"genesisHash" bson:"genesis_hash"`
}
type Block struct {
	Sender          string
	Receiver        string
	Amount          int8
	TimeStamp       time.Time
	TransactionData string
	Difficulty      string
	Nonce           string
	Hash            string
}
type BlockData struct {
	User            string `bson:"user"`
	Id              int    `json:"block_id" bson:"block_id"`
	Hash            string `json:"hash"bson:"hash"` // sha256
	TransactionData []byte `json:"transaction_data"bson:"transactionData"`
}

// todo email can be converted to bson _id to make it unique key
// todo there can be added role int which can control user access
