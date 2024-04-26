package response

import (
	"github.com/cariapo/cservice/cuan/enum/productPlan"
	"github.com/cariapo/cservice/cuan/enum/productProvider"
	"github.com/cariapo/cservice/cuan/enum/productType"
	"github.com/cuwand/pondasi/models"
)

type Provider struct {
	Code      int    `json:"code"`
	Success   bool   `json:"success"`
	Data      string `json:"data"`
	Timestamp int64  `json:"timestamp"`
}

type ListProduct struct {
	Code      int                   `json:"code"`
	Success   bool                  `json:"success"`
	Products  []Product             `json:"data"`
	Paging    models.PagingResponse `json:"paging"`
	Timestamp int64                 `json:"timestamp"`
}

type ListProductAggregate struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
	Data    struct {
		Type     productType.ProductType          `json:"type"`
		Plan     productPlan.ProductPlan          `json:"plan"`
		Provider *productProvider.ProductProvider `json:"provider"`
		Products []Product                        `json:"products,omitempty"`
	} `json:"data"`
	Paging    models.PagingResponse `json:"paging"`
	Timestamp int64                 `json:"timestamp"`
}

type GetProduct struct {
	Code      int     `json:"code"`
	Success   bool    `json:"success"`
	Data      Product `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type Product struct {
	Id          string                          `json:"-"`
	Name        string                          `json:"name"`
	Description string                          `json:"description,omitempty"`
	Information string                          `json:"information,omitempty"`
	Denom       string                          `json:"denom"`
	Identity    string                          `json:"identity"`
	Type        productType.ProductType         `json:"type"`
	Plan        productPlan.ProductPlan         `json:"plan"`
	Provider    productProvider.ProductProvider `json:"provider"`
	Price       int64                           `json:"price"`
	AdminFee    int64                           `json:"admin_fee"`
	IsInquiry   bool                            `json:"is_inquiry"`
	IsActive    bool                            `json:"is_active"`
}
