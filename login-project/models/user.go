package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" bson:"_id"`
	Gender    string `json:"gender"`
	Password  []byte `json:"password"`
	Path 	  string `json:"path" bson:"path"`
	JWT     string 
}

// todo email can be converted to bson _id to make it unique key
// todo there can be added role int which can control user access
