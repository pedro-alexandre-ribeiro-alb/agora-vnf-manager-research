package types

import (
	"database/sql/driver"
	"fmt"
)

type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Labels []Label

func (dst Labels) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case Labels:
		return Labels(src), nil
	}
	return nil, fmt.Errorf("cannt scan %T", src)
}

func (src Labels) Value() (driver.Value, error) {
	return Labels(src), nil
}
