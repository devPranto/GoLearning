package models

type User struct {
	//ID        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Password  []byte `json:"password"`
}

// todo email can be converted to bson _id to make it unique key
// todo there can be added role int which can control user access
