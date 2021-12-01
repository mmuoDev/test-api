package internal

//User represents an internal user
type User struct {
	ID       uuid.v4 `bson:"user_id"`
	Username string  `bson:"username"`
	Password string  `bson:"password"`
	Email    string  `bson:"email"`
}
