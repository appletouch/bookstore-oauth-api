package access_token

import (
	"github.com/appletouch/bookstore-oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	//number of hours a at is valid.
	accesstokenTTL = 24
)

//ACCESS TOKEN
type AccessToken struct {
	Access_Token string `json:"access_token"`
	UserId       int64  `json:"user_id"`
	ClientId     int64  `json:"client_id"`
	Expires      int64  `json:"expires"`
}

func (AT *AccessToken) IsExpired() bool {
	return time.Unix(AT.Expires, 0).Before(time.Now().UTC())
}

func GetNewAccesstoken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(accesstokenTTL * time.Hour).Unix(),
	}
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.Access_Token = strings.TrimSpace(at.Access_Token)
	if at.Access_Token == "" {
		return errors.New(500)
	}
	if at.UserId <= 0 {
		return errors.New(400, "invalid user")
	}
	if at.ClientId <= 0 {
		return errors.New(400, "invalid client id")
	}
	if at.Expires <= 0 {
		return errors.New(400, "invalid Expired time")
	}

	return nil
}

//ACCESS TOKEN REQUEST
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	//Used for password grant type
	Email    string `json:"access_token"`
	Password string `json:"user_id"`

	//Used for client credentials
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Scope string `json:"scope"`
}

func (atr *AccessTokenRequest) Validate() *errors.RestErr {

	atr.GrantType = strings.TrimSpace(atr.GrantType)

	switch atr.GrantType {
	case "password":
		atr.Email = strings.TrimSpace(atr.Email)
		atr.Password = strings.TrimSpace(atr.Password)
		if atr.Email == "" {
			return errors.New(400)
		}
		if atr.Password == "" {
			return errors.New(400)
		}
		break

	case "ClientCredentials":
		if atr.ClientId == "" {
			return errors.New(400)
		}
		if atr.ClientSecret == "" {
			return errors.New(400)
		}
		break

	default:
		return errors.New(400, "invalid grantype")
	}

	return nil
}
