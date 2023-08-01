package types

import "agora-vnf-manager/core/optional"

type VnfDeviceMapper struct {
	Id            int                                `json:"id"`
	DeviceId      optional.Optional[optional.String] `json:"deviceId"`
	VnfInstanceId optional.Optional[optional.Int]    `json:"vnfInstanceId"`
	ProxyId       optional.Optional[optional.Int]    `json:"proxyId"`
}

type VnfDeviceMapperDocs struct {
	Id            int    `json:"id"`
	DeviceId      string `json:"deviceId"`
	VnfInstanceId int    `json:"vnfInstanceId"`
	ProxyId       int    `json:"proxyId"`
}
