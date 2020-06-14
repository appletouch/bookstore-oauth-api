package rest

import (
	"encoding/json"
	"github.com/appletouch/bookstore-oauth-api/src/domain/users"
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"net/http"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:3000",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRepositoryInterface interface {
	Login(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRespository() UsersRepositoryInterface {
	return &usersRepository{}
}

func (us *usersRepository) Login(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/user/login", request)

	// no response when we have een timeou
	if response == nil || response.Response == nil {
		return nil, &errors.RestErr{http.StatusInternalServerError, "Timeout", "Somthing timed-out while calling the users Api"}
	}

	// Some error has occurred if code is above 299 and you should be able to map it as all errors should have the same interface
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, &errors.RestErr{502, "Error interface broken", "Error interface is not correct"}
		}
		// if a other error occures
		return nil, &errors.RestErr{404, "Invalid credentials", "invalid login credential"}
	}
	// no error than process response
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, &errors.RestErr{500, "Bad response", "error while unmarshal response"}
	}
	return &user, nil

}
