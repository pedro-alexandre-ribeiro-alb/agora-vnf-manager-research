package types

import (
	"database/sql/driver"
	"fmt"
)

type Type string

const (
	VomciService Type = "vomci/service"
	VomciProxy   Type = "vomci/proxy"
	DhcpService  Type = "dhcp/service"
	PppoeService Type = "pppoe/service"
)

var NetconfMapType map[Type]string = map[Type]string{
	VomciService: "bbf-nf-types:vomci-proxy-type",
	VomciProxy:   "bbf-nf-types:vomci-proxy-type",
	PppoeService: "bbf-nf-types:pppoe-ia-type",
}

var MapType map[string]Type = map[string]Type{
	"VOMCI":   VomciService,
	"VPROXY":  VomciProxy,
	"DHCPRA":  DhcpService,
	"PPPOEIA": PppoeService,
}

var from_string_map map[string]Type = map[string]Type{
	"vomci/service": VomciService,
	"vomci/proxy":   VomciProxy,
	"dhcp/service":  DhcpService,
	"pppoe/service": PppoeService,
}

func FromString(value string) (Type, error) {
	if value, ok := from_string_map[value]; ok {
		return value, nil
	}
	return "", fmt.Errorf("invalid Type: %s", value)
}

func (dst Type) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case Type:
		return Type(src), nil
	case string:
		build_type, err := FromString(src)
		if err != nil {
			return nil, err
		}
		return build_type, nil
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Type) Value() (driver.Value, error) {
	return Type(src), nil
}
