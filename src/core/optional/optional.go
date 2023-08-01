package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"errors"
)

type optionable interface {
	Scan(any) (any, error)
	Value() (driver.Value, error)
}

type Optional[T optionable] struct {
	value   T
	defined bool
}

func New[DataType optionable](v DataType) Optional[DataType] {
	return Optional[DataType]{value: v, defined: true}
}

func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.defined = true
}

func (opt Optional[T]) Get() (T, error) {
	if !opt.Present() {
		var nilOpt Optional[T]
		return nilOpt.value, errors.New("value not present")
	}
	return opt.value, nil
}

func (opt Optional[T]) Present() bool {
	return opt.defined
}

func (opt Optional[T]) OrElse(v T) T {
	if opt.Present() {
		return opt.value
	}
	return v
}

func (opt Optional[T]) Fn(fn func(T)) {
	if opt.Present() {
		fn(opt.value)
	}
}

func (opt Optional[T]) MarshalJSON() ([]byte, error) {
	if opt.Present() {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

func (opt *Optional[T]) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		(*opt).defined = false
		return nil
	}

	var value T

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	(*opt).value = value
	(*opt).defined = true
	return nil

}

func (opt Optional[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(opt.value, start)
}

func (opt *Optional[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s T
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	(*opt).value = s
	(*opt).defined = true

	return nil
}

func (opt *Optional[T]) Scan(value interface{}) error {
	if value == nil {
		opt.defined = false
		return nil
	}
	if typedValue, err := opt.value.Scan(value); err != nil {
		return err
	} else {
		if typedValue == nil {
			opt.defined = false
			return nil
		}
		opt.defined, opt.value = true, typedValue.(T)
		return nil
	}
}

func (opt Optional[T]) Value() (driver.Value, error) {
	if !opt.defined {
		return nil, nil
	}
	return opt.value.Value()
}
