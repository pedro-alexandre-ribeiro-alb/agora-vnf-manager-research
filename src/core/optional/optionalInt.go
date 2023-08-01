package optional

import (
	"database/sql/driver"
	"fmt"
)

type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64

func (dst Uint) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Uint(src), nil
	case int:
		return Uint(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Uint) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Uint8) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Uint8(src), nil
	case int:
		return Uint8(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Uint8) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Uint16) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Uint16(src), nil
	case int:
		return Uint16(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Uint16) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Uint32) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Uint32(src), nil
	case int:
		return Uint32(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Uint32) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Uint64) GetTypedValue(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Uint64(src), nil
	case int:
		return Uint64(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Uint64) Value() (driver.Value, error) {
	return int64(src), nil
}

type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64

func (dst Int) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Int(src), nil
	case int:
		return Int(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Int) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Int8) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Int8(src), nil
	case int:
		return Int8(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Int8) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Int16) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Int16(src), nil
	case int:
		return Int16(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Int16) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Int32) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Int32(src), nil
	case int:
		return Int32(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Int32) Value() (driver.Value, error) {
	return int64(src), nil
}

func (dst Int64) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case int64:
		return Int64(src), nil
	case int:
		return Int64(src), nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Int64) Value() (driver.Value, error) {
	return int64(src), nil
}
