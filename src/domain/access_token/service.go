package access_token

//Service has the business logic, validation and coordinates

import (
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"strings"
)

type RepositoryInterface interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationtime(AccessToken) *errors.RestErr
}

type ServiceInterface interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationtime(AccessToken) *errors.RestErr
}

type service struct {
	repository RepositoryInterface
}

func NewService(repo RepositoryInterface) ServiceInterface {
	return &service{
		repository: repo,
	}
}

func (service *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.New(400, "Invalid access token Id")
	}

	accesstoken, err := service.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accesstoken, nil
}

func (service *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return service.repository.Create(at)
}

func (service *service) UpdateExpirationtime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return service.repository.UpdateExpirationtime(at)
}
