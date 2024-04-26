package response

import (
	"github.com/cariapo/cservice/cuan/enum/productPlan"
	"github.com/cariapo/cservice/cuan/enum/productProvider"
	"github.com/cariapo/cservice/cuan/enum/productType"
	"github.com/cuwand/pondasi/helper/mapperHelper"
)

type Inquiry struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		CustomerReferenceNumber  string      `json:"customer_reference_number,omitempty"`
		ReferenceNumber          string      `json:"reference_number,omitempty"`
		RetrievalReferenceNumber string      `json:"retrieval_reference_number,omitempty"`
		Identifier               string      `json:"identifier,omitempty"`
		TrxAmount                int64       `json:"trx_amount"`
		TrxAdminFee              int64       `json:"trx_admin_fee"`
		TotalAmount              int64       `json:"total_amount"`
		BillInfo                 interface{} `json:"bill_info"`
		Product                  struct {
			Name        string                          `json:"name"`
			Description string                          `json:"description"`
			Information string                          `json:"information"`
			Identity    string                          `json:"identity"`
			Type        productType.ProductType         `json:"type"`
			Plan        productPlan.ProductPlan         `json:"plan"`
			Provider    productProvider.ProductProvider `json:"provider"`
		} `json:"product"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (i Inquiry) ToElectricityPrepaidBill() (bill ElectricityPrepaidBill) {
	if err := mapperHelper.InterfaceToStruct(i.Data.BillInfo, &bill); err != nil {
		panic(err)
	}

	return bill
}

type ElectricityPrepaidBill struct {
	CustomerId    string `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerClass string `json:"customer_class"` // R2
	Power         string `json:"power"`          // 3500VA
}
