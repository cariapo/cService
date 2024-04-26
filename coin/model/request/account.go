package request

import "github.com/cuwand/pondasi/models"

type FindAccountByNumber struct {
	AccountNumber string
	AccessToken   *string
}

type FindAccountById struct {
	Id          string
	AccessToken *string
}

type StoreAccount struct {
	CIF         string             `json:"cif"`
	ProductCode string             `json:"product_code"`
	Number      *string            `json:"number"`
	AccessToken *string            `json:"-" form:"-"`
	User        models.UserRequest `json:"-" form:"-"`
}
