package models

import (
	"context"
	"fmt"
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
)

var collection *mongo.Collection

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
func Update() {

}
func Find(mail string) *User {
	var result User
	filter := bson.D{{"_id", mail}}
	collection.FindOne(context.Background(), filter).Decode(&result)
	return &result
}

//fixme help me with some of the things
