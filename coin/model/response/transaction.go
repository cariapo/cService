package response

import (
	"github.com/cuwand/pondasi/enum/transactionStatusEnums"
	"github.com/cuwand/pondasi/helper/dateHelper/dateProperty"
)

type Overbooking struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ReferenceNumber          string                                   `json:"reference_number"`
		CustomerReferenceNumber  string                                   `json:"customer_reference_number"`
		SourceAccountNumber      string                                   `json:"source_account_number"`
		BeneficiaryAccountNumber string                                   `json:"beneficiary_account_number"`
		TrxAmount                float64                                  `json:"trx_amount"`
		TrxFee                   float64                                  `json:"trx_fee"`
		Amount                   float64                                  `json:"amount"`
		Remark                   string                                   `json:"remark"`
		TransactionStatus        transactionStatusEnums.TransactionStatus `json:"transaction_status"`
		TransactionTime          dateProperty.DateTime                    `json:"transaction_time"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type OverbookingMultiBeneficiary struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		CustomerReferenceNumber string                                   `json:"customer_reference_number"`
		SourceAccountNumber     string                                   `json:"source_account_number"`
		BeneficiaryAccounts     []BeneficiaryAccount                     `json:"beneficiary_accounts"`
		Amount                  float64                                  `json:"amount"`
		TransactionStatus       transactionStatusEnums.TransactionStatus `json:"transaction_status"`
		TransactionTime         dateProperty.DateTime                    `json:"transaction_time"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type BeneficiaryAccount struct {
	ReferenceNumber          string  `json:"reference_number"`
	BeneficiaryAccountNumber string  `json:"beneficiary_account_number"`
	TrxAmount                float64 `json:"trx_amount"`
	TrxFee                   float64 `json:"trx_fee"`
	Remark                   string  `json:"remark"`
}

type Reversal struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		CustomerReferenceNumber string                                   `json:"customer_reference_number"`
		ReferenceNumber         string                                   `json:"reference_number"`
		Amount                  float64                                  `json:"amount"`
		TransactionStatus       transactionStatusEnums.TransactionStatus `json:"transaction_status"`
		TransactionTime         dateProperty.DateTime                    `json:"transaction_time"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
