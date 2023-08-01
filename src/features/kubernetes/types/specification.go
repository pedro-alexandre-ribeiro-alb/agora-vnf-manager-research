package types

import optional "agora-vnf-manager/core/optional"

type Specifications struct {
	ConfigurationFile string                    `json:"configurationFile"`
	Labels            optional.Optional[Labels] `json:"labels"`
}
