package types

import (
	optional "agora-vnf-manager/core/optional"
)

type VnfInstance struct {
	Id                  int                                `json:"id"`
	Name                optional.Optional[optional.String] `json:"name"`
	Description         optional.Optional[optional.String] `json:"description"`
	Type                optional.Optional[Type]            `json:"type"`
	VnfInfraId          optional.Optional[optional.Int]    `json:"vnfInfraId"`
	Discovered          optional.Optional[optional.Bool]   `json:"discovered"`
	ManagementInterface optional.Optional[optional.String] `json:"managementInterface"`
	ControlInterface    optional.Optional[optional.String] `json:"controlInterface"`
	Vendor              optional.Optional[optional.String] `json:"vendor"`
	Version             optional.Optional[optional.String] `json:"version"`
}

type VnfInstanceDocs struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                Type   `json:"type"`
	VnfInfraId          int    `json:"vnfInfraId"`
	Discovered          bool   `json:"discovered"`
	ManagementInterface string `json:"managementInterface"`
	ControlInterface    string `json:"controlInterface"`
	Vendor              string `json:"vendor"`
	Version             string `json:"version"`
}
