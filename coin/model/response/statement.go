package response

import (
	"github.com/cariapo/cservice/coin/enum/transactionTypeEnums"
	"github.com/cuwand/pondasi/enum/transactionFlowEnums"
	"github.com/cuwand/pondasi/enum/transactionStatusEnums"
	"github.com/cuwand/pondasi/models"
)

type ListStatement struct {
	Code       int           		 `json:"code"`
	Success    bool          		 `json:"success"`
	Statements []Statement   		 `json:"data"`
	Paging     models.PagingResponse `json:"paging"`
	Timestamp  int64         		 `json:"timestamp"`
}

type Statement struct {
	ReferenceNumber         string                                   `json:"reference_number"`
	CustomerReferenceNumber string                                   `json:"customer_reference_number"`
	TransactionType         transactionTypeEnums.TransactionType     `json:"transaction_type"`   // OVERBOOKING
	TransactionFlow         transactionFlowEnums.TransactionFlow     `json:"transaction_flow"`   // DEBIT
	TransactionStatus       transactionStatusEnums.TransactionStatus `json:"transaction_status"` // SUCCESS
	Amount                  float64                                  `json:"amount"`             // 10000
	Remark                  string                                   `json:"remark"`             // ok
	TransactionTime         string                    				 `json:"transaction_time"`
}

type AccountTransaction struct {
	AccountId     string  `json:"account_id"`
	AccountNumber string  `json:"account_number"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
}
