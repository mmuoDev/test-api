package app

import "net/http"

func AddUserHandler(addUser db.AddUserFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//parse the request
		var req pkg.User
		httputils.JSONToDTO(&req, w, r)

		add := workflow.AddUser(addUser)
		res, err := add(req)
		if err != nil {
			httputils.ServeError(err, w)
			return
		}
		hhttputils.ServeJSON(res, w)

	}
}
