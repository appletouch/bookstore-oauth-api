package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	// example of using asserts package
	assert.EqualValues(t, 24, accesstokenTTL, "TTL of a accesstoken should be 24 hours\"")
}

//by default Golang does not have asserts and test are done in the way below (conversion to assert not done for this purpose)
func TestGetNewAccesstoken(t *testing.T) {
	AT := GetNewAccesstoken()
	if AT.IsExpired() {
		t.Error("brand new access token should not be expired")
	}
	if AT.Access_Token != "" {
		t.Error("New access token should not have defined access token id")
	}

	if AT.UserId != 0 {
		t.Error("New access token should not have a associated user id")
	}

}

//by default Golang does not have asserts and test are done in the way below (conversion to assert not done for this purpose)
func TestAccesstoken_IsExpired(t *testing.T) {
	AT := AccessToken{}
	if !AT.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	AT.Expires = time.Now().UTC().Add(time.Hour * 3).Unix()
	if AT.IsExpired() {
		t.Error("Access token expiring three hours from now should not be expired")
	}
}
