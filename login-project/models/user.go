package models

type User struct {
	//ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}
