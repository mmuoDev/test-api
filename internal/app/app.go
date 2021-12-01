package app

import "net/http"

//App contains handlers
type App struct {
	AddUserHandler http.HandlerFunc
}

//Handler returns the main handler for this application
func (a App) Handler() http.HandlerFunc {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user", a.AddUserHandler)
	return http.HandlerFunc(router.ServeHTTP)
}

// Options is a type for application options to modify the app
type Options func(o *OptionalArgs)

// /OptionalArgs optional arguments for this application
type OptionalArgs struct {
	AddUser db.AddUserFunc
}

//New creates a new instance of the App
func New(options ...Options) App {
	o := OptionalArgs{
		AddUser:                db.AddUser(),
	}
	for _, option := range options {
		option(&o)
	}
	addUser := AddUserHandler(o.AddUser)
	return App{
		AddUserHandler: addUser,
	}
}