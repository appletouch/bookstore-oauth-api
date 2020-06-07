package db

import (
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
)

type DbRepositoryInterface interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (repo *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.New(500, "Database connection is down")
}

func NewRepository() DbRepositoryInterface {
	return &dbRepository{}
}
