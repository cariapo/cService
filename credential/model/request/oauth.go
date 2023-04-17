package request

import "github.com/cuwand/pondasi/enum/grandTypeEnums"

type StoreToken struct {
	GrantType    grandTypeEnums.GrandType `json:"grant_type" form:"grant_type" binding:"required"`
	Username     string                   `json:"username" form:"username"`
	Password     string                   `json:"password" form:"password"`
	RefreshToken string                   `json:"refresh_token" form:"refresh_token"`
}
