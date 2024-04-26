package productPlan

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type ProductPlan string

const (
	PREPAID  ProductPlan = `PREPAID`
	POSTPAID ProductPlan = `POSTPAID`
	NONBILL  ProductPlan = `NONBILL`
	DATA     ProductPlan = `DATA`
	ROAMING  ProductPlan = `ROAMING`
	FINE     ProductPlan = `FINE`
)

func (p ProductPlan) ToCode() string {
	switch p {
	case PREPAID:
		return "00"
	case POSTPAID:
		return "01"
	case NONBILL:
		return "02"
	case DATA:
		return "03"
	case ROAMING:
		return "04"
	case FINE:
		return "05"
	default:
		return "undefined"
	}
}

func (s *ProductPlan) String() string {
	switch *s {
	case PREPAID:
		return "PREPAID"
	case POSTPAID:
		return "POSTPAID"
	case NONBILL:
		return "NONBILL"
	case DATA:
		return "DATA"
	case ROAMING:
		return "ROAMING"
	case FINE:
		return "FINE"
	default:
		return "undefined"
	}
}

func (s *ProductPlan) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux ProductPlan
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case PREPAID, POSTPAID, NONBILL, DATA, ROAMING, FINE:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for PREPAID, POSTPAID, NONBILL, DATA, ROAMING, FINE")
	}
}
