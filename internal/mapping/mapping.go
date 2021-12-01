package mapping


func ToDBUser(req pkg.User) (User, error) {
	//generate uuid
	id := 4
	return User{
		ID: id,
		Username: req.Username,
	}
}
