package response

type Account struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id               string  `json:"id"`
		CustomerId       string  `json:"customer_id"`
		CIF              string  `json:"cif"`               // 312401
		Number           string  `json:"number"`            // 0000001
		EffectiveBalance float64 `json:"effective_balance"` // 23000
		ProductCode      string  `json:"product_code"`
		ProductName      string  `json:"product_name"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
