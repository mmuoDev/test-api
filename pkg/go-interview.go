package pkg

//User describes request for adding a user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//UserResponse represents response after adding a user
type UserResponse struct {
	UserID string `json:"user_id"`
}
