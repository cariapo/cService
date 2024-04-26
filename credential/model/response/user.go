package response

type User struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id                     string `json:"id"`
		Username               string `json:"username"`
		IsActive               bool   `json:"is_active"`
		DefaultPassword        string `json:"default_password,omitempty"`
		DefaultPasswordExpired string `json:"default_password_expired,omitempty"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
