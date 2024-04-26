package cuan

import (
	"github.com/cariapo/cservice/cuan/model/request"
	"github.com/cariapo/cservice/cuan/model/response"
)

type CuanOutbound interface {
	FindProduct(req request.FindProduct) (*response.ListProduct, error)
	FindProductTelco(req request.FindProductTelco) (*response.ListProductAggregate, error)
	GetProductProvider(req request.GetProvider) (*response.Provider, error)
	GetProductByProductCode(req request.GetProduct) (*response.GetProduct, error)

	StoreInquiry(req request.StoreInquiry) (*response.Inquiry, error)

	StorePayment(req request.StorePayment) (*response.Payment, error)
	GetPaymentByCustomerRefNumber(req request.GetPaymentByCustomerRefNumber) (*response.Payment, error)
}
