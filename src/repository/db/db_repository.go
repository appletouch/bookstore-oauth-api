package db

import (
	"github.com/appletouch/bookstore-oauth-api/src/clients/cassandra"
	"github.com/appletouch/bookstore-oauth-api/src/domain/access_token"
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken                  = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?; "
	queryCreateAccessToken               = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpirationtimeAccessToken = ""
)

type DbRepositoryInterface interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationtime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (repo *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {

	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.Access_Token,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.New(404)
		}

		return nil, errors.New(500, err.Error())

	}
	return &result, nil
}

func (repo *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.Access_Token,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		return errors.New(500, err.Error())
	}
	return nil
}

func (repo *dbRepository) UpdateExpirationtime(at access_token.AccessToken) *errors.RestErr {

	if err := cassandra.GetSession().Query(queryUpdateExpirationtimeAccessToken,
		at.Expires,
		at.Access_Token,
	).Exec(); err != nil {
		return errors.New(500, err.Error())
	}
	return nil

}

func NewRepository() DbRepositoryInterface {
	return &dbRepository{}
}
