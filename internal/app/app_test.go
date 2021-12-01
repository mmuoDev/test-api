package app

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestAddUserWorksAsExpected(t *testing.T) {
	insertIntoDBInvoked := false

	mockAuthorDBInsert := func(o *app.OptionalArgs) {
		o.AddUser = func(a internal.User) error {
			insertIntoDBInvoked = true
			return nil
		}
	}

	//optional args
	opts := []app.Options{
		mockAuthorDBInsert,
	}

	ap := app.New(opts...)
	serverURL, cleanUpServer := app.NewTestServer(ap.Handler())
	defer cleanUpServer()

	reqPayload, _ := os.Open(filepath.Join("testdata", "request_user_data.json"))
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/users", serverURL), reqPayload)

	client := &http.Client{}
	res, _ := client.Do(req)

	t.Run("Http Status Code is 201", func(t *testing.T) {
		assert.Equal(t, res.StatusCode, http.StatusCreated)
	})

	t.Run("Insert to DB invoked", func(t *testing.T) {
		assert.True(t, insertIntoDBInvoked)
	})
}
