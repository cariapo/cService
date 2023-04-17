package transactionTypeEnums

import (
	"encoding/json"
	"github.com/cuwand/pondasi/errors"
)

type TransactionType string

const (
	Overbooking TransactionType = `OVERBOOKING`
	Reversal    TransactionType = `REVERSAL`
)

func (s *TransactionType) String() string {
	switch *s {
	case Overbooking:
		return "OVERBOOKING"
	case Reversal:
		return "REVERSAL"
	default:
		return "undefined"
	}
}

func (s *TransactionType) UnmarshalJSON(data []byte) error {
	// Define a secondary type so that we don't end up with a recursive call to json.Unmarshal
	type Aux TransactionType
	var a *Aux = (*Aux)(s)
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}

	// Validate the valid enum values
	switch *s {
	case Overbooking, Reversal:
		return nil
	default:
		*s = ""
		return errors.BadRequest("invalid value for TransactionStatus type, must Overbooking, Reversal")
	}
}
