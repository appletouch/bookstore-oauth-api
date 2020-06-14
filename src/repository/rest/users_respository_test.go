package rest

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// entry point of every test suite
func TestMain(m *testing.M) {
	fmt.Println("About to start Testcases")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromAPI(t *testing.T) {

	//remove old mocks
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:3000/users/user/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"kooltjes@yahoo.com", "password":"welkom01"}`,
		RespHTTPCode: -1,
	})

	repository := usersRepository{}
	user, err := repository.Login("kooltjes@yahoo.com", "welkom01")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Somthing timed-out while calling the users Api", err.Detail)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	//remove old mocks
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:3000/users/user/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"kool@yahoo.com", "password":"welkom01"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":401, "error": "not found"}`,
	})

	repository := usersRepository{}
	user, err := repository.Login("kool@yahoo.com", "welkom01")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credential", err.Detail)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	//remove old mocks
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:3000/users/user/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"kooltjes@yahoo.com", "password":"welkom01"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"1"","first_name":"Pet","last_name":"Kool","email":"kool@yahoo.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.Login("kool@yahoo.com", "welkom01")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error while unmarshal response", err.Detail)
}

func TestLoginUserNoError(t *testing.T) {
	//remove old mocks
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:3000/users/user/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"kooltjes@yahoo.com", "password":"welkom01"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":1,"first_name":"Pet","last_name":"Kool","email":"kool@yahoo.com"}`,
	})

	repository := usersRepository{}
	user, err := repository.Login("kool@yahoo.com", "welkom01")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "Pet", user.FirstName)
	assert.EqualValues(t, "Kool", user.LastName)
	assert.EqualValues(t, "kool@yahoo.com", user.Email)
	assert.EqualValues(t, 1, user.Id)
}
