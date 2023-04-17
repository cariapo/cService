package response

import (
	"fmt"
	"github.com/cuwand/pondasi/enum/tokenTypeEnums"
)

type StoreToken struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		AccessToken  string                   `json:"access_token"`
		RefreshToken string                   `json:"refresh_token"`
		TokenType    tokenTypeEnums.TokenType `json:"token_type"`
		ExpiresIn    int                      `json:"expires_in"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (s StoreToken) GetAccessToken() string {
	return fmt.Sprintf("%s %s", s.Data.TokenType, s.Data.AccessToken)
}

func (s StoreToken) GetRefreshToken() string {
	return s.Data.RefreshToken
}
