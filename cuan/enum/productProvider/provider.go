package productProvider

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type ProductProvider string

const (
	TELKOMSEL ProductProvider = `TELKOMSEL`
	INDOSAT   ProductProvider = `INDOSAT`
	PLN       ProductProvider = `PLN`
)

func (p ProductProvider) ToCode() string {
	switch p {
	case TELKOMSEL:
		return "00"
	case INDOSAT:
		return "01"
	case PLN:
		return "02"
	default:
		return "undefined"
	}
}

func (s *ProductProvider) String() string {
	switch *s {
	case TELKOMSEL:
		return "TELKOMSEL"
	case INDOSAT:
		return "INDOSAT"
	case PLN:
		return "PLN"
	default:
		return "undefined"
	}
}

func (s *ProductProvider) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux ProductProvider
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case TELKOMSEL, INDOSAT, PLN:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for TELKOMSEL, INDOSAT, PLN")
	}
}
