package credential

import (
	"github.com/cariapo/cservice/credential/model/request"
	"github.com/cariapo/cservice/credential/model/response"
)

type CredentialOutbound interface {
	Oauth(req request.StoreToken) (*response.StoreToken, error)
	Authorize(token string) (*response.Authorize, error)
}
