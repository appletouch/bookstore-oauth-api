package services

//Service has the business logic, validation and coordinates

import (
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/appletouch/bookstore-oauth-api/src/repository/db"
	"github.com/appletouch/bookstore-oauth-api/src/repository/rest"
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"strings"
)

//type RepositoryInterface interface {
//	GetById(string) (*AccessToken, *errors.RestErr)
//	Create(AccessTokenRequest) *errors.RestErr
//	UpdateExpirationtime(AccessToken) *errors.RestErr
//}

type ServiceInterface interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationtime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.UsersRepositoryInterface
	dbRepo       db.DbRepositoryInterface
}

func NewService(userRepo rest.UsersRepositoryInterface, dpRepo db.DbRepositoryInterface) ServiceInterface {
	return &service{
		restUserRepo: userRepo,
		dbRepo:       dpRepo,
	}
}

func (service *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.New(400, "Invalid access token Id")
	}

	accesstoken, err := service.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accesstoken, nil
}

func (service *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//Authenicate User against the user API with email / password
	user, err := service.restUserRepo.Login(request.UserName, request.Password)
	if err != nil {
		return nil, err
	}
	//Genreate a new Access token
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	//save the new accesstoken in the cassandra db
	if err := service.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil

}

func (service *service) UpdateExpirationtime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return service.dbRepo.UpdateExpirationtime(at)
}
