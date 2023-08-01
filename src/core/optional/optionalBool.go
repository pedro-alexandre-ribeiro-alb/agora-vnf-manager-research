package optional

import (
	"database/sql/driver"
	"fmt"
)

type Bool bool

func (dst Bool) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case bool:
		return Bool(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Bool) Value() (driver.Value, error) {
	return bool(src), nil
}
