package access_token

import "github.com/appletouch/bookstore-oauth-api/src/utils/errors"

type RepositoryInterface interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type ServiceInterface interface {
	GetById(string) (*AccessToken, *errors.RestErr)
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
	accesstoken, err := service.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accesstoken, nil
}
