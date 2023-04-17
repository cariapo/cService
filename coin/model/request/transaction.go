package request

type Overbooking struct {
	AccessToken              string  `json:"-" form:"-"`
	CustomerReferenceNumber  string  `json:"customer_reference_number" binding:"required"`
	SourceAccountNumber      string  `json:"source_account_number" binding:"required"`
	BeneficiaryAccountNumber string  `json:"beneficiary_account_number" binding:"required"`
	TrxAmount                float64 `json:"trx_amount" binding:"required"`
	Remark                   string  `json:"remark" binding:"required"`
}

type OverbookingMultiBeneficiary struct {
	AccessToken             string               `json:"-" form:"-"`
	CustomerReferenceNumber string               `json:"customer_reference_number" binding:"required"`
	SourceAccountNumber     string               `json:"source_account_number" binding:"required"`
	BeneficiaryAccounts     []BeneficiaryAccount `json:"beneficiary_accounts" binding:"required"`
	TotalTrxAmount          float64              `json:"total_trx_amount" binding:"required"`
}

type BeneficiaryAccount struct {
	BeneficiaryAccountNumber string  `json:"beneficiary_account_number" binding:"required"`
	TrxAmount                float64 `json:"trx_amount" binding:"required"`
	Remark                   string  `json:"remark" binding:"required"`
}

type Reversal struct {
	AccessToken             string  `json:"-" form:"-"`
	CustomerReferenceNumber string  `json:"customer_reference_number"  binding:"required"`
	ReferenceNumber         string  `json:"reference_number"  binding:"required"`
	Amount                  float64 `json:"amount"  binding:"required"`
}
