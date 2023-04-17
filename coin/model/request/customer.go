package request

type FindCustomerByCif struct {
	Cif         string
	AccessToken string
}

type FindCustomerById struct {
	Id          string
	AccessToken string
}
