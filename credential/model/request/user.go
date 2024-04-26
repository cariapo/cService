package request

import "github.com/cuwand/pondasi/models"

type Register struct {
	//AccessToken *string            `json:"-" form:"-"`
	User     models.UserRequest `json:"-" form:"-"`
	Username string             `json:"username" binding:"required"`
	Password string             `json:"password" binding:"required"`
	Roles    []string           `json:"roles" binding:"required"`
}

type ChangePassword struct {
	//AccessToken     *string            `json:"-" form:"-"`
	User            models.UserRequest `json:"-" form:"-"`
	UserId          string             `json:"user_id" binding:"required"`
	CurrentPassword string             `json:"current_password" binding:"required"`
	NewPassword     string             `json:"new_password" binding:"required"`
	ConfirmPassword string             `json:"confirm_password" binding:"required"`
}

type ResetPassword struct {
	//AccessToken *string            `json:"-" form:"-"`
	User   models.UserRequest `json:"-" form:"-"`
	UserId string             `json:"user_id" binding:"required"`
}
