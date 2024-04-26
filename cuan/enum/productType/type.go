package productType

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type ProductType string

//type ProductTypeCode string

const (
	TELECOMMUNICATION ProductType = `TELECOMMUNICATION`
	ELECTRICITY       ProductType = `ELECTRICITY`
	INSURANCE         ProductType = `INSURANCE`
	EWALLET           ProductType = `EWALLET`
	INTERNET          ProductType = `INTERNET`
	MULTIFINANCE      ProductType = `MULTIFINANCE`
	TRANSPORTATION    ProductType = `TRANSPORTATION`
	ENERGY            ProductType = `ENERGY`
)

func (p ProductType) ToCode() string {
	switch p {
	case TELECOMMUNICATION:
		return "00"
	case ELECTRICITY:
		return "01"
	default:
		return "undefined"
	}
}

func (s *ProductType) String() string {
	switch *s {
	case TELECOMMUNICATION:
		return "TELECOMMUNICATION"
	case ELECTRICITY:
		return "ELECTRICITY"
	case INSURANCE:
		return "INSURANCE"
	case EWALLET:
		return "EWALLET"
	case INTERNET:
		return "INTERNET"
	case MULTIFINANCE:
		return "MULTIFINANCE"
	case TRANSPORTATION:
		return "TRANSPORTATION"
	case ENERGY:
		return "ENERGY"
	default:
		return "undefined"
	}
}

func (s *ProductType) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux ProductType
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case ENERGY, TRANSPORTATION, MULTIFINANCE, INTERNET, EWALLET, ELECTRICITY, TELECOMMUNICATION:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for ENERGY, TRANSPORTATION, MULTIFINANCE, INTERNET, EWALLET, ELECTRICITY, TELECOMMUNICATION")
	}
}
