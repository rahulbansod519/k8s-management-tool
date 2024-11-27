package models

type User struct {
	ID       string `bson:"_id,omitempty"` // MongoDB uses _id field
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Role     string `bson:"role"` // e.g., admin, client
}
