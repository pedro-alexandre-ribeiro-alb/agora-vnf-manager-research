package optional

import (
	"database/sql/driver"
	"fmt"
)

type String string

func (dst String) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case string:
		if String(src) == "" {
			return nil, nil
		} else {
			return String(src), nil
		}
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src String) Value() (driver.Value, error) {
	return string(src), nil
}
