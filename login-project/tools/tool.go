package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"log"
	"login-project/models"
	"time"
)

var (
	MySigningKey = []byte("AllYourBase")
	expireTime   = time.Now().Add(2 * time.Hour)
	key          = []byte("the-key-has-to-be-32-bytes-long!")
)

func CreatJWT(user *models.User) string {

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireTime),
		Issuer:    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(MySigningKey)
	return ss
}

func CreateBlock(email string, genesis bool, prevHash string) (string, []byte) {
	// block created with Json format
	//fixme : design spec: block number in general or for a particular user ?
	blockData := &models.BlockData{}
	b := &models.Block{}
	b.New()
	blockData.User, b.Sender = email, email

	if genesis {
		blockData.Id = 0
		b.Hash = "00000"
	} else {
		serial := models.BlockNo(email)
		blockData.Id = serial
		b.Hash = prevHash
	}
	block, err := json.Marshal(b)
	if err != nil {
		fmt.Println("Error marshal : ", err.Error())
		return "nil", nil
	}

	//Encryption Mechanism
	ciphertext, err := encrypt(block)
	checker, _ := Decrypt(ciphertext)
	fmt.Println("Decrypted Cypher Text ", checker)
	if err != nil {
		log.Fatal(err)
	}
	//hashString := fmt.Sprintf("%x", ciphertext)

	//Inserting block data into database (MongoDB)
	blockData.TransactionData = ciphertext
	blockData.Hash = createHash(ciphertext)
	blockData.Insert()
	return blockData.Hash, block
}

func encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("in decrypt 1")
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("in decrypt 2")
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
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
