package response

type Authorize struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Id          string      `json:"id"`
		Username    string      `json:"username,omitempty"`
		ClientId    string      `json:"client_id,omitempty"`
		Roles       []Role      `json:"roles,omitempty"`
		Authorities []Authority `json:"authorities,omitempty"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Role struct {
	Role        string       `json:"role"`
	Application string       `json:"application"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Permission  string `json:"permission"`
	Description string `json:"description"`
}

type Authority struct {
	Authority   string  `json:"authority"`
	Description string  `json:"description"`
	Scopes      []Scope `json:"scopes"`
}

type Scope struct {
	Scope       string `json:"scope"`
	Description string `json:"description"`
}
