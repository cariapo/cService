package response

type Customer struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id   string `json:"id"`
		CIF  string `json:"cif"`  // 312401
		Name string `json:"name"` // ICHWAN ALMAZA
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type ExistsCustomer struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		CIF    string `json:"cif,omitempty"` // 312401
		Exists bool `json:"exists"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
