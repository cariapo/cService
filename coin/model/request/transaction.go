package request

type Overbooking struct {
	AccessToken              *string `json:"-"`
	CustomerReferenceNumber  string  `json:"customer_reference_number"`
	SourceAccountNumber      string  `json:"source_account_number"`
	BeneficiaryAccountNumber string  `json:"beneficiary_account_number"`
	TrxAmount                float64 `json:"trx_amount"`
	Remark                   string  `json:"remark"`
}

type OverbookingMultiBeneficiary struct {
	AccessToken             *string              `json:"-" form:"-"`
	CustomerReferenceNumber string               `json:"customer_reference_number"`
	SourceAccountNumber     string               `json:"source_account_number"`
	BeneficiaryAccounts     []BeneficiaryAccount `json:"beneficiary_accounts"`
	TotalTrxAmount          float64              `json:"total_trx_amount"`
}

type BeneficiaryAccount struct {
	BeneficiaryAccountNumber string  `json:"beneficiary_account_number"`
	TrxAmount                float64 `json:"trx_amount"`
	Remark                   string  `json:"remark"`
}

type Reversal struct {
	AccessToken             *string `json:"-" form:"-"`
	CustomerReferenceNumber string  `json:"customer_reference_number"`
	ReferenceNumber         string  `json:"reference_number"`
	Amount                  float64 `json:"amount"`
}
