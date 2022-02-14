package orderEnum

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/andhikagama/os-api/config/consts"
)

type (
	// PaymentTypeConst .
	PaymentTypeConst struct {
		Label   string `json:"label"`
		Name    string `json:"name"`
		DBValue string `json:"-"`
	}

	// PaymentType .
	PaymentType int
)

const (
	PaymentTypeCOD PaymentType = iota + 1
	PaymentTypeVA
	PaymentTypeTransfer
)

var PaymentTypeConsts = []PaymentTypeConst{
	{"COD", "cod", "cod"},
	{"VA", "va", "va"},
	{"Transfer", "transfer", "transfer"},
}

// Name .
func (c PaymentType) Name() string {
	if c < 1 || int(c) > len(PaymentTypeConsts) {
		return ""
	}
	return PaymentTypeConsts[c-1].Name
}

// DBValue .
func (c PaymentType) DBValue() string {
	if c < 1 || int(c) > len(PaymentTypeConsts) {
		return ""
	}
	return PaymentTypeConsts[c-1].DBValue
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *PaymentType) UnmarshalParam(src string) error {
	index := c.findIndex(src, func(c PaymentTypeConst) string {
		return c.Name
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentType(index)
	return nil
}

// MarshalJSON presents value to the client
func (c PaymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name())
}

// UnmarshalJSON parses value from the client
func (c *PaymentType) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := c.findIndex(rawVal, func(c PaymentTypeConst) string {
		return c.Name
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentType(index)
	return nil
}

// Scan retrieves value from the DB
func (c *PaymentType) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return consts.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := c.findIndex(dbVal, func(c PaymentTypeConst) string {
		return c.DBValue
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentType(index)
	return nil
}

// Value encodes value to the DB
func (c PaymentType) Value() (driver.Value, error) {
	return string(c.DBValue()), nil
}

func (c PaymentType) findIndex(code string, selector func(c PaymentTypeConst) string) int {
	for i, v := range PaymentTypeConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents PaymentTypeConsts as their complete object form
func (c PaymentType) AsCompleteConstants() []PaymentTypeConst {
	list := make([]PaymentTypeConst, len(PaymentTypeConsts))
	copy(list, PaymentTypeConsts)
	return list
}
