package credential

import (
	"github.com/cariapo/cservice/credential/model/request"
	"github.com/cariapo/cservice/credential/model/response"
)

type CredentialOutbound interface {
	// Oauth
	Oauth(req request.StoreToken) (*response.StoreToken, error)
	Authorize(token string) (*response.Authorize, error)

	// User
	Register(req request.Register) (*response.User, error)
	ChangePassword(req request.ChangePassword) (*response.User, error)
	ResetPassword(req request.ResetPassword) (*response.User, error)
}
