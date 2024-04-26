package request

import "github.com/cuwand/pondasi/models"

type StorePayment struct {
	CustomerReferenceNumber  string             `json:"customer_reference_number" binding:"required"`
	RetrievalReferenceNumber *string            `json:"retrieval_reference_number"`
	ProductCode              string             `json:"product_code" binding:"required"`
	Identifier               string             `json:"identifier" binding:"required"`
	SourceAccount            string             `json:"source_account" binding:"required"`
	AccessToken              *string            `json:"-"`
	User                     models.UserRequest `json:"-"`
}

type GetPaymentByCustomerRefNumber struct {
	CustomerReferenceNumber string
	AccessToken             *string
}
