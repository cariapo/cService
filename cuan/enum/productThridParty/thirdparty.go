package productThridParty

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type ProductThridparty string

const (
	SERPUL ProductThridparty = `SERPUL`
	DFLASH ProductThridparty = `DFLASH`
)

func (s *ProductThridparty) String() string {
	switch *s {
	case SERPUL:
		return "SERPUL"
	case DFLASH:
		return "DFLASH"
	default:
		return "undefined"
	}
}

func (s *ProductThridparty) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux ProductThridparty
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case SERPUL, DFLASH:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for SERPUL, DFLASH")
	}
}
