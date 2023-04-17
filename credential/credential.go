package credential

import (
	"github.com/cuwand/cservice/credential/model/request"
	"github.com/cuwand/cservice/credential/model/response"
)

type CredentialOutbound interface {
	Oauth(req request.StoreToken) (*response.StoreToken, error)
	Authorize(token string) (*response.Authorize, error)
}
