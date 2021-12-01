package workflow

//AddUserFunc returns functionality to add a user
type AddUserFunc func(req pkg.User) (pkg.UserResponse, error)

//AddUser adds a user
func AddUser(addUser db.AddUserFunc) AddUserFunc {
	return func(req pkg.User) (pkg.UserResponse, error) {
		user := mapping.ToDBUser(req)
		//insert into db
		if err := addUser(user); err != nil {
			return pkg.UserResponse, errors.Wrapf(err, "workflow - unable to insert into DB")
		}
		return pkg.UserResponse{UserID: user.ID}, nil 
	}
}
