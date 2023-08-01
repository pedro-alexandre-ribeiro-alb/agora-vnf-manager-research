package types

import "agora-vnf-manager/core/optional"

type VnfInfrastructure struct {
	Id                int                                `json:"id"`
	Name              optional.Optional[optional.String] `json:"name"`
	Description       optional.Optional[optional.String] `json:"description"`
	ConfigurationFile optional.Optional[optional.String] `json:"configurationFile"`
}

type VnfInfrastructureDocs struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	ConfigurationFile string `json:"configurationFile"`
}
