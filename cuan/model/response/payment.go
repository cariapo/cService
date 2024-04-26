package response

import (
	"github.com/cuwand/pondasi/enum/transactionStatusEnums"
	"github.com/cuwand/pondasi/helper/mapperHelper"
)

type Payment struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		ReferenceNumber          string                                   `json:"reference_number"`
		CustomerReferenceNumber  string                                   `json:"customer_reference_number"`
		RetrievalReferenceNumber *string                                  `json:"retrieval_reference_number,omitempty"`
		RequestDateTime          string                                   `json:"request_date_time"`
		ResponseDateTime         string                                   `json:"response_date_time,omitempty"`
		Identifier               string                                   `json:"identifier"`
		Status                   transactionStatusEnums.TransactionStatus `json:"status"`
		Reason                   string                                   `json:"reason"`
		BillAmount               int64                                    `json:"bill_amount"`
		TrxFeeAmount             int64                                    `json:"trx_fee_amount"`
		TrxAmount                int64                                    `json:"trx_amount"`
		AdditionalData           interface{}                              `json:"additional_data,omitempty"`
		Product                  struct {
			Name        string                          `json:"name"`
			Description string                          `json:"description"`
			Information string                          `json:"information"`
			Identity    string                          `json:"identity"`
			Type        string         `json:"type"`
			Plan        string         `json:"plan"`
			Provider    string `json:"provider"`
		}                              `json:"product"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (p Payment) ToElectricityPrepaidAdditionalData() (additionalData ElectricityPrepaidAdditionalData) {
	if err := mapperHelper.InterfaceToStruct(p.Data.AdditionalData, &additionalData); err != nil {
		panic(err)
	}

	return additionalData
}

func (p Payment) ToTelcoPrepaidAdditionalData() (additionalData TelcoPrepaidAdditionalData) {
	if err := mapperHelper.InterfaceToStruct(p.Data.AdditionalData, &additionalData); err != nil {
		panic(err)
	}

	return additionalData
}

type ElectricityPrepaidAdditionalData struct {
	CustomerName  string `json:"customer_name" bson:"customer_name"`   // PERUM MONAKO B
	CustomerClass string `json:"customer_class" bson:"customer_class"` // R2
	Power         string `json:"power" bson:"power"`                   // 3500VA
	Token         string `json:"token" bson:"token"`                   // 0621-9203-6330-2872-3082
	KWHTotal      string `json:"kwh_total" bson:"kwh_total"`           // 2,7KWH
}

type TelcoPrepaidAdditionalData struct {
	SerialNumber string `json:"serial_number" bson:"serial_number"`
}
