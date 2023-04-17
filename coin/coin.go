package coin

import (
	"github.com/cariapo/cservice/coin/model/request"
	"github.com/cariapo/cservice/coin/model/response"
)

type CoinOutbound interface {
	FindAccountByNumber(req request.FindAccountByNumber) (*response.Account, error)
	FindAccountById(req request.FindAccountById) (*response.Account, error)

	FindCustomerByCIF(req request.FindCustomerByCif) (*response.Customer, error)
	FindCustomerById(req request.FindCustomerById) (*response.Customer, error)

	Overbooking(req request.Overbooking) (*response.Overbooking, error)
	OverbookingMultiBeneficiary(req request.OverbookingMultiBeneficiary) (*response.OverbookingMultiBeneficiary, error)
	Reversal(req request.Reversal) (*response.Reversal, error)

	Statements(req request.Statement) (*response.ListStatement, error)
}
