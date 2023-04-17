package request

type FindAccountByNumber struct {
	AccountNumber string
	AccessToken   string
}

type FindAccountById struct {
	Id          string
	AccessToken string
}
