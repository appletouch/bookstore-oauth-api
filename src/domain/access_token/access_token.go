package access_token

import "time"

const (
	//number of hours a at is valid.
	accesstokenTTL = 24
)

type AccessToken struct {
	Access_Token string `json:"access_token"`
	UserId       int64  `json:"user_id"`
	ClientId     int64  `json:"client_id"`
	Expires      int64  `json:"expires"`
}

func GetNewAccesstoken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(accesstokenTTL * time.Hour).Unix(),
	}
}

func (AT *AccessToken) IsExpired() bool {
	return time.Unix(AT.Expires, 0).Before(time.Now().UTC())
}
