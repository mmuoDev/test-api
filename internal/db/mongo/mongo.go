package db

type AddUserFunc func() (User, error)

func AddUser() AddUserFunc {
	return func()  error {
		return nil 
	}
}
