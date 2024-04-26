package request

import "github.com/cuwand/pondasi/models"

type FindCustomerByCif struct {
	CIF         string
	AccessToken *string
}

type FindCustomerById struct {
	Id          string
	AccessToken *string
}

type ExistsCustomerByPhoneNumber struct {
	PhoneNumber string
	AccessToken *string
}

type StoreCustomer struct {
	User        models.UserRequest `json:"-" form:"-"`
	Name        string             `json:"name"`
	PhoneNumber string             `json:"phone_number"`
	AccessToken *string            `json:"-" form:"-"`
}
