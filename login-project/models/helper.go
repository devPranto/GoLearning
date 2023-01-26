package models

import (
	"context"
	"fmt"
	"github.com/larrybattle/nonce-golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	connectionString = "mongodb+srv://pranto:password18@cluster0.ittknvy.mongodb.net/?retryWrites=true&w=majority"
	dbName           = "loginApp"
	collectionName   = "user"
	collectionName2  = "UserBlock"
)

var collection *mongo.Collection
var collection2 *mongo.Collection

func init() {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb connection successful")
	collection = client.Database(dbName).Collection(collectionName)
	collection2 = client.Database(dbName).Collection(collectionName2)
	fmt.Printf("Collection1 : %v \n Collection2 : %v \n", collection, collection2)
}
func (user *User) Insert() {
	result, err := collection.InsertOne(context.Background(), &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func InsertToken(token *User) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", token.Email}}
	update := bson.D{{"$set", bson.D{{"jwt", token.JWT}}}}
	collection.UpdateOne(context.TODO(), filter, update, opts)

}
func Update(user *User) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", user.Email}}
	update := bson.D{{"$set", bson.D{{"firstname", user.FirstName}, {"lastname", user.LastName}, {"password", user.Password}}}}
	result, _ := collection.UpdateOne(context.TODO(), filter, update, opts)
	fmt.Println(result)
}
func UpdateBlockStatus(user *User) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", user.Email}}
	update := bson.D{{"$set", bson.D{{"has_block", user.HasBlock}, {"genesis_hash", user.LastHash}}}}
	result, _ := collection.UpdateOne(context.TODO(), filter, update, opts)
	fmt.Println(result)
}
func Find(mail string) *User {
	var result User
	filter := bson.D{{"_id", mail}}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("User doesn't Exist")
	}
	return &result
}

// block functions
// todo use collection2 for block related database works
func (b *Block) New() {
	if b.Difficulty == "" {
		b.Difficulty = "Hard"
	}
	if b.Nonce == "" {
		b.Nonce = nonce.NewToken()
	}
	b.TimeStamp = time.Now()
	b.Receiver = "Admin"
	b.Amount = 71
}

func (block *BlockData) Insert() {
	result, err := collection2.InsertOne(context.Background(), &block)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func BlockNo(email string) int {
	fmt.Println(email)
	filter := bson.D{{"user", email}}
	result, err := collection2.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("couldn't find user block in db")
		return 0
	}
	var results []BlockData
	result.All(context.TODO(), &results)
	return len(results)
}

func FindBlock(user string) []BlockData {
	var results []BlockData
	filter := bson.D{{"user", user}}
	result, err := collection2.Find(context.Background(), filter)
	if err != nil {
		log.Fatal("Couldn't Find Block")
		return nil
	}
	result.All(context.TODO(), &results)
	return results
}
func FindBlockData(hash string) *BlockData {
	var result BlockData
	filter := bson.D{{"hash", hash}}
	err := collection2.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Block doesn't Exist")
	}
	return &result
}
