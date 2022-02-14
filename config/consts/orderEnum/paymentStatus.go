package orderEnum

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/andhikagama/os-api/config/consts"
)

type (
	// PaymentStatusConst .
	PaymentStatusConst struct {
		Label   string `json:"label"`
		Name    string `json:"name"`
		DBValue string `json:"-"`
	}

	// PaymentStatus .
	PaymentStatus int
)

const (
	PaymentStatusPending PaymentStatus = iota + 1
	PaymentStatusPaid
	PaymentStatusExpired
)

var PaymentStatusConsts = []PaymentStatusConst{
	{"Pending", "pending", "pending"},
	{"Paid", "paid", "paid"},
	{"Expired", "expired", "expired"},
}

// Name .
func (c PaymentStatus) Name() string {
	if c < 1 || int(c) > len(PaymentStatusConsts) {
		return ""
	}
	return PaymentStatusConsts[c-1].Name
}

// DBValue .
func (c PaymentStatus) DBValue() string {
	if c < 1 || int(c) > len(PaymentStatusConsts) {
		return ""
	}
	return PaymentStatusConsts[c-1].DBValue
}

// UnmarshalParam parses value from the client (handled by gorm)
func (c *PaymentStatus) UnmarshalParam(src string) error {
	index := c.findIndex(src, func(c PaymentStatusConst) string {
		return c.Name
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentStatus(index)
	return nil
}

// MarshalJSON presents value to the client
func (c PaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name())
}

// UnmarshalJSON parses value from the client
func (c *PaymentStatus) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	index := c.findIndex(rawVal, func(c PaymentStatusConst) string {
		return c.Name
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentStatus(index)
	return nil
}

// Scan retrieves value from the DB
func (c *PaymentStatus) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return consts.ErrConstantParsing
	}
	dbVal := string(rawVal)

	index := c.findIndex(dbVal, func(c PaymentStatusConst) string {
		return c.DBValue
	})

	if index == 0 {
		return consts.ErrUnknownConstant
	}

	*c = PaymentStatus(index)
	return nil
}

// Value encodes value to the DB
func (c PaymentStatus) Value() (driver.Value, error) {
	return string(c.DBValue()), nil
}

func (c PaymentStatus) findIndex(code string, selector func(c PaymentStatusConst) string) int {
	for i, v := range PaymentStatusConsts {
		if selector(v) == code {
			return i + 1
		}
	}
	return 0
}

// AsCompleteConstants presents PaymentStatusConsts as their complete object form
func (c PaymentStatus) AsCompleteConstants() []PaymentStatusConst {
	list := make([]PaymentStatusConst, len(PaymentStatusConsts))
	copy(list, PaymentStatusConsts)
	return list
}
