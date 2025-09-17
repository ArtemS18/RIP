package ds

import (
	"database/sql/driver"
	"fmt"
)

type enumStatus string

const (
	DRAFT     enumStatus = "DRAFT"
	DELETED   enumStatus = "DELETED"
	COMPLITED enumStatus = "COMPLITED"
	FORMED    enumStatus = "FORMED"
	REJECTED  enumStatus = "REJECTED"
)

func (en *enumStatus) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type for enumStatus: %T", value)
	}
	*en = enumStatus(v)
	return nil
}
func (en enumStatus) Value() (driver.Value, error) {
	return string(en), nil
}
