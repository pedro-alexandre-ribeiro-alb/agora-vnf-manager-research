package types

import optional "agora-vnf-manager/core/optional"

type Specification struct {
	ReleaseName string                             `json:"releaseName"`
	Namespace   string                             `json:"namespace"`
	ChartPath   optional.Optional[optional.String] `json:"chartPath"`
	ValuesPath  optional.Optional[optional.String] `json:"valuesPath"`
}
