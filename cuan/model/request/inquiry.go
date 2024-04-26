package request

import "github.com/cuwand/pondasi/models"

type StoreInquiry struct {
	CustomerReferenceNumber string             `json:"customer_reference_number" binding:"required"`
	ProductCode             string             `json:"product_code" binding:"required"`
	Identifier              string             `json:"identifier" binding:"required"`
	Amount                  uint64             `json:"amount,default=10"`
	AccessToken             *string            `json:"-"`
	User                    models.UserRequest `json:"-"`
}
