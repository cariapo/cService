package request

import (
	"github.com/cuwand/pondasi/enum/transactionFlowEnums"
	"github.com/cuwand/pondasi/enum/transactionStatusEnums"
	"github.com/cuwand/pondasi/models"
)

type Statement struct {
	models.Paging
	AccountNumber     string                                    `json:"-" form:"-"`
	AccessToken       string                                    `json:"-" form:"-"`
	TransactionStatus *transactionStatusEnums.TransactionStatus `form:"transaction_status"`
	TransactionFlow   *transactionFlowEnums.TransactionFlow     `form:"transaction_flow"`
}
