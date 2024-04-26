package request

import (
	"github.com/cariapo/cservice/cuan/enum/productPlan"
	"github.com/cariapo/cservice/cuan/enum/productProvider"
	"github.com/cariapo/cservice/cuan/enum/productType"
	"github.com/cuwand/pondasi/models"
)

type FindProduct struct {
	models.Paging
	Provider    *productProvider.ProductProvider `form:"provider"`
	Plan        *productPlan.ProductPlan         `form:"plan"`
	Type        *productType.ProductType         `form:"type"`
	VendorCode  *string                          `form:"vendor_code"`
	AccessToken *string
}

type FindProductTelco struct {
	models.Paging
	PhoneNumber string
	Plan        productPlan.ProductPlan
	AccessToken *string
}

type GetProduct struct {
	ProductCode string
	AccessToken *string
}

type GetProvider struct {
	PhoneNumber string
	AccessToken *string
}
